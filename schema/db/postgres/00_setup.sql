
-- migrate up
CREATE SCHEMA IF NOT EXISTS "auth";


CREATE OR REPLACE FUNCTION "auth".function_updated_at()
  RETURNS TRIGGER AS $$
  BEGIN
   NEW.updated_at = now();
   RETURN NEW;
  END;
  $$ LANGUAGE 'plpgsql';

CREATE TABLE "auth"."user" (
	id_user 		    TEXT PRIMARY KEY,
	first_name		    TEXT NOT NULL,
	last_name			TEXT NOT NULL,
	email 				TEXT UNIQUE NOT NULL,
	password_hash   	TEXT NOT NULL,
	refresh_token		TEXT NOT NULL,
	active				BOOLEAN DEFAULT TRUE NOT NULL,
	created_at			TIMESTAMP DEFAULT NOW(),
	updated_at			TIMESTAMP DEFAULT NOW()
);

CREATE TRIGGER trigger_user_updated_at BEFORE UPDATE
  ON "auth"."user" FOR EACH ROW EXECUTE PROCEDURE "auth".function_updated_at();


-- migrate down
DROP TRIGGER trigger_user_updated_at ON auth."user"

DROP TABLE "auth"."user";

DROP FUNCTION "auth".function_updated_at;

DROP SCHEMA "auth";
