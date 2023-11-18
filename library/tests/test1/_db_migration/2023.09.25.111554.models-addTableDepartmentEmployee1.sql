        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `departm`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `isActive` BOOLEAN NOT NULL,
                `name_dep` VARCHAR(MAX) NOT NULL

                ,PRIMARY KEY (`ID`)
            )
            TABLESPACE pg_default;
                ;        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `employee1`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `isActive` BOOLEAN NOT NULL,
                `user_ID` INT NULL, ADD CONSTRAINT `fk_user_ID` FOREIGN KEY (`user_ID`) REFERENCES `user`(`ID`),
                `departm_ID` INT NULL, ADD CONSTRAINT `fk_departm_ID` FOREIGN KEY (`departm_ID`) REFERENCES `departm`(`ID`)

                ,PRIMARY KEY (`ID`)
            )
            TABLESPACE pg_default;
                