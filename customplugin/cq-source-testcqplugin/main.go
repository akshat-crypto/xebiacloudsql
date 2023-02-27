package main

import (
	"github.com/akshat-crypto/cq-source-testcqplugin/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
