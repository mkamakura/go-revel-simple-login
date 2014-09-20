package models

type Manager struct {
    ID int `db:"id" json:"id"`
    Name string `db:"name" json:"name"`
    Password string `db:"password" json:"password"`
}
