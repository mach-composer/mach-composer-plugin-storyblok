package internal

type StoryblokConfig struct {
	URL     string `mapstructure:"url"`
	Token   string `mapstructure:"token"`
	SpaceID string `mapstructure:"space_id"`
}
