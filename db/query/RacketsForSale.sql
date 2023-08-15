-- name: CreateRacketsForSale :one
INSERT INTO "RacketsForSale" (
  id,
  price,
  seller_id,
  status
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: ListRaquets :many
SELECT * FROM "RacketsForSale"
WHERE "price" BETWEEN $2 AND $3
  AND "grip_type" = $4
  AND "carbon_type" = $5
  AND "status" = 'Selling'
LIMIT $1;

-- name: ListAllRackets :many
SELECT * FROM "RacketsForSale"
WHERE "status" = 'Selling'
LIMIT $1;

-- name: UpdateAccountFeatures :one
UPDATE "RacketsForSale"
SET "price" = $1, "grip_type" = $2, "carbon_type" = $3
WHERE "id" = $4 AND "status" = 'Selling'
RETURNING *;

-- name: UpdateAccountStatus :one
UPDATE "RacketsForSale"
SET "status" = 'Sold'
WHERE "id" = $1
RETURNING *;

-- name: DeleteSellingRacket :exec
DELETE FROM accounts 
WHERE id = $1 and "status" = 'Selling';








