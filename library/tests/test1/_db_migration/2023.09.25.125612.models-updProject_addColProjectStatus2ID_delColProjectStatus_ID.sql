
            ALTER TABLE public."project"
            ADD COLUMN "projStatus2_ID" integer NULL CONSTRAINT "fk_projStatus2_ID" REFERENCES public."projstatus" ("ID");
            ALTER TABLE public."project"
            DROP COLUMN "projStatus_ID"