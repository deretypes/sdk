package manga

type BaseProvider struct {
	Settings *Settings
}

func NewBaseProvider(settings *Settings) *BaseProvider {
	return &BaseProvider{
		Settings: settings,
	}
}

func (p *BaseProvider) GetSettings() (*Settings, error) {
	return p.Settings, nil
}

func (p *BaseProvider) SetSettings(opts ...SettingOpt) {
	for _, opt := range opts {
		opt(p.Settings)
	}
}

func SetSessionToken(sessionToken string) SettingOpt {
	return func(settings *Settings) {
		settings.SessionToken = &sessionToken
	}
}
