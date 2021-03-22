-- phpMyAdmin SQL Dump
-- version 5.0.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 08 Feb 2021 pada 15.58
-- Versi server: 10.4.14-MariaDB
-- Versi PHP: 7.4.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bwastartup`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `short_description` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `perks` text NOT NULL,
  `backer_count` int(11) NOT NULL,
  `goal_amount` int(11) NOT NULL,
  `current_amount` int(11) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int(11) NOT NULL,
  `campaign_id` int(11) NOT NULL,
  `file_name` varchar(255) NOT NULL,
  `is_primary` tinyint(4) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `campaign_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `status` varchar(255) NOT NULL,
  `code` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `occupation` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `avatar_file_name` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `created_at`, `updated_at`) VALUES
(1, 'ramses', 'programmer', 'ramseshutasoit77@gmail.com', '$2a$04$T2tIbfn5C7P/bxuVS8HAb.knlNhBjdYw8NUhFr6zoNDQJ0BcmEUE2', 'avatar.jpg', 'user', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(2, 'Test Simpan', '', '', '', '', '', '2021-02-06 15:52:36', '2021-02-06 15:52:36'),
(3, 'Test simpan dari service', 'Dev', 'ramses1@gmail.com', '$2a$04$T2tIbfn5C7P/bxuVS8HAb.knlNhBjdYw8NUhFr6zoNDQJ0BcmEUE2', '', 'user', '2021-02-06 16:16:13', '2021-02-06 16:16:13'),
(4, 'Hutasoit', 'Designer', 'hutasoit@gmail.com', '$2a$04$A5bwcA/yP5q3yVkIwycM8ei.8BP2I/WDv3M55iq8lgYomAqgXapGu', '', 'user', '2021-02-06 16:53:14', '2021-02-06 16:53:14'),
(5, 'Budiman', 'Designer', 'hutasoit@gmail.com', '$2a$04$DrjOTRuVfxk.F.MyrO4tVeGf0LjpewsBbhJq3XvU22BXDfa6VgiXq', '', 'user', '2021-02-06 17:10:59', '2021-02-06 17:10:59'),
(6, 'Budiman', 'Designer', 'hutasoit@gmail.com', '$2a$04$qWqWY/ZhWY4nkJrFjxq91urU60bljm2/HJs1yNcmPVnOK9/Wa.Qfm', '', 'user', '2021-02-06 17:14:17', '2021-02-06 17:14:17'),
(7, 'Budiman', 'Designer', 'hutasoit@gmail.com', '$2a$04$P8bglHeXZJu62US93dHJ9eItPKR1K8PueU4WYQJOpASJ4SDBaJGf6', '', 'user', '2021-02-06 17:18:01', '2021-02-06 17:18:01'),
(8, 'Budiman', 'Designer', 'hutasoit@gmail.com', '$2a$04$Pf/w/VP9b1TeOGEF.O9sVe2T.NlIHb8jqf6f73s6b7QkuYtWZ2szW', '', 'user', '2021-02-06 17:29:22', '2021-02-06 17:29:22'),
(9, 'Budiman', 'Designer', 'hutasoit@gmail.com', '$2a$04$n2.GD7GtFW0oj99.b3FEmudAsNJZc6N85O69.Iawco7TmvKX5.Ikq', '', 'user', '2021-02-06 17:35:06', '2021-02-06 17:35:06'),
(10, 'Budiman', 'Designer', 'hutasoit@gmail.com', '$2a$04$um5BPsNUx5Xd/RernQ0CUu.w0ZtxG69sJwSNNyfzjsBjOb4F0jIwO', '', 'user', '2021-02-06 17:35:22', '2021-02-06 17:35:22'),
(11, 'Budiman', 'Designer', 'hutasoit@gmail.com', '$2a$04$lclxvjfeSnvfByqU4u3yAepmIKQRcVr9pEREtneF7cZtAIUAD9Le2', '', 'user', '2021-02-06 17:40:30', '2021-02-06 17:40:30'),
(12, 'Budiman', 'Designer', 'hutasoit@gmail.com', '$2a$04$B5vL/HrldamefUiQ5WOF5.rLkv2Ae1VQshjghoqOXeteYvFuMLWw6', '', 'user', '2021-02-06 17:42:27', '2021-02-06 17:42:27'),
(13, 'Budiman', 'Designer', 'hutasoit@gmail.com', '$2a$04$hg6uCBWWu2zjBSAfMVcP2eiLotONp5hOQ2k77UzbfBgWFLde0pldy', '', 'user', '2021-02-06 17:42:38', '2021-02-06 17:42:38'),
(14, 'Budiman', 'Designer', 'hutasoit@gmail.com', '$2a$04$O7V4gQgMeKfA3vs54zQWQuSnp6uwFOSA1L5Cf7D7EhTOLyZ3X3kcy', '', 'user', '2021-02-06 17:49:32', '2021-02-06 17:49:32');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fkCampaignsUsers` (`user_id`);

--
-- Indeks untuk tabel `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fkCampaignImagesCampaigns` (`campaign_id`);

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fkTransactionsUsers` (`user_id`),
  ADD KEY `fkTransactionsCampaigns` (`campaign_id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `campaigns`
--
ALTER TABLE `campaigns`
  ADD CONSTRAINT `fkCampaignsUsers` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD CONSTRAINT `fkCampaignImagesCampaigns` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `fkTransactionsCampaigns` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `fkTransactionsUsers` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
