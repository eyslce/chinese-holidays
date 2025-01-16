package chinese_holidays

type Config struct {
	SavePath string
}

type Option func(cfg *Config)

func WithSavePath(savePath string) Option {
	return func(cfg *Config) {
		cfg.SavePath = savePath
	}
}

func applyOptions(cfg *Config, opts ...Option) {
	for _, opt := range opts {
		opt(cfg)
	}
}
