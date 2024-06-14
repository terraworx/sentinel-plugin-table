package main

import (
	"github.com/hashicorp/sentinel-sdk/rpc"
	table "github.com/terrawork/sentinel-plugin-table/plugin"
)

func main() {
	rpc.Serve(&rpc.ServeOpts{
		PluginFunc: table.New,
	})
}
