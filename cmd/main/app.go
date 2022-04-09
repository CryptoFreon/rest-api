package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"rest-api/pkg/logging"

	"rest-api/internal/config"
	"rest-api/internal/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//asd := os.Environ()
	//fmt.Println(asd)
	//goProjects, exist := os.LookupEnv("GOPROJECTS")
	//if exist {
	//	err := os.Chdir(goProjects + "/rest-api")
	//	if err != nil {
	//		log.Fatal("Not correct $GOPROJECTS", err)
	//	}
	//}

	//curDir, _ := os.Getwd()
	//fmt.Printf("Current Dir: %s\n", curDir)

	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)

}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
	} else {
		logger.Info("listen tcp")

		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	logger.Fatal(server.Serve(listener))
}
