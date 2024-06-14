package main

import (
	"github.com/hashicorp/sentinel-sdk/rpc"
	table "github.com/terraworx/sentinel-plugin-table/plugin"
)

func main() {
	rpc.Serve(&rpc.ServeOpts{
		PluginFunc: table.New,
	})
}
