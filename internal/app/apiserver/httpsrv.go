package apiserver

import (
	"context"
	"github.com/kilimov/notificator/internal/app/business"
	"github.com/kilimov/notificator/internal/app/resources/users"
	"log"
	"net/http"
	"time"

	"github.com/kilimov/notificator/internal/app/resources"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

const compressLevel = 5

type HTTPServer struct {
	Address           string
	FilesDir          string
	CertFile, KeyFile *string
	IsTesting         bool

	userManager *business.UserManager

	idleConnsClosed chan struct{}
	masterCtx       context.Context
	version         string
}

func NewHTTPServer(ctx context.Context, opts Config, userManager *business.UserManager, version string) *HTTPServer {
	srv := &HTTPServer{
		Address:   opts.ListenAddr,
		FilesDir:  opts.FilesDir,
		IsTesting: opts.IsTesting,

		idleConnsClosed: make(chan struct{}),
		masterCtx:       ctx,
		version:         version,

		userManager: userManager,
	}

	if opts.CertFile != "" {
		srv.CertFile = &opts.CertFile
	}

	if opts.KeyFile != "" {
		srv.KeyFile = &opts.KeyFile
	}

	return srv
}

func (srv *HTTPServer) setupRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.NoCache)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.NewCompressor(compressLevel).Handler)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins(srv.IsTesting),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	filesRoute := "/files"
	r.Mount(filesRoute, resources.FilesResource{FilesDir: srv.FilesDir}.Routes())
	r.Mount("/swagger", resources.SwaggerResource{FilesPath: filesRoute}.Routes())
	r.Mount("/version", resources.VersionResource{Version: srv.version}.Routes())

	r.Mount("/api/v1/users", users.NewUserResource(srv.userManager).Routes())

	return r
}

func allowedOrigins(testing bool) []string {
	if testing {
		return []string{"*"}
	}

	return []string{}
}

func (srv *HTTPServer) Run() error {
	const (
		readTimeout  = 5 * time.Second
		writeTimeout = 30 * time.Second
	)

	s := &http.Server{
		Addr:         srv.Address,
		Handler:      chi.ServerBaseContext(srv.masterCtx, srv.setupRouter()),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	go srv.ListenCtxForGracefulTermination(s)
	log.Printf("[INFO] serving HTTP on \"%s\"", srv.Address)

	if srv.CertFile == nil && srv.KeyFile == nil {
		if err := s.ListenAndServe(); err != nil {
			return err
		}
	} else {
		if err := s.ListenAndServeTLS(*srv.CertFile, *srv.KeyFile); err != nil {
			return err
		}
	}

	return nil
}

func (srv *HTTPServer) ListenCtxForGracefulTermination(s *http.Server) {
	<-srv.masterCtx.Done()

	if err := s.Shutdown(srv.masterCtx); err != nil {
		log.Printf("[ERROR] HTTP server Shutdown: %v", err)
	}

	log.Println("Processed idle connections successfully before termination")
	close(srv.idleConnsClosed)
}

func (srv *HTTPServer) WaitForGracefulTermination() {
	<-srv.idleConnsClosed
}
