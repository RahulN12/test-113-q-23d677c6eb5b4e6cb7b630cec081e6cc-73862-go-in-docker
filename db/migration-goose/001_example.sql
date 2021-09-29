-- +goose Up

  CREATE TABLE `quiz` (
  `id` bigint(10) PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL);

  CREATE TABLE `questions` (
  `id` bigint(10) PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `options` varchar(255) NOT NULL,
  `correct_option` bigint NOT NULL,
  `quiz` bigint NOT NULL,
  `points` bigint NOT NULL);