DROP PROCEDURE IF EXISTS `proc_confirm_item`;
DELIMITER $
CREATE PROCEDURE `proc_confirm_item` (IN residentID INT, IN itemType INT)
BEGIN
    DECLARE idCurrState INT;
    DECLARE idPreState INT;
	DECLARE idNextState INT;
	DECLARE idNextNextState INT;
	DECLARE typeNextState INT;
	DECLARE errorOccured INT;
	DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET errorOccured = 1;
	SET errorOccured = 0;	
    SET SQL_SAFE_UPDATES = 0;
	START TRANSACTION;
		UPDATE tb_item SET confirmed = 1
			WHERE resident_id = residentID AND type = itemType AND enabled = 1;
		SELECT id, next_item_id INTO idCurrState, idNextState FROM tb_item
			WHERE resident_id = residentID AND type = itemType AND enabled = 1;
		IF idNextState != -1 THEN
			-- 将下一个状态item设为enabled
			UPDATE tb_item SET enabled = 1 WHERE id = idNextState;
			SELECT next_item_id, type INTO idNextNextState, typeNextState
				FROM tb_item
				WHERE id = idNextState;
			IF typeNextState = 1 THEN -- 刚刚设置的状态为“在寺”
-- 				-- 将对应的“在寺”item设为confirmed
-- 				UPDATE tb_item SET confirmed = 1 WHERE id = idNextState;
				-- 将“在寺”的下一个状态设为enabled
				UPDATE tb_item SET enabled = 1 WHERE id = idNextNextState;
			END IF;
		ELSE -- 没有下一个状态（此处默认当前状态为PLAN_TO_LEAVE），将该住众与原先分配的孤邸解除绑定
			DELETE FROM tb_resi_status WHERE resident_id = residentID;
            -- 将对应的“在寺”状态设为confirmed
            SELECT id INTO idPreState FROM tb_item WHERE next_item_id = idCurrState;
            UPDATE tb_item SET confirmed = 1 WHERE id = idPreState;
		END IF;
	IF errorOccured = 0 THEN
		COMMIT;
        SELECT 0;
	ELSE
		ROLLBACK;
        SELECT -1;
	END IF;
	SET SQL_SAFE_UPDATES = 1;
END;