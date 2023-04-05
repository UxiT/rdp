package bootstrap

import (
	"github.com/UxiT/rdp/db"
)

type Application struct {
	Env *Env
	DB  db.Database
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DB = NewDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	closeDatabaseConnection(&app.DB)
}
