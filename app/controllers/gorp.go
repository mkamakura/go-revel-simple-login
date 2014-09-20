package controllers

import (
    "github.com/coopernurse/gorp"
    "database/sql"
    "github.com/revel/revel"
	_ "github.com/go-sql-driver/mysql"
	_ "fmt"
	_ "strings"
	"LoginSample/app/models"
)

var (
    Dbm *gorp.DbMap
)

type GorpController struct {
    *revel.Controller
    Txn *gorp.Transaction
}

var InitDB func() = func() {
    connectionString := getConnectionString()
    if db, err := sql.Open("mysql", connectionString); err != nil {
        revel.ERROR.Fatal(err)
    } else {
		Dbm = &gorp.DbMap {
			Db: db,
			Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	}

	defineManagerTable(Dbm)
	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		revel.ERROR.Fatal(err)
	}
}

func defineManagerTable(dbm *gorp.DbMap) {
	t := dbm.AddTable(models.Manager{}).SetKeys(true, "id")
	t.ColMap("name").SetMaxSize(25)
}


func (c *GorpController) Begin() revel.Result {
    txn, err := Dbm.Begin()
    if err != nil {
        panic(err)
    }
    c.Txn = txn
    return nil
}

func (c *GorpController) Commit() revel.Result {
    if c.Txn == nil {
        return nil
    }
    if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}

func (c *GorpController) Rollback() revel.Result {
    if c.Txn == nil {
        return nil
    }
    if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}
