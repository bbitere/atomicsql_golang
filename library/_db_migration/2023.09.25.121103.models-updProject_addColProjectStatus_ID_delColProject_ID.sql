
            ALTER TABLE public."project"
            ADD COLUMN "projStatus_ID" INT NOT NULL;
            ALTER TABLE public."project"
            DROP COLUMN "proj_ID"