-- CREATE TABLE "accounts" (
--   "id" bigserial PRIMARY KEY,
--   "owner" varchar NOT NULL,
--   "balance" bigint NOT NULL,
--   "currency" varchar NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT 'now()'
-- );

-- CREATE TABLE "entries" (
--   "id" bigserial PRIMARY KEY,
--   "account_id" bigint NOT NULL,
--   "amount" bigint NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT 'now()'
-- );

-- CREATE TABLE "transfers" (
--   "id" bigserial PRIMARY KEY,
--   "source_account_id" bigint NOT NULL,
--   "target_account_id" bigint NOT NULL,
--   "amount" bigint NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT 'now()'
-- );

-- CREATE INDEX ON "accounts" ("owner");

-- CREATE INDEX ON "entries" ("account_id");

-- CREATE INDEX ON "transfers" ("source_account_id");

-- CREATE INDEX ON "transfers" ("target_account_id");

-- COMMENT ON COLUMN "entries"."amount" IS 'Positive or Negative';

-- COMMENT ON COLUMN "transfers"."amount" IS 'Positive Only';

-- ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

-- ALTER TABLE "transfers" ADD FOREIGN KEY ("source_account_id") REFERENCES "accounts" ("id");

-- ALTER TABLE "transfers" ADD FOREIGN KEY ("target_account_id") REFERENCES "accounts" ("id");

CREATE TYPE "carbon_type" AS ENUM (
  'Wood',
  'Inner',
  'Outer'
);

CREATE TYPE "grip_type" AS ENUM (
  'Penhold',
  'Handshake'
);

CREATE TYPE "status" AS ENUM (
  'sold',
  'not_sold'
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "source_account_id" bigint NOT NULL,
  "target_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "RacketsForSale" (
  "id" bigint UNIQUE,
  "price" bigint NOT NULL,
  "seller_id" bigint NOT NULL,
  "posted_time" TIMESTAMP NOT NULL DEFAULT 'now()',
  "status" status NOT NULL,
  "buyer_id" bigint
);

CREATE TABLE "RacketInventory" (
  "id" bigserial PRIMARY KEY,
  "carbon_type" carbon_type NOT NULL,
  "grip_type" grip_type NOT NULL,
  "owner_id" bigint,
  "isSelling" boolean NOT NULL DEFAULT false
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("source_account_id");

CREATE INDEX ON "transfers" ("target_account_id");

CREATE INDEX ON "transfers" ("source_account_id", "target_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'Positive or Negative';

COMMENT ON COLUMN "transfers"."amount" IS 'Positive Only';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("source_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("target_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "RacketsForSale" ADD FOREIGN KEY ("id") REFERENCES "RacketInventory" ("id");

ALTER TABLE "RacketsForSale" ADD FOREIGN KEY ("seller_id") REFERENCES "accounts" ("id");

ALTER TABLE "RacketsForSale" ADD FOREIGN KEY ("buyer_id") REFERENCES "accounts" ("id");

ALTER TABLE "RacketInventory" ADD FOREIGN KEY ("owner_id") REFERENCES "accounts" ("id");
