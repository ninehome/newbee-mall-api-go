-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: May 06, 2023 at 06:22 PM
-- Server version: 5.7.28
-- PHP Version: 7.3.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test`
--

-- --------------------------------------------------------

--
-- Table structure for table `tb_newbee_mall_user_recharge`
--

CREATE TABLE `tb_newbee_mall_user_recharge` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `user_name` varchar(11) COLLATE utf8_unicode_ci NOT NULL DEFAULT '''''',
  `money` int(11) NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `agent_id` varchar(11) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `tb_newbee_mall_user_recharge`
--

INSERT INTO `tb_newbee_mall_user_recharge` (`id`, `user_id`, `user_name`, `money`, `create_time`, `agent_id`) VALUES
(1, 631, '', 5000, '2023-05-06 15:43:53', '0'),
(2, 631, '123456788', 10000, '2023-05-06 15:47:06', '0'),
(3, 629, '9013639251', 3193, '2023-05-06 18:21:04', '6001');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tb_newbee_mall_user_recharge`
--
ALTER TABLE `tb_newbee_mall_user_recharge`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tb_newbee_mall_user_recharge`
--
ALTER TABLE `tb_newbee_mall_user_recharge`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
