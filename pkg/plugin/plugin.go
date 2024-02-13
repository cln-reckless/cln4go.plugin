package plugin

import (
	"os"

	"github.com/vincenzopalazzo/cln4go/plugin"
)

type PluginState struct{}

func Hello(plugin *plugin.Plugin[*PluginState], request map[string]any) (map[string]any, error) {
	return map[string]any{"message": "hello from cln4go.template"}, nil
}

func OnShutdown(plugin *plugin.Plugin[*PluginState], request map[string]any) {
	os.Exit(0)
}
