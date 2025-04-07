-- Create "job_states" table
CREATE TABLE "job_states" (
  "job_name" text NOT NULL,
  "attributes" jsonb NULL,
  PRIMARY KEY ("job_name")
);
