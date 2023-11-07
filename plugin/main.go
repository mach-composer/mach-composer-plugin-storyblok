package plugin

import (
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"

	"github.com/mach-composer/mach-composer-plugin-storyblok/internal"
)

func Serve() {
	p := internal.NewStoryblokPlugin()
	plugin.ServePlugin(p)
}
