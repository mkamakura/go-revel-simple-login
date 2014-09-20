package controllers

import (
    "github.com/revel/revel"
    "LoginSample/app/models"
    "encoding/json"
)

type ManagerCtrl struct {
    GorpController
}

func (c ManagerCtrl) parseManager() (models.Manager, error) {
    manager := models.Manager{}
    err := json.NewDecoder(c.Request.Body).Decode(&manager)
    return manager, err
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
