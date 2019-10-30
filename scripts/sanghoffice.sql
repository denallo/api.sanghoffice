-- MySQL dump 10.13  Distrib 8.0.16, for Win64 (x86_64)
--
-- Host: localhost    Database: sanghoffice
-- ------------------------------------------------------
-- Server version	8.0.16

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tb_item`
--

DROP TABLE IF EXISTS `tb_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `tb_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resident_id` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL COMMENT '事项类型：0-人员于当天离开 1-已预约人员于当天到达',
  `confirmed` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已确认',
  `activate_date` varchar(10) DEFAULT NULL COMMENT '事项激活日期',
  PRIMARY KEY (`id`),
  KEY `resident_id_idx` (`resident_id`) /*!80000 INVISIBLE */,
  CONSTRAINT `resident_id` FOREIGN KEY (`resident_id`) REFERENCES `tb_resi_status` (`resident_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tb_kuti`
--

DROP TABLE IF EXISTS `tb_kuti`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `tb_kuti` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `number` int(11) NOT NULL DEFAULT '0',
  `type` int(11) NOT NULL DEFAULT '0',
  `for_sex` int(11) NOT NULL DEFAULT '0',
  `broken` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=148 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tb_resi_history`
--

DROP TABLE IF EXISTS `tb_resi_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `tb_resi_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resident_id` int(11) NOT NULL DEFAULT '0',
  `kuti_id` int(11) NOT NULL DEFAULT '0',
  `arrive_date` varchar(255) NOT NULL DEFAULT '',
  `leave_date` varchar(255) NOT NULL DEFAULT '',
  `comment` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tb_resi_status`
--

DROP TABLE IF EXISTS `tb_resi_status`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `tb_resi_status` (
  `resident_id` int(11) NOT NULL DEFAULT '0',
  `kuti_id` int(11) NOT NULL DEFAULT '0',
  `arrive_date` varchar(255) NOT NULL DEFAULT '',
  `plan_to_stay_days` int(11) NOT NULL DEFAULT '0',
  `plan_to_leave_date` varchar(255) NOT NULL DEFAULT '',
  `turned_phone_card` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`resident_id`),
  KEY `fk_kuti_id_idx` (`kuti_id`),
  CONSTRAINT `fk_kuti_id` FOREIGN KEY (`kuti_id`) REFERENCES `tb_kuti` (`id`),
  CONSTRAINT `fk_resident_id` FOREIGN KEY (`resident_id`) REFERENCES `tb_resident` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER `tr_create_item` AFTER INSERT ON `tb_resi_status` FOR EACH ROW BEGIN
	DECLARE currDate VARCHAR(10);
	INSERT INTO tb_item (resident_id, tb_item.type, confirmed, activate_date) VALUES (NEW.resident_id, 0, 0, '');
	SET currDate = CURDATE();
	IF NEW.arrive_date > currDate THEN
		INSERT INTO tb_item (resident_id, tb_item.type, confirmed, activate_date) VALUES (NEW.resident_id, 1, 0, '');
	END IF;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `tb_resident`
--

DROP TABLE IF EXISTS `tb_resident`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `tb_resident` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `dhamame` varchar(255) NOT NULL DEFAULT '',
  `sex` int(11) NOT NULL DEFAULT '0',
  `identifier` varchar(255) NOT NULL DEFAULT '',
  `age` int(11) NOT NULL DEFAULT '0',
  `type` int(11) NOT NULL DEFAULT '0',
  `folk` varchar(255) NOT NULL DEFAULT '',
  `native_place` varchar(255) NOT NULL DEFAULT '',
  `ability` varchar(255) NOT NULL DEFAULT '',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `emergency_contact` varchar(255) NOT NULL DEFAULT '',
  `emergency_contact_phone` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=344 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Temporary view structure for view `v_resi_status`
--

DROP TABLE IF EXISTS `v_resi_status`;
/*!50001 DROP VIEW IF EXISTS `v_resi_status`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8mb4;
/*!50001 CREATE VIEW `v_resi_status` AS SELECT 
 1 AS `kuti_number`,
 1 AS `kuti_type`,
 1 AS `for_sex`,
 1 AS `resident`,
 1 AS `sex`*/;
SET character_set_client = @saved_cs_client;

--
-- Final view structure for view `v_resi_status`
--

/*!50001 DROP VIEW IF EXISTS `v_resi_status`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `v_resi_status` (`kuti_number`,`kuti_type`,`for_sex`,`resident`,`sex`) AS select `tmp2`.`number` AS `number`,`tmp2`.`type` AS `type`,`tmp2`.`for_sex` AS `for_sex`,`tmp3`.`name` AS `name`,`tmp3`.`sex` AS `sex` from ((select `tb_resident`.`id` AS `id`,`tb_resident`.`name` AS `name`,`tb_resident`.`sex` AS `sex` from `tb_resident`) `tmp3` join (select `tb_kuti`.`number` AS `number`,`tb_kuti`.`type` AS `type`,`tb_kuti`.`for_sex` AS `for_sex`,`tmp`.`resident_id` AS `resident_id` from (`tb_kuti` join (select `tb_resi_status`.`resident_id` AS `resident_id`,`tb_resi_status`.`kuti_id` AS `kuti_id` from `tb_resi_status`) `tmp` on((`tb_kuti`.`id` = `tmp`.`kuti_id`)))) `tmp2` on((`tmp3`.`id` = `tmp2`.`resident_id`))) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-10-30 23:30:14
