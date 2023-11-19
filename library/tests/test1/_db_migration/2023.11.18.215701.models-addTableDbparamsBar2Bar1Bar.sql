        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `dbparams`
            (
                `ID` INT NOT NULL AUTO_INCREMENT,
                `ParamName` VARCHAR(255) NOT NULL,
                `ParamValue` VARCHAR(255) NOT NULL

                ,PRIMARY KEY (`ID`)
                
            ) ENGINE=InnoDB;        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `bar2`
            (
                `Id` BIGINT NOT NULL AUTO_INCREMENT,
                `foo_id` BIGINT NULL,
                `unique_int` INT NULL,
                `nonunique_int` INT NULL

                ,PRIMARY KEY (`Id`)
                
            ) ENGINE=InnoDB;        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `bar1`
            (
                `Id` BIGINT NOT NULL AUTO_INCREMENT,
                `foo_id` BIGINT NULL,
                `unique_int` INT NULL,
                `nonunique_int` INT NULL

                ,PRIMARY KEY (`Id`)
                
            ) ENGINE=InnoDB;        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS `bar`
            (
                `Id` BIGINT NOT NULL AUTO_INCREMENT,
                `foo_id` BIGINT NULL,
                `unique_int` INT NULL,
                `nonunique_int` INT NULL

                ,PRIMARY KEY (`Id`)
                
            ) ENGINE=InnoDB