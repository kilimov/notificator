package resources

import (
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

// FilesResource для раздачи статичных файлов
type FilesResource struct {
	FilesDir string
}

func (fr FilesResource) Routes() chi.Router {
	r := chi.NewRouter()
	filesRoot := http.Dir(fr.FilesDir)

	NewFileServer(r, "/", filesRoot)

	return r
}

// NewFileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func NewFileServer(r chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		ctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
