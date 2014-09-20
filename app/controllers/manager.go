package controllers

import (
	_ "github.com/revel/revel"
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

func (c ManagerCtrl) Add() revel.Result {
	if manager, err := c.parseManager(); err != nil {
		return c.RenderText("Unable to parse the Manager from JSON.")
	} else {
		manager.Validate(c.Validate)
		if c.Validate.HasErrors() {
			return c.RenderText("You have error in your Manager.")
		} else {
			if err := c.Txn.Insert(&Manager); err != nil {
				return c.Render.Text(
					"Error inserting record into database!")
			} else {
				return c.RenderJson(manager)
			}
		}
	}
}
