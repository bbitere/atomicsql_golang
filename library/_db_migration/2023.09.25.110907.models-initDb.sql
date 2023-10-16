        
            /*------------------------------------------------------------------*/
            CREATE TABLE IF NOT EXISTS `DBParams`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `ParamName` VARCHAR(20000) NOT NULL,
                `ParamValue` VARCHAR(20000) NOT NULL

                ,PRIMARY KEY (`ID`)
            )
            TABLESPACE pg_default;
                ;        
            /*------------------------------------------------------------------*/
            CREATE TABLE IF NOT EXISTS `statusRole`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `status_name` VARCHAR(20000) NULL

                ,PRIMARY KEY (`ID`)
            )
            TABLESPACE pg_default;
                ;        
            /*------------------------------------------------------------------*/
            CREATE TABLE IF NOT EXISTS `user_role`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `isActive` BOOLEAN NOT NULL,
                `roleName` VARCHAR(20000) NOT NULL,
                `role_status_ID` INT NULL, CONSTRAINT `fk_role_status_ID` FOREIGN KEY (`role_status_ID`) REFERENCES `statusRole`(`ID`),
                `my_active` BOOLEAN NOT NULL

                ,PRIMARY KEY (`ID`)
            )
            TABLESPACE pg_default;
                ;        
            /*------------------------------------------------------------------*/
            CREATE TABLE IF NOT EXISTS `user`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `userName` VARCHAR(20000) NOT NULL,
                `UUID` VARCHAR(20000) NOT NULL,
                `user_psw` VARCHAR(20000) NOT NULL,
                `userRole_ID` INT NULL, CONSTRAINT `fk_userRole_ID` FOREIGN KEY (`userRole_ID`) REFERENCES `user_role`(`ID`),
                `time1` TIMESTAMP NULL,
                `money` DOUBLE PRECISION NOT NULL

                ,PRIMARY KEY (`ID`)
            )
            TABLESPACE pg_default;
                