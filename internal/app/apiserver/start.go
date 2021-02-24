// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package apiserver

import (
	"context"
	"fmt"
	"github.com/kilimov/notificator/internal/app/database"
	"github.com/kilimov/notificator/internal/app/database/drivers"
	"github.com/kilimov/notificator/internal/app/log"
	"github.com/pkg/errors"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/logutils"
)

// Start is a command to start new api server.
func Start(version string) {
	fmt.Printf("notificator %s\n", version)

	appCtx, cancelAppCtx := context.WithCancel(context.Background())
	defer cancelAppCtx()
	go listenSystemSignals(cancelAppCtx)

	opts := ConfigWithParsedFlags()
	setupLog(opts.InDebugMode)

	ds, err := setupDS(*opts)
	if err != nil {
		stdlog.Println(err)
		return
	}
	defer log.ErrorFnPrintln(func() error { return ds.Close(context.Background()) })

	app := NewHTTPServer(appCtx, *opts, version)
	if err := app.Run(); err != nil {
		stdlog.Println(err)
		return
	}

	app.WaitForGracefulTermination()
	stdlog.Printf("[INFO] process terminated")
}

// setupLog sets up log levels and logs output.
func setupLog(inDebugMode bool) {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("INFO"),
		Writer:   os.Stdout,
	}

	stdlog.SetFlags(stdlog.Ldate | stdlog.Ltime)

	if inDebugMode {
		stdlog.SetFlags(stdlog.Ldate | stdlog.Ltime | stdlog.Lmicroseconds | stdlog.Lshortfile)
		filter.MinLevel = "DEBUG"
	}

	stdlog.SetOutput(filter)
}

func setupDS(opts Config) (drivers.DataStore, error) {
	ds, err := database.New(drivers.DataStoreConfig{
		URL:           opts.DSURL,
		DataStoreName: opts.DSName,
		DataBaseName:  opts.DSDB,
	})
	if err != nil {
		errMsg := fmt.Sprintf("[ERROR] cannot create datastore %s: %v", opts.DSName, err)
		return nil, errors.New(errMsg)
	}

	if err := ds.Connect(); err != nil {
		errMsg := fmt.Sprintf("[ERROR] cannot connect to datastore %s: %v", opts.DSName, err)
		return nil, errors.New(errMsg)
	}

	stdlog.Printf("[INFO] connected to %s", ds.Name())

	return ds, nil
}

func listenSystemSignals(cancelAppCtx context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	stdlog.Print("[WARN] interrupt signal")
	cancelAppCtx()
}
