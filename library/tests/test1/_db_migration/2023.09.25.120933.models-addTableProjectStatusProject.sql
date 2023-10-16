        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."projstatus"
            (
                "ID" serial PRIMARY KEY NOT NULL,
                "name" VARCHAR NOT NULL
            )
            TABLESPACE pg_default;
                ;        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."project"
            (
                "ID" serial PRIMARY KEY NOT NULL,
                "name" VARCHAR NOT NULL
            )
            TABLESPACE pg_default;
                