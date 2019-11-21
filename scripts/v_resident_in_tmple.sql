DROP VIEW IF EXISTS `v_residents`;
CREATE 
    ALGORITHM = UNDEFINED 
    DEFINER = `root`@`localhost` 
    SQL SECURITY DEFINER
VIEW `v_residents` (`resident_id` , `name` , `dhamame` , `sex` , `type` , `kuti_number` , `kuti_type` , `arrive` , `leave`) AS
    SELECT 
        `tmp_resident`.`resident_id` AS `resident_id`,
        `tmp_resident`.`name` AS `name`,
        `tmp_resident`.`dhamame` AS `dhamame`,
        `tmp_resident`.`sex` AS `sex`,
        `tmp_resident`.`type` AS `type`,
        `tmp_resi_status`.`kuti_number` AS `kuti_number`,
        `tmp_resi_status`.`kuti_type` AS `kuti_type`,
        `tmp_resi_status`.`arrive` AS `arrive`,
        `tmp_resi_status`.`leave` AS `leave`
    FROM
        ((SELECT 
            `tb_resident`.`id` AS `resident_id`,
                `tb_resident`.`name` AS `name`,
                `tb_resident`.`dhamame` AS `dhamame`,
                `tb_resident`.`sex` AS `sex`,
                `tb_resident`.`type` AS `type`
        FROM
            `tb_resident`) `tmp_resident`
        JOIN (SELECT 
            `tmp_resi_status`.`id` AS `id`,
                `tmp_kuti`.`kuti_number` AS `kuti_number`,
                `tmp_kuti`.`kuti_type` AS `kuti_type`,
                `tmp_resi_status`.`arrive` AS `arrive`,
                `tmp_resi_status`.`leave` AS `leave`
			FROM
				((SELECT 
				`tb_resi_status`.`resident_id` AS `id`,
					`tb_resi_status`.`kuti_id` AS `kuti_id`,
					`tb_resi_status`.`arrive_date` AS `arrive`,
					`tb_resi_status`.`plan_to_leave_date` AS `leave`
				FROM
					`tb_resi_status`) `tmp_resi_status`
			JOIN (SELECT 
				`tb_kuti`.`id` AS `kuti_id`,
					`tb_kuti`.`type` AS `kuti_type`,
					`tb_kuti`.`number` AS `kuti_number`
				FROM
					`tb_kuti`) `tmp_kuti`
			ON ((`tmp_resi_status`.`kuti_id` = `tmp_kuti`.`kuti_id`)))) `tmp_resi_status`
		ON ((`tmp_resident`.`resident_id` = `tmp_resi_status`.`id`)))