
            ALTER TABLE `employee1`
            ADD COLUMN `departm2_ID` INT NULL, ADD CONSTRAINT `fk_departm2_ID` FOREIGN KEY (`departm2_ID`) REFERENCES `departm`(`ID`)