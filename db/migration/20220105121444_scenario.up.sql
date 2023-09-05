CREATE TABLE IF NOT EXISTS "users"(
	"id" BIGSERIAL NOT NULL PRIMARY KEY,
	"login"   VARCHAR NOT NULL,
	"password_hash"  VARCHAR NOT NULL 
);

CREATE TABLE IF NOT EXISTS "scenarios"(
	"id"   BIGSERIAL NOT NULL PRIMARY KEY,
	"name"  VARCHAR NOT NULL,
	"owner_id" INT NOT NULL,
	"steps_count" INT NOT NULL,
	"deleted_at" TIME,
	"data" JSONB NOT NULL,
	FOREIGN KEY ("owner_id")  REFERENCES "users" ("id")
);
CREATE INDEX ON "scenarios"("owner_id");

CREATE TABLE IF NOT EXISTS "simulations"(
	"id"   BIGSERIAL NOT NULL PRIMARY KEY,
	"scenario_id" INT NOT NULL,
	"name"  VARCHAR NOT NULL,
	"owner_id" INT NOT NULL,
	FOREIGN KEY ("scenario_id")  REFERENCES "scenarios" ("id"),
	FOREIGN KEY ("owner_id")  REFERENCES "users" ("id")
);
CREATE INDEX ON "simulations"("owner_id");
CREATE INDEX ON "simulations"("scenario_id");

CREATE TABLE IF NOT EXISTS "steps"(
	"id"   BIGSERIAL NOT NULL PRIMARY KEY,
	"simulation_id" INT NOT NULL,
	"step" INT not NULL,
	"data" JSONB,
	FOREIGN KEY ("simulation_id")  REFERENCES "simulations" ("id")
);

CREATE INDEX ON "steps"("simulation_id");

CREATE TABLE IF NOT EXISTS "events"(
	"id"   BIGSERIAL NOT NULL PRIMARY KEY,
	"name"  VARCHAR NOT NULL,
	"description"  VARCHAR NOT NULL
);