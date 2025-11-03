package models

type Note struct {
	ID       int    `db:"id"`
	Title    string `db:"title"`
	Content  string `db:"content"`
	Language string `db:"language"`
}
