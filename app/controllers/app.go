package controllers

import "github.com/revel/revel"

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {
    greeting := "Monster Hunter 4G Refarence"
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

    return c.Render(name)
}
