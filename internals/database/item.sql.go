// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: item.sql

package database

import (
	"context"
)

const createItem = `-- name: CreateItem :one
INSERT INTO Items (
        classid,
        ItemName,
        ImageUrl,
        DayChange,
        WeekChange
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING classid, itemname, daychange, weekchange, imageurl
`

type CreateItemParams struct {
	Classid    int64
	Itemname   string
	Imageurl   string
	Daychange  float64
	Weekchange float64
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.Classid,
		arg.Itemname,
		arg.Imageurl,
		arg.Daychange,
		arg.Weekchange,
	)
	var i Item
	err := row.Scan(
		&i.Classid,
		&i.Itemname,
		&i.Daychange,
		&i.Weekchange,
		&i.Imageurl,
	)
	return i, err
}

const getItemByName = `-- name: GetItemByName :one
SELECT classid
FROM Items
WHERE itemname = $1
`

func (q *Queries) GetItemByName(ctx context.Context, itemname string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getItemByName, itemname)
	var classid int64
	err := row.Scan(&classid)
	return classid, err
}

const getItemsIds = `-- name: GetItemsIds :many
SELECT classid
FROM Items
`

func (q *Queries) GetItemsIds(ctx context.Context) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getItemsIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var classid int64
		if err := rows.Scan(&classid); err != nil {
			return nil, err
		}
		items = append(items, classid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDailyChange = `-- name: UpdateDailyChange :exec
UPDATE Items
SET DayChange = $1
WHERE classid = $2
`

type UpdateDailyChangeParams struct {
	Daychange float64
	Classid   int64
}

func (q *Queries) UpdateDailyChange(ctx context.Context, arg UpdateDailyChangeParams) error {
	_, err := q.db.ExecContext(ctx, updateDailyChange, arg.Daychange, arg.Classid)
	return err
}

const updateWeeklyChange = `-- name: UpdateWeeklyChange :exec
UPDATE Items
SET WeekChange = $1
WHERE classid = $2
`

type UpdateWeeklyChangeParams struct {
	Weekchange float64
	Classid    int64
}

func (q *Queries) UpdateWeeklyChange(ctx context.Context, arg UpdateWeeklyChangeParams) error {
	_, err := q.db.ExecContext(ctx, updateWeeklyChange, arg.Weekchange, arg.Classid)
	return err
}
