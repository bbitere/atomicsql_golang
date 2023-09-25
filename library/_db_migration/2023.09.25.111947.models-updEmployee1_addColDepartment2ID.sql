
            ALTER TABLE public."employee1"
            ADD COLUMN "departm2_ID" integer NULL CONSTRAINT "fk_departm2_ID" REFERENCES public."departm" ("ID")