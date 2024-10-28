package app

import (
	"contactApp/repository"
	"contactApp/utils/log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	sync.Mutex
	Name       string
	Router     *mux.Router
	DB         *gorm.DB
	Log        log.Logger
	Server     *http.Server
	WG         *sync.WaitGroup
	Repository *repository.Repository
}

func newApp(name string,
	router *mux.Router,
	db *gorm.DB,
	log log.Logger,
	server *http.Server,
	wg *sync.WaitGroup,
	repository *repository.Repository,
) *App {
	return &App{
		Name:       name,
		DB:         db,
		Log:        log,
		WG:         wg,
		Repository: repository,
	}
}

func NewDBConnection(log log.Logger) *gorm.DB {
	db, err := gorm.Open("mysql", "root:Bank1mbha!Bank1mbha!@StructureApp?charsetutf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return db
}

func (app *App) initialiseRouter() {
	app.Log.Info(app.Name + " App Route initialising")
	app.Router = mux.NewRouter().StrictSlash()
	app.Router = app.Router.PathPrefix("/api/v1/contactapp")
}

func initialiseServer(){
	headers :=handlers.AllowedHeaders
}