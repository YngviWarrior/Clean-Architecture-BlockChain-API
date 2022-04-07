CREATE DATABASE `go`;

CREATE TABLE IF NOT EXISTS `exchange` (
  `exchange` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `website` varchar(255) NOT NULL,
  PRIMARY KEY (`exchange`)
);