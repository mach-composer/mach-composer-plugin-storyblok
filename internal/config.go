package internal

type StoryblokConfig struct {
	URL     string `mapstructure:"url"`
	Token   string `mapstructure:"token"`
	SpaceID int    `mapstructure:"space_id"`
}
