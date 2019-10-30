CREATE TABLE `tb_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resident_id` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL COMMENT '事项类型：0-人员于当天离开 1-已预约人员于当天到达',
  `confirmed` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已确认',
  `activate_date` varchar(10) DEFAULT NULL COMMENT '是否已激活提示',
  PRIMARY KEY (`id`),
  KEY `resident_id_idx` (`resident_id`) /*!80000 INVISIBLE */,
  CONSTRAINT `resident_id` FOREIGN KEY (`resident_id`) REFERENCES `tb_resi_status` (`resident_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
