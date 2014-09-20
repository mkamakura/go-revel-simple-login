package controllers

import (
	"github.com/revel/revel"
	"LoginSample/app/models"
	"LoginSample/app/routes"
	"fmt"
)

type App struct {
	GorpController
}

func (c App) Index() revel.Result {
    greeting := "Hello World"
    return c.Render(greeting)
}

func (c App) Login(name string, password string) revel.Result {
    c.Validation.Required(name).Message("Your name is required!")
    c.Validation.Required(password).Message("Your password is required!")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Index)
	}

	manager := c.findManagerByName(name)

	if manager != nil {

	}

	c.Flash.Out["name"] = name
	c.Flash.Error("Login failed")
	return c.Redirect(routes.App.Index())
}

func (c App) findManagerByName(name string) *models.Manager {
	manager, err := c.Txn.Select(models.Manager{},
		`SELECT * FROM Manager WHERE name = ?`, name)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(manager) == 0 {
		return nil
	}

	return manager[0].(*models.Manager)
}
