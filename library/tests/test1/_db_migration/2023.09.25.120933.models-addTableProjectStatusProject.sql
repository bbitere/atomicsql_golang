        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `projstatus`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `name` VARCHAR(255) NOT NULL

                ,PRIMARY KEY (`ID`)
                
            ) ENGINE=InnoDB;        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `project`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `name` VARCHAR(255) NOT NULL

                ,PRIMARY KEY (`ID`)
                
            ) ENGINE=InnoDB