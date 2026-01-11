package state

const INVALID_PID = -1

type PwdRegistry struct {
	currentPid int
	pwdMap     map[int]string
}

func NewPwdRegistry() *PwdRegistry {
	return &PwdRegistry{
		currentPid: INVALID_PID,
		pwdMap:     make(map[int]string),
	}
}

func (r *PwdRegistry) UpsertEntry(pid int, pwd string) {
	r.pwdMap[pid] = pwd
}

func (r *PwdRegistry) DeleteEntry(pid int) {
	delete(r.pwdMap, pid)
}

func (r *PwdRegistry) GetPwd(pid int) string {
	return r.pwdMap[pid]
}

func (r *PwdRegistry) UpdateCurrentPid(pid int) {
	r.currentPid = pid
}

func (r *PwdRegistry) GetCurrentPid() int {
	return r.currentPid
}

func (r *PwdRegistry) GetCurrentPidPwd() string {
	if r.currentPid == INVALID_PID {
		return ""
	}

	return r.GetPwd(r.currentPid)
}
