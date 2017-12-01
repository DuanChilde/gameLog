package controllers

var res JsonRes

type JsonRes struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func (j *JsonRes) Success(msg string, data map[string]interface{}) JsonRes {
	res.Code = 1
	res.Data = data
	res.Msg = msg
	return res
}

func (j *JsonRes) Fail(errMsg string) JsonRes {
	res.Code = 0
	res.Msg = errMsg
	res.Data = make(map[string]interface{})
	return res
}
