CREATE TABLE "location" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "longitude" numeric,
  "latitude" numeric,
  "count" integer DEFAULT 0
);

CREATE INDEX ON "location" ("name");

COMMENT ON COLUMN "location"."name" IS 'name of city';

COMMENT ON COLUMN "location"."count" IS 'to count visited';
