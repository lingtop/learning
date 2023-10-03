package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"template/service/database/postgresql"
	"template/service/external/integration"
	"template/service/global_variable"
	"template/service/interface/http"
	"template/service/logger"
	"template/service/model"
	"template/service/util"

	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
)

// Global Variable
var wg sync.WaitGroup
var configFile string

func initConfig() {
	// Init viper
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}

	// Read Config
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "unable to read config: %v\n", err)
		os.Exit(1)
	}

	// Parse Json Decimal to Float
	decimal.MarshalJSONWithoutQuotes = true
}

func initPostgresql() {
	wg.Add(1)
	postgresql.InitDatabase()
}

func shutdownPostgresql() {
	postgresql.ShutdownDatabase()
	wg.Done()
}

func initRedis() {
	// wg.Add(1)
	// redis.Init()
}

func shutdownRedis() {
	// redis.ShutDown()
	// wg.Done()
}

func initListenInterface() {
	wg.Add(1)
	go func() {
		http.InitHttpServer()
		wg.Done()
	}()
}

func initListenOsSignal() {
	wg.Add(1)
	go func() {
		var count int
		chanOsSignal := make(chan os.Signal, 2)
		signal.Notify(chanOsSignal, syscall.SIGTERM, os.Interrupt)

		go func() {
			for getSignal := range chanOsSignal {
				if getSignal == os.Interrupt || getSignal == syscall.SIGTERM {
					count++
					if count == 2 {
						logger.Logger.Info("Forcefully exiting")
						os.Exit(1)
					}

					go func() {
						shutdownPostgresql()
					}()

					go func() {
						http.ShutdownHttpServer()
					}()

					go func() {
						shutdownRedis()
					}()

					logger.Logger.Info("Signal SIGKILL caught. shutting down")
					logger.Logger.Info("Catching SIGKILL one more time will forcefully exit")

					wg.Done()
				}
			}
			close(chanOsSignal)
		}()
	}()
}

func initTimezone() {
	loc, err := time.LoadLocation(global_variable.TimeZone)
	if err != nil {
		logger.Logger.Error("unable to set timezone cuz: %s", err)
		os.Exit(0)
	}
	time.Local = loc
}

func initComponent() {
	// Init Global Variable
	global_variable.InitVariable()

	// Init Utility Function
	util.InitUtil()

	// Init Model
	model.InitModel()

	// Init Integration
	integration.InitIntegrationInfo()
}
