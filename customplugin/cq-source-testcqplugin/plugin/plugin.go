package plugin

import (
	// "github.com/akshat-crypto/cq-source-testcqplugin/client"
	// "github.com/akshat-crypto/cq-source-testcqplugin/resources"
	// "github.com/cloudquery/plugin-sdk/plugins/source"
	// "github.com/cloudquery/plugin-sdk/schema"
	"github.com/akshat-crypto/cq-source-testcqplugin/client"
	"github.com/akshat-crypto/cq-source-testcqplugin/resources"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"testcqplugin",
		Version,
		schema.Tables{
			resources.Comics(),
		},
		client.New,
	)
}
