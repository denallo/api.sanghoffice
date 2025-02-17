DROP TABLE IF EXISTS `tb_resi_status`;
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
