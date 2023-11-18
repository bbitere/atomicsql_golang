
            ALTER TABLE `project`
            ADD COLUMN `projStatus2_ID` INT NULL, ADD CONSTRAINT `fk_projStatus2_ID` FOREIGN KEY (`projStatus2_ID`) REFERENCES `projstatus`(`ID`);
            ALTER TABLE `project`
            DROP COLUMN `projStatus_ID`