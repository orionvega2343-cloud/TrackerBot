package models

type Habits struct {
	Id     int    `db:"id"`
	UserId int64  `db:"user_id"`
	Title  string `db:"title"`
}
