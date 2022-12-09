-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.6.21-log - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 test 的数据库结构
CREATE DATABASE IF NOT EXISTS `test` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `test`;

-- 导出  表 test.tb_newbee_mall_user_bank 结构
CREATE TABLE IF NOT EXISTS `tb_newbee_mall_user_bank` (
  `bank_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '银行id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `bank_name` char(50) NOT NULL DEFAULT '0' COMMENT '开户行名称',
  `user_name` char(50) NOT NULL DEFAULT '0' COMMENT '账户持有人名字',
  `bank_number` char(50) NOT NULL DEFAULT '0' COMMENT '银行卡号码',
  `default` tinyint(3) unsigned zerofill NOT NULL DEFAULT '000' COMMENT '是否 默认选项',
  PRIMARY KEY (`bank_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='用户银行账户';

-- 正在导出表  test.tb_newbee_mall_user_bank 的数据：~2 rows (大约)
INSERT INTO `tb_newbee_mall_user_bank` (`bank_id`, `user_id`, `bank_name`, `user_name`, `bank_number`, `default`) VALUES
	(1, 7, 'russbank', 'noe', '123333', 001),
	(2, 7, 'ssusus', 'neo', '34343', 000);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
