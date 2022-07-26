package main

import (
	"os"

	"github.com/robfig/cron/v3"
	"github.com/samber/lo"

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
	// Dependencies
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

	// Persistance Layer

	appRepo := &sqlite.AppRepo{
		Logger: logger.Scoped("app-repo"),
	}
	envRepo := &sqlite.EnvRepo{
		Logger: logger.Scoped("env-repo"),
	}
	routeRepo := &sqlite.RouteRepo{
		Logger: logger.Scoped("route-repo"),
	}
	pluginRepo := &sqlite.PluginRepo{
		Logger: logger.Scoped("plugin-repo"),
	}
	runtimeRepo := docker.NewRuntimeRepo(
		logger.Scoped("runtime-repo"),
		dockerClient,
	)
	runtimeServiceRepo := docker.NewRuntimeServiceRepo(
		logger.Scoped("runtime-service-repo"),
		dockerClient,
		certResolverName,
	)
	runtimeImageRepo := docker.NewRuntimeImageRepo(
		logger.Scoped("runtime-image-repo"),
		dockerClient,
	)
	runtimeNodeRepo := docker.NewRuntimeNodeRepo(
		logger.Scoped("runtime-node-repo"),
		dockerClient,
	)
	runtimeTaskRepo := docker.NewRuntimeTaskRepo(
		logger.Scoped("runtime-task-repo"),
		dockerClient,
	)

	// Service Layer

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
		AppRepo:          appRepo,
		PluginRepo:       pluginRepo,
		CertResolverName: certResolverName,
		RuntimeService:   runtimeService,
		AppService:       appService,
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

	// Jobs

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

	go scheduleJobs(jobs, logger.Scoped("cron"))

	// Controllers

	server := graphql.NewServer(logger.Scoped("gql-server"), db, resolver, ACCESS_TOKEN)
	server.ServeGraphql()
}

func scheduleJobs(jobs []scheduledJob, logger *fmt.Logger) {
	c := cron.New(cron.WithChain(
		cron.Recover(logger),
	))
	logger.I("Scheduled Jobs:")
	lo.ForEach(jobs, func(job scheduledJob, _ int) {
		c.AddFunc(job.expression, job.job)
		logger.I(" - %s (%s)", job.name, job.expression)
	})
	c.Start()
}
