-- name: GetAllFeedsWithUser :many
SELECT "feeds"."name", "feeds"."url", "users"."name"
FROM "feeds"
JOIN "users"
ON "feeds"."user_id" = "users"."id";