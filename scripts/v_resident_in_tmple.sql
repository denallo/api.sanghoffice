DROP VIEW IF EXISTS `v_resident_in_temple`;
CREATE VIEW v_resident_in_temple (`resident_id`, `name`, `dhamame`, `sex`, `type`, `kuti_number`, `kuti_type`, `arrive`, `leave`) AS
  SELECT `resident_id`, `name`, `dhamame`, `sex`, `type`, `kuti_number`, `kuti_type`, `arrive`, `leave` FROM (
    (SELECT `id` AS `resident_id`, `name`, `dhamame`, `sex`, `type` FROM `tb_resident`) `tmp_resident`
    JOIN
    (SELECT `id`, `kuti_number`, `kuti_type`, `arrive`, `leave` FROM (
      (SELECT `resident_id` AS `id`, `kuti_id`, `arrive_date` AS `arrive`, `plan_to_leave_date` AS `leave`
       FROM `tb_resi_status`) `tmp_resi_status`
       -- WHERE `arrive_date` < DATE_FORMAT(NOW(),'%Y-%m-%d')) `tmp_resi_status`
      JOIN
      (SELECT `id` AS `kuti_id`, `type` AS `kuti_type`, `number` AS `kuti_number` from `tb_kuti`) `tmp_kuti`
      ON (`tmp_resi_status`.`kuti_id` = `tmp_kuti`.`kuti_id`))) `tmp_resi_status` 
    ON (`tmp_resident`.`resident_id` = `tmp_resi_status`.`id`));