package resp

type ResultMessage struct {
	Code int `json:"code"`
	Msg string  `json:"msg"`
	Data interface{} `json:"data"`
}
