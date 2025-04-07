-- Create enum type "user_type"
CREATE TYPE "user_type" AS ENUM ('REGULAR', 'REFERENCE');
-- Create "users" table
CREATE TABLE "users" (
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
-- Create index "idx_users_gh_id" to table: "users"
CREATE UNIQUE INDEX "idx_users_gh_id" ON "users" ("gh_id");
-- Create index "idx_users_username" to table: "users"
CREATE UNIQUE INDEX "idx_users_username" ON "users" ("username");
-- Create "user_followers" table
CREATE TABLE "user_followers" (
  "user_id" uuid NOT NULL,
  "follower_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "follower_id"),
  CONSTRAINT "fk_user_followers_followers" FOREIGN KEY ("follower_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_user_followers_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
