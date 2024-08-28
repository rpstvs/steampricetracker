-- name: CreateItem :one
INSERT INTO Items (
        classid,
        ItemName,
        ImageUrl,
        DayChange,
        WeekChange
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetItemsIds :many
SELECT classid
FROM Items;
-- name: GetItemByName :one
SELECT classid
FROM Items
WHERE itemname = $1;
-- name: UpdateDailyChange :exec
UPDATE Items
SET DayChange = $1
WHERE classid = $2;
-- name: UpdateWeeklyChange :exec
UPDATE Items
SET WeekChange = $1
WHERE classid = $2;