package main

import (
	"os"

	"github.com/robfig/cron/v3"
	"github.com/samber/lo"

	cron2 "github.com/aklinker1/miasma/internal/server/cron"
	"github.com/aklinker1/miasma/internal/server/docker"
	"github.com/aklinker1/miasma/internal/server/fmt"
	"github.com/aklinker1/miasma/internal/server/graphql"
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

	db := sqlite.NewDB(databasePath, logger)
	err := db.Open()
	if err != nil {
		logger.E("Failed to open database: %v", err)
		os.Exit(1)
	}

	runtime, err := docker.NewRuntimeService(logger, certResolverName)
	apps := sqlite.NewAppService(db, runtime, logger)
	env := sqlite.NewEnvService(db, runtime, logger)
	routes := sqlite.NewRouteService(db, logger)
	plugins := sqlite.NewPluginService(db, apps, runtime, logger, certResolverName)
	if err != nil {
		logger.E("Failed to initialize docker runtime: %v", err)
		os.Exit(1)
	}
	resolver := &graphql.Resolver{
		Apps:       apps,
		Routes:     routes,
		EnvService: env,
		Plugins:    plugins,
		Runtime:    runtime,
		Version:    VERSION,
		Logger:     logger,
	}
	server := graphql.NewServer(logger, db, resolver, ACCESS_TOKEN)

	pollingUpdater := cron2.PollingUpgrader{
		Logger:  logger,
		Apps:    apps,
		Runtime: runtime,
		Routes:  routes,
		Plugins: plugins,
		Env:     env,
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
		cron.Recover(logger),
	))
	logger.I("Scheduled Jobs:")
	lo.ForEach(jobs, func(job scheduledJob, _ int) {
		c.AddFunc(job.expression, job.job)
		logger.I(" - %s (%s)", job.name, job.expression)
	})
	c.Start()
}
