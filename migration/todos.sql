-- phpMyAdmin SQL Dump
-- version 4.8.2
-- https://www.phpmyadmin.net/
--
-- Host: mysql
-- Generation Time: Jul 25, 2018 at 04:36 PM
-- Server version: 8.0.11
-- PHP Version: 7.2.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";

--
-- Database: `todos`
--
CREATE DATABASE IF NOT EXISTS `todos` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `todos`;

GRANT USAGE ON *.* TO `todo_app`@`%`;
GRANT SELECT, INSERT, UPDATE, DELETE ON `todos`.* TO `todo_app`@`%`;

-- --------------------------------------------------------

--
-- Table structure for table `todos`
--

CREATE TABLE IF NOT EXISTS `todos` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` text COLLATE utf8mb4_unicode_ci,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `todos` (`id`, `title`, `content`) VALUES
('260954e3-8a3c-41d9-87c4-371db1241e83', 'Comprar leite', NULL),
('7777266a-bed7-47b2-addb-277cf5cff410', 'Pegar as crianças na escola', 'Não esquecer de pagar a mensalidade também.'),
('c989d151-e48f-4b88-b8fa-07ec9917f4e7', 'Trocar o chuveiro', NULL);

COMMIT;
