package state

const INVALID_PID = -1

type PwdRegistry struct {
	currentPwd string
	currentApp string
}

func NewPwdRegistry() *PwdRegistry {
	return &PwdRegistry{
		currentPwd: "~/",
		currentApp: "",
	}
}

func (r *PwdRegistry) GetCurrentPwd() string {
	if r.currentApp != "Alacritty" {
		return "~/"
	}
	return r.currentPwd
}

func (r *PwdRegistry) UpdateCurrentPwd(pwd string) {
	r.currentPwd = pwd
}

func (r *PwdRegistry) GetCurrentApp() string {
	return r.currentApp
}

func (r *PwdRegistry) UpdateCurrentApp(app string) {
	r.currentApp = app
}
