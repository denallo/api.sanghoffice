CREATE TABLE `tb_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `next_item_id` int(11) DEFAULT NULL,
  `resident_id` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL COMMENT '事项类型：0-预约到达 1-在寺 2-计划离开',
  `enabled` tinyint(4) NOT NULL DEFAULT '0',
  `confirmed` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已确认',
  `activate_date` varchar(10) DEFAULT NULL COMMENT '事项激活日期',
  PRIMARY KEY (`id`),
  KEY `resident_id_idx` (`resident_id`) /*!80000 INVISIBLE */,
  CONSTRAINT `resident_id` FOREIGN KEY (`resident_id`) REFERENCES `tb_resident` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
