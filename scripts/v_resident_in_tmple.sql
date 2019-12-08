DROP VIEW IF EXISTS `v_residents`;
CREATE 
    ALGORITHM = UNDEFINED 
    DEFINER = `root`@`localhost` 
    SQL SECURITY DEFINER
VIEW `v_residents` (`resident_id` , `name` , `dhamame` , `sex` , `type` , `kuti_number` , `kuti_type` , `arrive` , `leave`) AS
	SELECT tb_resident.id AS resident_id, tb_resident.name, tb_resident.dhamame, tb_resident.sex, tb_resident.type, 
			tb_kuti.number AS kuti_number, tb_kuti.type AS kuti_type,
			tb_resi_status.arrive_date AS arrive, tb_resi_status.plan_to_leave_date AS `leave`
	FROM tb_resident, tb_kuti, tb_resi_status
	WHERE tb_resident.id = tb_resi_status.resident_id AND tb_kuti.id = tb_resi_status.kuti_id;
