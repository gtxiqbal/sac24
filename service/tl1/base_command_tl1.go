package tl1

type BaseCommandTL1 interface {
	SetValue()
	Login()
	Logout()
	Ont(status int, onuType map[string]string, name string)
}
