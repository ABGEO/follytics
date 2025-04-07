-- Create enum type "event_type"
CREATE TYPE "event_type" AS ENUM ('FOLLOW', 'UNFOLLOW');
-- Create "events" table
CREATE TABLE "events" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "type" "event_type" NULL,
  "user_id" uuid NULL,
  "reference_user_id" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_events_reference_user" FOREIGN KEY ("reference_user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_events_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_events_reference_user_id" to table: "events"
CREATE INDEX "idx_events_reference_user_id" ON "events" ("reference_user_id");
-- Create index "idx_events_user_id" to table: "events"
CREATE INDEX "idx_events_user_id" ON "events" ("user_id");
