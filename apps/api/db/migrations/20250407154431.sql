-- Create enum type "event_type"
CREATE TYPE "event_type" AS ENUM ('FOLLOW', 'UNFOLLOW');
-- Create "event" table
CREATE TABLE "event" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "type" "event_type" NULL,
  "user_id" uuid NULL,
  "reference_user_id" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_event_reference_user" FOREIGN KEY ("reference_user_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_event_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_event_reference_user_id" to table: "event"
CREATE INDEX "idx_event_reference_user_id" ON "event" ("reference_user_id");
-- Create index "idx_event_user_id" to table: "event"
CREATE INDEX "idx_event_user_id" ON "event" ("user_id");
