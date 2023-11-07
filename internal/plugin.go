package internal

import (
	"fmt"

	"github.com/mach-composer/mach-composer-plugin-helpers/helpers"
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"
	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"
)

func NewStoryblokPlugin() schema.MachComposerPlugin {
	state := &Plugin{
		provider:    "0.5.4",
		siteConfigs: map[string]*StoryblokConfig{},
	}

	return plugin.NewPlugin(&schema.PluginSchema{
		Identifier: "storyblok",

		Configure: state.Configure,
		IsEnabled: func() bool { return state.enabled },

		GetValidationSchema: state.GetValidationSchema,

		// Config
		SetGlobalConfig: state.SetGlobalConfig,
		SetSiteConfig:   state.SetSiteConfig,

		// Renders
		RenderTerraformProviders: state.TerraformRenderProviders,
		RenderTerraformResources: state.TerraformRenderResources,
		RenderTerraformComponent: state.RenderTerraformComponent,
	})
}

type Plugin struct {
	environment  string
	provider     string
	globalConfig *StoryblokConfig
	siteConfigs  map[string]*StoryblokConfig
	enabled      bool
}

func (p *Plugin) Configure(environment string, provider string) error {
	p.environment = environment
	if provider != "" {
		p.provider = provider
	}
	return nil
}

func (p *Plugin) GetValidationSchema() (*schema.ValidationSchema, error) {
	result := getSchema()
	return result, nil
}

func (p *Plugin) SetGlobalConfig(data map[string]any) error {
	cfg := StoryblokConfig{}
	if err := mapstructure.Decode(data, &cfg); err != nil {
		return err
	}
	p.globalConfig = &cfg
	p.enabled = true
	return nil
}

func (p *Plugin) SetSiteConfig(site string, data map[string]any) error {
	cfg := StoryblokConfig{}
	if err := mapstructure.Decode(data, &cfg); err != nil {
		return err
	}
	p.siteConfigs[site] = &cfg
	p.enabled = true
	return nil
}

func (p *Plugin) getSiteConfig(site string) *StoryblokConfig {
	result := &StoryblokConfig{}
	if p.globalConfig != nil {
		result.URL = p.globalConfig.URL
		result.Token = p.globalConfig.Token
		result.SpaceID = p.globalConfig.SpaceID
	}

	cfg, ok := p.siteConfigs[site]
	if ok {
		if cfg.URL != "" {
			result.URL = cfg.URL
		}
		if cfg.Token != "" {
			result.Token = cfg.Token
		}
		if cfg.SpaceID != "" {
			result.SpaceID = cfg.SpaceID
		}
	}

	if result.URL == "" {
		return nil
	}
	return result
}

func (p *Plugin) TerraformRenderProviders(site string) (string, error) {
	cfg := p.getSiteConfig(site)
	if cfg == nil {
		return "", nil
	}

	result := fmt.Sprintf(`
	storyblok = {
		source = "labd/storyblok"
		version = "%s"
	}`, helpers.VersionConstraint(p.provider))
	return result, nil
}

func (p *Plugin) TerraformRenderResources(site string) (string, error) {
	cfg := p.getSiteConfig(site)
	if cfg == nil {
		return "", nil
	}

	template := `
		provider "storyblok" {
			{{ renderProperty "url" .URL }}
			{{ renderProperty "token" .Token }}
		}
	`
	return helpers.RenderGoTemplate(template, cfg)
}

func (p *Plugin) RenderTerraformComponent(site string, component string) (*schema.ComponentSchema, error) {
	cfg := p.getSiteConfig(site)
	if cfg == nil {
		return nil, nil
	}

	template := `
		{{ renderProperty "storyblok_space_id" .SpaceID }}
	`
	vars, err := helpers.RenderGoTemplate(template, cfg)
	if err != nil {
		return nil, err
	}
	result := &schema.ComponentSchema{
		Variables: vars,
	}
	return result, nil
}
