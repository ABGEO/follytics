-- Create enum type "user_type"
CREATE TYPE "user_type" AS ENUM ('REGULAR', 'REFERENCE');
-- Create "user" table
CREATE TABLE "user" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "type" "user_type" NULL DEFAULT 'REGULAR',
  "gh_id" bigint NULL,
  "username" text NULL,
  "name" text NULL,
  "email" text NULL,
  "avatar" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_user_gh_id" to table: "user"
CREATE UNIQUE INDEX "idx_user_gh_id" ON "user" ("gh_id");
-- Create index "idx_user_username" to table: "user"
CREATE UNIQUE INDEX "idx_user_username" ON "user" ("username");
-- Create "user_followers" table
CREATE TABLE "user_followers" (
  "user_id" uuid NOT NULL,
  "follower_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "follower_id"),
  CONSTRAINT "fk_user_followers_followers" FOREIGN KEY ("follower_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_user_followers_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
