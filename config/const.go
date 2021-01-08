package config

type CodeMsg struct {
	RetCode string
	RetMsg  string
}

var (
	CodeMsgSuccess     = &CodeMsg{"RP0000", "成功"}
	CodeMsgBadRequest  = &CodeMsg{"RP4000", "请求参数错误"}
	CodeMsgServerError = &CodeMsg{"RP5000", "服务处理异常"}
)

const (
	DefaultVersionLimit         = 3
	DefaultVersionReservedLimit = 30
)
