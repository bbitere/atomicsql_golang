        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."dbparams"
            (
                "ID" serial PRIMARY KEY NOT NULL,
                "ParamName" VARCHAR NOT NULL,
                "ParamValue" VARCHAR NOT NULL
            )
            TABLESPACE pg_default;
                ;        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."bar2"
            (
                "Id" bigserial NOT NULL AUTO_INCREMENT,
                "foo_id" BIGINT NULL,
                "unique_int" INT NULL,
                "nonunique_int" INT NULL
            )
            TABLESPACE pg_default;
                ;        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."bar1"
            (
                "Id" bigserial NOT NULL AUTO_INCREMENT,
                "foo_id" BIGINT NULL,
                "unique_int" INT NULL,
                "nonunique_int" INT NULL
            )
            TABLESPACE pg_default;
                ;        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS public."bar"
            (
                "Id" bigserial NOT NULL AUTO_INCREMENT,
                "foo_id" BIGINT NULL,
                "unique_int" INT NULL,
                "nonunique_int" INT NULL
            )
            TABLESPACE pg_default;
                