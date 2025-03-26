use exchange;

DROP PROCEDURE IF EXISTS DropTablesInSchema;

DELIMITER //

CREATE PROCEDURE DropTablesInSchema(IN schema_name VARCHAR(255))
BEGIN
    DECLARE done INT DEFAULT 0;
    DECLARE var_table_name VARCHAR(255);

    DECLARE cur CURSOR FOR
        SELECT table_name
        FROM information_schema.tables
        WHERE table_schema = schema_name;

    DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = 1;

    OPEN cur;

    read_loop: LOOP
        FETCH cur INTO var_table_name;
        IF done THEN
            LEAVE read_loop;
        END IF;

        IF var_table_name IS NOT NULL THEN
            SET @stmt = CONCAT('DROP TABLE IF EXISTS `', schema_name, '`.`', var_table_name, '`;');
            PREPARE drop_statement FROM @stmt;
            EXECUTE drop_statement;
            DEALLOCATE PREPARE drop_statement;
        END IF;
    END LOOP;

    CLOSE cur;
END //

DELIMITER ;


/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
CALL DropTablesInSchema('exchange');
DROP PROCEDURE IF EXISTS DropTablesInSchema;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
