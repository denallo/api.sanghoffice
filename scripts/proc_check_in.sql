DROP PROCEDURE IF EXISTS `proc_check_in`;
DELIMITER $
CREATE PROCEDURE `proc_check_in` (
  IN name VARCHAR(50),
  IN dhamame VARCHAR(50),
  IN identifier VARCHAR(50),
  IN sex INT,
  IN age INT,
  IN type INT,
  IN folk VARCHAR(50),
  IN nativePlace VARCHAR(50),
  IN ability VARCHAR(50),
  IN phone VARCHAR(50),
  IN emergencyContact VARCHAR(50),
  IN emergencyContactPhone VARCHAR(50),
  IN kutiNumber INT,
  IN kutiType INT,
  IN isMonk INT,
  IN arriveDate VARCHAR(50),
  IN leaveDate VARCHAR(50)
)
BEGIN
  DECLARE rowCnt INT;
  DECLARE residentID INT;
  DECLARE kutiID INT;
  DECLARE currDate VARCHAR(10);
  DECLARE lastInsertedID INT;
  DECLARE idPlanToLeave INT;
  DECLARE errorOccured INT;
  DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET errorOccured = 1;
  SET errorOccured = 0;
  SET residentID = -1;
  START TRANSACTION;
    -- ------------------- 如果resident不存在则创建 ------------------------
	SELECT id INTO residentID FROM tb_resident
	WHERE (name = @name AND name != '') 
	OR (dhamame = @dhamame AND dhamame != '') LIMIT 1;
    IF -1 = residentID THEN
      INSERT INTO tb_resident (name, dhamame, sex, identifier, age, type, folk, native_place, ability, phone, emergency_contact, emergency_contact_phone)
      VALUES (name, dhamame, sex, identifier, age, type, folk, native_place, ability, phone, emergencyContact, emergencyContactPhone);
      SET residentID = LAST_INSERT_ID();
	END IF;
    -- ------------------ 创建该resident本次入住的resi_status -----------------------
    SELECT id INTO kutiID FROM tb_kuti WHERE for_sex = sex AND number = kutiNumber AND tb_kuti.type = kutiType;
    -- SELECT id INTO kutiID FROM tb_kuti WHERE for_sex = sex AND number = kutiNumber;
    INSERT INTO tb_resi_status 
		(resident_id, kuti_id, arrive_date, plan_to_stay_days, plan_to_leave_date, turned_phone_card)
	VALUES (residentID, kutiId, arriveDate, 0, leaveDate, 0);
    -- ------------------ 创建对应的待办事项 ---------------------------
	-- 计划离开
	INSERT INTO tb_item (next_item_id, resident_id, tb_item.type, enabled, confirmed, activate_date)
		VALUES (-1, residentID, 2, 0, 0, leaveDate);
	SET lastInsertedID = LAST_INSERT_ID();
	-- 在寺
	INSERT INTO tb_item (next_item_id, resident_id, tb_item.type, enabled, confirmed, activate_date)
		VALUES (lastInsertedID, residentID, 1, 0, 0, '');
	SET lastInsertedID = LAST_INSERT_ID();
	-- 预约到达
	SET currDate = CURDATE();
	IF arriveDate > currDate THEN
		INSERT INTO tb_item (next_item_id, resident_id, tb_item.type, enabled, confirmed, activate_date)
			VALUES (lastInsertedID, residentID, 0, 1, 0, arriveDate);
	ELSE -- 当天入住，将代表“在寺”状态的record设为enabled；将“计划离开”设为enabled
		UPDATE tb_item SET enabled = 1 WHERE id = lastInsertedID;
        SELECT next_item_id INTO idPlanToLeave FROM tb_item WHERE id = lastInsertedID; 
        UPDATE tb_item SET enabled = 1 WHERE id = idPlanToLeave;
	END IF;
  IF errorOccured = 0 THEN
    COMMIT;
    SELECT residentID;
  ELSE
	ROLLBACK;
    SELECT -100;
  END IF;
END