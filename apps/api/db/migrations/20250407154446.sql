-- Create "job_state" table
CREATE TABLE "job_state" (
  "job_name" text NOT NULL,
  "attributes" jsonb NULL,
  PRIMARY KEY ("job_name")
);
