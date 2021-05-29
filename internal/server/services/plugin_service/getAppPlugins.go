package plugin_service

import (
	"github.com/aklinker1/miasma/internal/server/utils/server_models"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func GetAppPlugins(tx *gorm.DB, appID string) (*server_models.AppPlugins, error) {
	log.V("plugin_server.GetAppPlugins(%v)", appID)

	traefik := &models.TraefikPluginConfig{}
	err := tx.First(traefik, "app_id = ?", appID).Error
	if err != nil {
		traefik = nil
	}

	return &server_models.AppPlugins{
		Traefik: traefik,
	}, nil
}
