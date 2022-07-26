package main

import (
	"os"

	"github.com/robfig/cron/v3"
	"github.com/samber/lo"

	"github.com/aklinker1/miasma/internal/server"
	cron2 "github.com/aklinker1/miasma/internal/server/cron"
	"github.com/aklinker1/miasma/internal/server/docker"
	"github.com/aklinker1/miasma/internal/server/fmt"
	"github.com/aklinker1/miasma/internal/server/graphql"
	"github.com/aklinker1/miasma/internal/server/services"
	"github.com/aklinker1/miasma/internal/server/sqlite"
)

// Compile time variables
var (
	VERSION    string
	BUILD      string
	BUILD_HASH string
	BUILD_DATE string
)

// Environment Variables
var (
	ACCESS_TOKEN      = os.Getenv("ACCESS_TOKEN")
	AUTO_UPGRADE_CRON = os.Getenv("AUTO_UPGRADE_CRON")
)

// Other constants
var (
	databasePath           = "/data/miasma/apps.db"
	certResolverName       = "miasmaresolver"
	defaultAutoUpgradeCron = "@daily"
)

type scheduledJob struct {
	name       string
	expression string
	job        func()
}

func main() {
	logger := &fmt.Logger{}

	db := sqlite.NewDB(databasePath, logger.Scoped("db"))
	err := db.Open()
	if err != nil {
		logger.E("Failed to open database: %v", err)
		os.Exit(1)
	}

	dockerClient, err := docker.NewDefaultClient()
	if err != nil {
		logger.E("Failed to initialize docker runtime: %v", err)
		os.Exit(1)
	}

	panic(server.NewNotImplementedError("Implement repos"))
	var appRepo server.AppRepo
	var envRepo server.EnvRepo
	var routeRepo server.RouteRepo
	var pluginRepo server.PluginRepo
	var runtimeRepo server.RuntimeRepo
	var runtimeServiceRepo server.RuntimeServiceRepo
	var runtimeImageRepo server.RuntimeImageRepo
	var runtimeNodeRepo server.RuntimeNodeRepo
	var runtimeTaskRepo server.RuntimeTaskRepo

	appService := &services.AppService{
		DB:                 db,
		Logger:             logger.Scoped("app-service"),
		AppRepo:            appRepo,
		RouteRepo:          routeRepo,
		EnvRepo:            envRepo,
		PluginRepo:         pluginRepo,
		RuntimeServiceRepo: runtimeServiceRepo,
	}
	runtimeService := &services.RuntimeService{
		DB:                 db,
		Logger:             logger.Scoped("runtime-service"),
		AppRepo:            appRepo,
		RouteRepo:          routeRepo,
		EnvRepo:            envRepo,
		PluginRepo:         pluginRepo,
		RuntimeServiceRepo: runtimeServiceRepo,
		AppService:         appService,
	}
	pluginService := &services.PluginService{
		DB:               db,
		Logger:           logger.Scoped("plugin-service"),
		PluginRepo:       pluginRepo,
		CertResolverName: certResolverName,
		RuntimeService:   runtimeService,
	}

	resolver := &graphql.Resolver{
		Version: VERSION,
		DB:      db,
		Logger:  logger.Scoped("resolver"),

		AppRepo:            appRepo,
		RouteRepo:          routeRepo,
		EnvRepo:            envRepo,
		PluginRepo:         pluginRepo,
		RuntimeServiceRepo: runtimeServiceRepo,
		RuntimeRepo:        runtimeRepo,
		RuntimeNodeRepo:    runtimeNodeRepo,
		RuntimeTaskRepo:    runtimeTaskRepo,
		RuntimeImageRepo:   runtimeImageRepo,

		AppService:     appService,
		PluginService:  pluginService,
		RuntimeService: runtimeService,
	}
	server := graphql.NewServer(logger.Scoped("gql-server"), db, resolver, ACCESS_TOKEN)

	pollingUpdater := cron2.PollingUpgrader{
		DB:               db,
		Logger:           logger.Scoped("polling-upgrader"),
		AppRepo:          appRepo,
		RuntimeImageRepo: runtimeImageRepo,
		PluginRepo:       pluginRepo,
		RuntimeService:   runtimeService,
	}
	jobs := []scheduledJob{
		{
			name:       "Auto Upgrade",
			expression: lo.Ternary(AUTO_UPGRADE_CRON == "", defaultAutoUpgradeCron, AUTO_UPGRADE_CRON),
			job:        pollingUpdater.Cron,
		},
	}

	go scheduleJobs(jobs, logger)
	server.ServeGraphql()
}

func scheduleJobs(jobs []scheduledJob, logger *fmt.Logger) {
	c := cron.New(cron.WithChain(
		cron.Recover(logger.Scoped("cron")),
	))
	logger.I("Scheduled Jobs:")
	lo.ForEach(jobs, func(job scheduledJob, _ int) {
		c.AddFunc(job.expression, job.job)
		logger.I(" - %s (%s)", job.name, job.expression)
	})
	c.Start()
}
