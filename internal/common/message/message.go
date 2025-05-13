package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserName string `json:"userName"`
	UserPwd  string `json:"userPwd"`
}

type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type RegisterMes struct {
	UserName string `json:"userName"`
	UserPwd  string `json:"userPwd"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
