        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."departm"
            (
                "ID" serial PRIMARY KEY NOT NULL,
                "isActive" BOOLEAN NOT NULL,
                "name_dep" VARCHAR NOT NULL
            )
            TABLESPACE pg_default;
                ;        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."employee1"
            (
                "ID" serial PRIMARY KEY NOT NULL,
                "isActive" BOOLEAN NOT NULL,
                "user_ID" integer NULL CONSTRAINT "fk_user_ID" REFERENCES public."user" ("ID"),
                "departm_ID" integer NULL CONSTRAINT "fk_departm_ID" REFERENCES public."departm" ("ID")
            )
            TABLESPACE pg_default;
                