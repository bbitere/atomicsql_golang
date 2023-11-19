        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `DBParams`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `ParamName` VARCHAR(255) NOT NULL,
                `ParamValue` VARCHAR(255) NOT NULL

                ,PRIMARY KEY (`ID`)
                
            ) ENGINE=InnoDB;        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `statusRole`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `status_name` VARCHAR(255) NULL

                ,PRIMARY KEY (`ID`)
                
            ) ENGINE=InnoDB;        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `user_role`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `isActive` BOOLEAN NOT NULL,
                `roleName` VARCHAR(255) NOT NULL,
                `role_status_ID` INT NULL,
                `my_active` BOOLEAN NOT NULL

                ,PRIMARY KEY (`ID`)
                ,
                CONSTRAINT `fk_role_status_ID` FOREIGN KEY (`role_status_ID`) REFERENCES `statusRole`(`ID`)
            ) ENGINE=InnoDB;        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `user`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `userName` VARCHAR(255) NOT NULL,
                `UUID` VARCHAR(255) NOT NULL,
                `user_psw` VARCHAR(255) NOT NULL,
                `userRole_ID` INT NULL,
                `time1` TIMESTAMP NULL,
                `money` DOUBLE PRECISION NOT NULL

                ,PRIMARY KEY (`ID`)
                ,
                CONSTRAINT `fk_userRole_ID` FOREIGN KEY (`userRole_ID`) REFERENCES `user_role`(`ID`)
            ) ENGINE=InnoDB