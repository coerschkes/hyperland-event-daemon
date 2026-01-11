package state

import "fmt"

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
	fmt.Println("UPDATING PWD")
	r.currentPwd = pwd
	fmt.Println("PWD: " + r.currentPwd)
}

func (r *PwdRegistry) GetCurrentApp() string {
	return r.currentApp
}

func (r *PwdRegistry) UpdateCurrentApp(app string) {
	r.currentApp = app
}
