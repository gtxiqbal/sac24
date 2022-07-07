package tl1

type DtoCmdDetail struct {
	Target    string   `json:"target"`
	Desc      string   `json:"desc"`
	Cmd       string   `json:"cmd"`
	Cmds      []string `json:"cmds"`
	SkipError bool     `json:"skip_error"`
	TimeSleep int      `json:"time_sleep"`
}
