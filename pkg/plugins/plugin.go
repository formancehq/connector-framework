package plugins

import "github.com/formancehq/connector-framework/pkg/models"

type Plugin struct {
	models.PSPPlugin
	models.BankingBridgePlugin
}
