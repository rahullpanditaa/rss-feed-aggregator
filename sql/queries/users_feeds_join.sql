-- name: GetAllFeedsWithCreatorUsername :many
SELECT "feeds"."name", "feeds"."url", "users"."name" AS "user_name"
FROM "feeds"
JOIN "users"
ON "feeds"."user_id" = "users"."id";