-- phpMyAdmin SQL Dump
-- version 4.0.4.2
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Dec 18, 2022 at 08:47 AM
-- Server version: 5.6.13
-- PHP Version: 5.4.17

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `plutus-am`
--
CREATE DATABASE IF NOT EXISTS `plutus-am` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `plutus-am`;

-- --------------------------------------------------------

--
-- Table structure for table `assets`
--

CREATE TABLE IF NOT EXISTS `assets` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(32) NOT NULL,
  `type_id` int(11) NOT NULL,
  `currency_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=23 ;

--
-- Dumping data for table `assets`
--

INSERT INTO `assets` (`id`, `title`, `type_id`, `currency_id`) VALUES
(1, 'VTBR', 1, 810),
(10, 'YNDX', 1, 810),
(11, 'MAGN', 1, 810),
(12, 'FEES', 1, 810),
(14, 'NLMK', 1, 810),
(15, 'WIM Gold ETF', 1, 810),
(17, 'LQDT', 1, 810),
(18, 'MGNT', 1, 810),
(19, 'PHOR', 1, 810),
(20, 'MOEX', 1, 810),
(21, 'LKOH', 1, 810),
(22, 'GAZP', 1, 810);

-- --------------------------------------------------------

--
-- Table structure for table `assets_codes`
--

CREATE TABLE IF NOT EXISTS `assets_codes` (
  `asset` int(11) NOT NULL,
  `code` varchar(36) NOT NULL,
  KEY `asset` (`asset`,`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `assets_codes`
--

INSERT INTO `assets_codes` (`asset`, `code`) VALUES
(9, '0'),
(10, 'NL0009805522'),
(10, 'YNDX.RX'),
(11, 'MAGN'),
(15, 'GOLD');

-- --------------------------------------------------------

--
-- Table structure for table `bank_accounts`
--

CREATE TABLE IF NOT EXISTS `bank_accounts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `owner_id` int(11) NOT NULL,
  `title` varchar(64) NOT NULL,
  `currency_id` int(11) NOT NULL,
  `code` varchar(32) NOT NULL,
  `broker` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=3 ;

--
-- Dumping data for table `bank_accounts`
--

INSERT INTO `bank_accounts` (`id`, `owner_id`, `title`, `currency_id`, `code`, `broker`) VALUES
(2, 1, 'bank 1', 810, 'FKKD12', 'FFIN');

-- --------------------------------------------------------

--
-- Table structure for table `bank_inout`
--

CREATE TABLE IF NOT EXISTS `bank_inout` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `inout_date` date NOT NULL,
  `operation_id` int(11) NOT NULL,
  `bank_account_id` int(11) NOT NULL,
  `amount` float NOT NULL,
  `inout_type` varchar(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `currencies`
--

CREATE TABLE IF NOT EXISTS `currencies` (
  `id` int(11) NOT NULL,
  `title` varchar(32) NOT NULL,
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `currencies`
--

INSERT INTO `currencies` (`id`, `title`) VALUES
(810, 'RUB'),
(840, 'USD'),
(978, 'EUR');

-- --------------------------------------------------------

--
-- Table structure for table `currency_rates`
--

CREATE TABLE IF NOT EXISTS `currency_rates` (
  `period` date NOT NULL,
  `currency_id` int(11) NOT NULL,
  `rate` decimal(10,4) NOT NULL,
  UNIQUE KEY `period` (`period`,`currency_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `currency_rates`
--

