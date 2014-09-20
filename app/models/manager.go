package models

type Manager struct {
	ID int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
}

