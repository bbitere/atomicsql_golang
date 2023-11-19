        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `departm`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `isActive` BOOLEAN NOT NULL,
                `name_dep` VARCHAR(255) NOT NULL

                ,PRIMARY KEY (`ID`)
                
            ) ENGINE=InnoDB;        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `employee1`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `isActive` BOOLEAN NOT NULL,
                `user_ID` INT NULL,
                `departm_ID` INT NULL

                ,PRIMARY KEY (`ID`)
                ,
                CONSTRAINT `fk_user_ID` FOREIGN KEY (`user_ID`) REFERENCES `user`(`ID`),
                CONSTRAINT `fk_departm_ID` FOREIGN KEY (`departm_ID`) REFERENCES `departm`(`ID`)
            ) ENGINE=InnoDB