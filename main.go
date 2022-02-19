package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/payment"
	"bwastartup/transaction"
	"bwastartup/user"
	webhandler "bwastartup/web/handler"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	/**
	* Connecting to database
	 */
	dsn := "root:passw0rd@tcp(127.0.0.1:3306)/crowdfund?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// repository instance
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	// service instance
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	campaignService := campaign.NewService(campaignRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	// handler instance
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	userWebHandler := webhandler.NewUserHandler(userService)
	campaignWebHandler := webhandler.NewCampaignHandler(campaignService, userService)
	transactionWebHandler := webhandler.NewTransactionHandler(transactionService)
	sessionWebHandler := webhandler.NewSessionHandler(userService)

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))

	// init gin router
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(sessions.Sessions("bwastartup", cookieStore))
	// serving static files
	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")

	router.HTMLRender = loadTemplates("./web/templates")
	// set api group for `/api/v1`
	api := router.Group("/api/v1")

	// list of router
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)
	api.GET("/campaign/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransaction)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	authAdminMiddlewareIns := authAdminMiddleware()

	// router for CMS
	router.GET("/users", authAdminMiddlewareIns, userWebHandler.Index)
	router.GET("/users/new", authAdminMiddlewareIns, userWebHandler.New)
	router.POST("/users", authAdminMiddlewareIns, userWebHandler.Create)
	router.GET("/users/edit/:id", authAdminMiddlewareIns, userWebHandler.Edit)
	router.POST("/users/update/:id", authAdminMiddlewareIns, userWebHandler.Update)
	router.GET("/users/avatar/:id", authAdminMiddlewareIns, userWebHandler.NewAvatar)
	router.POST("/users/avatar/:id", authAdminMiddlewareIns, userWebHandler.CreateAvatar)

	router.GET("/campaigns", authAdminMiddlewareIns, campaignWebHandler.Index)
	router.GET("/campaigns/new", authAdminMiddlewareIns, campaignWebHandler.New)
	router.POST("/campaigns", authAdminMiddlewareIns, campaignWebHandler.Create)
	router.GET("/campaigns/images/:id", authAdminMiddlewareIns, campaignWebHandler.NewImage)
	router.POST("/campaigns/images/:id", authAdminMiddlewareIns, campaignWebHandler.CreateImage)
	router.GET("/campaigns/edit/:id", authAdminMiddlewareIns, campaignWebHandler.Edit)
	router.POST("/campaigns/update/:id", authAdminMiddlewareIns, campaignWebHandler.Update)
	router.GET("/campaigns/show/:id", authAdminMiddlewareIns, campaignWebHandler.Show)

	router.GET("/transactions", authAdminMiddlewareIns, transactionWebHandler.Index)

	router.GET("/login", sessionWebHandler.New)
	router.GET("/logout", sessionWebHandler.Logout)
	router.POST("/session", sessionWebHandler.Login)

	// listen server on port 3001
	router.Run(":3001")

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string = ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

func authAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		//userID
		userIDSession := session.Get("userID")

		if userIDSession == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}

	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
