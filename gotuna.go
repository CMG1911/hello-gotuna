package gotuna

import (
	"io/fs"
	"log"

	"github.com/gorilla/mux"
)

type App struct {
	Logger         *log.Logger
	Router         *mux.Router
	Static         fs.FS
	StaticPrefix   string
	ViewFiles      fs.FS
	ViewHelpers    []ViewHelperFunc
	Session        *Session
	UserRepository UserRepository
	Locale         Locale
}