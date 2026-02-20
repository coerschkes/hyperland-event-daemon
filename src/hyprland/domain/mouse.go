package domain

type Mouse struct {
	Name        string `json:"name"`
	Sensitivity string
}

var RazerViperConfigurations = []Mouse{
	{
		Name:        "razer-razer-viper-v3-pro",
		Sensitivity: "-1",
	},
	{
		Name:        "razer-razer-viper-v3-pro-1",
		Sensitivity: "-1",
	},
	{
		Name:        "razer-razer-viper-v3-pro-mouse",
		Sensitivity: "-1",
	},
}
