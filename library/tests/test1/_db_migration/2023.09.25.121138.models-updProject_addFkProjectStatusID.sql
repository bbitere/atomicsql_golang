
                ALTER TABLE public."project"
                ADD CONSTRAINT "fk_projStatus_ID" FOREIGN KEY ("projStatus_ID") REFERENCES public."projstatus" ("ID");