package internal

import (
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"
)

func Serve() {
	p := NewStoryblokPlugin()
	plugin.ServePlugin(p)
}
