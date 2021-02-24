package resources

import (
	_ "github.com/kilimov/notificator/api"
	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger"
	"path/filepath"
)

// SwaggerResource для размещения API документации
type SwaggerResource struct {
	FilesPath string
}

func (sr SwaggerResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/*", httpSwagger.Handler(
		httpSwagger.URL(filepath.Join(sr.FilesPath, "swagger.json")),
	))
	return r
}
