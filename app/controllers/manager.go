package controllers

import (
    "github.com/revel/revel"
    "LoginSample/app/models"
)

type ManagerCtrl struct {
    GorpController
}

func (c ManagerCtrl) findManagerByName(name string) revel.Result {
    manager := new(models.Manager)
    err := c.Txn.SelectOne(manager,
        `SELECT * FROM Manager WHERE name = ?`, name)
    if err != nil {
        return c.RenderText("Error. Manager doesn't exist.")
    }
    return c.RenderJson(manager)
}