INSERT INTO `currency_rates` (`period`, `currency_id`, `rate`) VALUES
('2022-07-30', 36, '42.9600'),
('2022-07-30', 51, '0.1503'),
('2022-07-30', 124, '47.7679'),
('2022-07-30', 156, '9.2163'),
('2022-07-30', 203, '2.5234'),
('2022-07-30', 208, '8.3364'),
('2022-07-30', 344, '7.8241'),
('2022-07-30', 348, '0.1550'),
('2022-07-30', 356, '0.7672'),
('2022-07-30', 392, '0.4609'),
('2022-07-30', 398, '0.1280'),
('2022-07-30', 410, '0.0471'),
('2022-07-30', 417, '0.7371'),
('2022-07-30', 498, '3.1714'),
('2022-07-30', 578, '6.3254'),
('2022-07-30', 702, '44.4760'),
('2022-07-30', 710, '3.7138'),
('2022-07-30', 752, '6.0486'),
('2022-07-30', 756, '64.3608'),
('2022-07-30', 826, '74.2711'),
('2022-07-30', 840, '61.3101'),
('2022-07-30', 860, '0.0056'),
('2022-07-30', 933, '23.4276'),
('2022-07-30', 934, '17.5172'),
('2022-07-30', 944, '36.0648'),
('2022-07-30', 946, '12.6641'),
('2022-07-30', 949, '3.4212'),
('2022-07-30', 960, '80.7405'),
('2022-07-30', 972, '5.9719'),
('2022-07-30', 975, '31.7291'),
('2022-07-30', 978, '62.5695'),
('2022-07-30', 980, '1.6755'),
('2022-07-30', 985, '13.1901'),
('2022-07-30', 986, '11.7574'),
('2022-08-02', 36, '43.4789'),
('2022-08-02', 51, '0.1522'),
('2022-08-02', 124, '48.3863'),
('2022-08-02', 156, '9.2822'),
('2022-08-02', 203, '2.5709'),
('2022-08-02', 208, '8.5009'),
('2022-08-02', 344, '7.9186'),
('2022-08-02', 348, '0.1578'),
('2022-08-02', 356, '0.7771'),
('2022-08-02', 392, '0.4683'),
('2022-08-02', 398, '0.1300'),
('2022-08-02', 410, '0.0475'),
('2022-08-02', 417, '0.7477'),
('2022-08-02', 498, '3.2125'),
('2022-08-02', 578, '6.4383'),
('2022-08-02', 702, '44.9544'),
('2022-08-02', 710, '3.7636'),
('2022-08-02', 752, '6.1038'),
('2022-08-02', 756, '65.1381'),
('2022-08-02', 826, '75.4659'),
('2022-08-02', 840, '62.0506'),
('2022-08-02', 860, '0.0056'),
('2022-08-02', 933, '23.3563'),
('2022-08-02', 934, '17.7287'),
('2022-08-02', 944, '36.5004'),
('2022-08-02', 946, '12.9221'),
('2022-08-02', 949, '3.4645'),
('2022-08-02', 960, '82.1302'),
('2022-08-02', 972, '6.0440'),
('2022-08-02', 975, '32.3534'),
('2022-08-02', 978, '63.2468'),
('2022-08-02', 980, '1.6801'),
('2022-08-02', 985, '13.4242'),
('2022-08-02', 986, '11.9602'),
('2022-08-30', 36, '41.3974'),
('2022-08-30', 51, '0.1490'),
('2022-08-30', 124, '46.4729'),
('2022-08-30', 156, '8.7216'),
('2022-08-30', 203, '2.4523'),
('2022-08-30', 208, '8.1213'),
('2022-08-30', 344, '7.7053'),
('2022-08-30', 348, '0.1460'),
('2022-08-30', 356, '0.7557'),
('2022-08-30', 392, '0.4342'),
('2022-08-30', 398, '0.1289'),
('2022-08-30', 410, '0.0447'),
('2022-08-30', 417, '0.7480'),
('2022-08-30', 498, '3.1190'),
('2022-08-30', 578, '6.1861'),
('2022-08-30', 702, '43.1477'),
('2022-08-30', 710, '3.5646'),
('2022-08-30', 752, '5.6435'),
('2022-08-30', 756, '62.2755'),
('2022-08-30', 826, '71.3860'),
('2022-08-30', 840, '60.3636'),
('2022-08-30', 860, '0.0055'),
('2022-08-30', 933, '23.5924'),
('2022-08-30', 934, '17.2467'),
('2022-08-30', 944, '35.5080'),
('2022-08-30', 946, '12.3400'),
('2022-08-30', 949, '3.3216'),
('2022-08-30', 960, '78.8632'),
('2022-08-30', 972, '5.8891'),
('2022-08-30', 975, '30.8844'),
('2022-08-30', 978, '59.9608'),
('2022-08-30', 980, '1.6424'),
('2022-08-30', 985, '12.7331'),
('2022-08-30', 986, '11.8593');

-- --------------------------------------------------------

--
-- Table structure for table `depo_accounts`
--

CREATE TABLE IF NOT EXISTS `depo_accounts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `owner_id` int(11) NOT NULL,
  `title` varchar(64) NOT NULL,
  `code` varchar(32) NOT NULL,
  `broker` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- Dumping data for table `depo_accounts`
--

INSERT INTO `depo_accounts` (`id`, `owner_id`, `title`, `code`, `broker`) VALUES
(1, 1, 'depo 1', 'FHFY12', 'FFIN');

-- --------------------------------------------------------

--
-- Table structure for table `depo_inout`
--

CREATE TABLE IF NOT EXISTS `depo_inout` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `inout_date` date NOT NULL,
  `operation_id` int(11) NOT NULL,
  `asset_id` int(11) NOT NULL,
  `depo_account_id` int(11) NOT NULL,
  `count` float NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `operations`
--

CREATE TABLE IF NOT EXISTS `operations` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `operation_date` date NOT NULL,
  `operation_number` varchar(32) NOT NULL,
  `type` varchar(16) NOT NULL,
  `trade_id` int(11) NOT NULL,
  `comment` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `owners`
--

CREATE TABLE IF NOT EXISTS `owners` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- Dumping data for table `owners`
--

INSERT INTO `owners` (`id`, `title`) VALUES
(1, 'Personal');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
