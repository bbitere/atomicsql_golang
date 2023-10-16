        
            /*------------------------------------------------------------------*/
            CREATE TABLE IF NOT EXISTS `projstatus`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `name` VARCHAR(20000) NOT NULL

                ,PRIMARY KEY (`ID`)
            )
            TABLESPACE pg_default;
                ;        
            /*------------------------------------------------------------------*/
            CREATE TABLE IF NOT EXISTS `project`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `name` VARCHAR(20000) NOT NULL

                ,PRIMARY KEY (`ID`)
            )
            TABLESPACE pg_default;
                