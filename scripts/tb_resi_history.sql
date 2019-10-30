CREATE TABLE `tb_resi_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resident_id` int(11) NOT NULL DEFAULT '0',
  `kuti_id` int(11) NOT NULL DEFAULT '0',
  `arrive_date` varchar(255) NOT NULL DEFAULT '',
  `leave_date` varchar(255) NOT NULL DEFAULT '',
  `comment` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;