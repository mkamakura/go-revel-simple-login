package controllers

import (
    "crypto/md5"
    "encoding/hex"
    "github.com/revel/revel"
    "LoginSample/app/models"
    "LoginSample/app/routes"
)

type App struct {
    GorpController
}

func (c App) Index() revel.Result {
    if c.connected() != nil {
        // return c.Redirect(routes.App.Index())
    }
    c.Flash.Error("Please log in first")

    headerTitle := "タイトル"
    return c.Render(headerTitle)
}

func (c App) connected() *models.Manager {
    if c.RenderArgs["manager"] != nil {
        return c.RenderArgs["manager"].(*models.Manager)
    }
    if name, ok := c.Session["manager"]; ok {
        return c.findManagerByName(name)
    }

    return nilå
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
        hashPassword := getMD5Hash(password)
        if manager.Password == hashPassword {
            c.Session["manager"] = name
            c.Session.SetNoExpiration()
            c.Flash.Success("Welcome, " + name)
            return c.Redirect(routes.App.Index())
        }
    }

    c.Flash.Error("Login failed")
    return c.Redirect(routes.App.Index())
}

func (c App) findManagerByName(name string) *models.Manager {
    manager, err := c.Txn.Select(models.Manager{},
        `SELECT * FROM manager WHERE name = ?`, name)

    if err != nil {
        revel.ERROR.Println(err)
        return nil
    }
    if len(manager) == 0 {
        revel.ERROR.Printf("manager not found[%s]", name)
        return nil
    }

    return manager[0].(*models.Manager)
}

func getMD5Hash(password string) string {
    hasher := md5.New()
    hasher.Write([]byte(password))
    return hex.EncodeToString(hasher.Sum(nil))
}
