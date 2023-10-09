
                ALTER TABLE `project`
                ADD CONSTRAINT `fk_projStatus_ID` FOREIGN KEY (`projStatus_ID`) REFERENCES `projstatus` (`ID`);