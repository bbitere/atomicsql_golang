        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."DBParams"
            (
                "ID" serial PRIMARY KEY NOT NULL,
                "ParamName" VARCHAR NOT NULL,
                "ParamValue" VARCHAR NOT NULL
            )
            TABLESPACE pg_default;
                ;        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."statusRole"
            (
                "ID" serial PRIMARY KEY NOT NULL,
                "status_name" VARCHAR NULL
            )
            TABLESPACE pg_default;
                ;        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."user_role"
            (
                "ID" serial PRIMARY KEY NOT NULL,
                "isActive" BOOLEAN NOT NULL,
                "roleName" VARCHAR NOT NULL,
                "role_status_ID" integer NULL CONSTRAINT "fk_role_status_ID" REFERENCES public."statusRole" ("ID"),
                "my_active" BOOLEAN NOT NULL
            )
            TABLESPACE pg_default;
                ;        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."user"
            (
                "ID" serial PRIMARY KEY NOT NULL,
                "userName" VARCHAR NOT NULL,
                "UUID" VARCHAR NOT NULL,
                "user_psw" VARCHAR NOT NULL,
                "userRole_ID" integer NULL CONSTRAINT "fk_userRole_ID" REFERENCES public."user_role" ("ID"),
                "time1" TIMESTAMP NULL,
                "money" double precision NOT NULL
            )
            TABLESPACE pg_default;
                