CREATE 
  ALGORITHM = UNDEFINED 
  DEFINER = `root`@`localhost` 
  SQL SECURITY DEFINER
VIEW `v_resi_status` (`kuti_number` , `kuti_type` , `for_sex` , `resident` , `sex`) AS
  SELECT 
    `tmp2`.`number` AS `number`,
    `tmp2`.`type` AS `type`,
    `tmp2`.`for_sex` AS `for_sex`,
    `tmp3`.`name` AS `name`,
    `tmp3`.`sex` AS `sex`
  FROM
  ((SELECT 
    `tb_resident`.`id` AS `id`,
      `tb_resident`.`name` AS `name`,
      `tb_resident`.`sex` AS `sex`
  FROM
	`tb_resident`) `tmp3`
  JOIN (SELECT 
    `tb_kuti`.`number` AS `number`,
      `tb_kuti`.`type` AS `type`,
      `tb_kuti`.`for_sex` AS `for_sex`,
      `tmp`.`resident_id` AS `resident_id`
  FROM
    (`tb_kuti`
  JOIN (SELECT 
    `tb_resi_status`.`resident_id` AS `resident_id`,
      `tb_resi_status`.`kuti_id` AS `kuti_id`
  FROM
    `tb_resi_status`) `tmp` ON ((`tb_kuti`.`id` = `tmp`.`kuti_id`)))) `tmp2` ON ((`tmp3`.`id` = `tmp2`.`resident_id`)))