-- --------------------------------------------------------
-- 主機:                           192.168.56.105
-- 伺服器版本:                        10.5.11-MariaDB-1:10.5.11+maria~focal - mariadb.org binary distribution
-- 伺服器作業系統:                      debian-linux-gnu
-- HeidiSQL 版本:                  11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 傾印 UserInfo 的資料庫結構
CREATE DATABASE IF NOT EXISTS `UserInfo` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `UserInfo`;

-- 傾印  資料表 UserInfo.AccountInfo 結構
CREATE TABLE IF NOT EXISTS `AccountInfo` (
  `username` varchar(50) NOT NULL COMMENT '帳號',
  `name` varchar(50) NOT NULL COMMENT '姓名',
  `password` varchar(50) NOT NULL COMMENT '密碼',
  `age` int(11) NOT NULL DEFAULT 0 COMMENT '年齡',
  `email` varchar(50) NOT NULL DEFAULT '0' COMMENT '電子郵件',
  `status` varchar(50) NOT NULL DEFAULT '0' COMMENT '帳號狀態',
  `kkc` int(11) NOT NULL COMMENT '代幣總額',
  `totalexchange` int(11) NOT NULL COMMENT '總交換次數',
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在傾印表格  UserInfo.AccountInfo 的資料：~0 rows (近似值)
/*!40000 ALTER TABLE `AccountInfo` DISABLE KEYS */;
/*!40000 ALTER TABLE `AccountInfo` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
