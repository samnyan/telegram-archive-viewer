package api

import "strconv"

type ApiResult struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *interface{} `json:"data"`
}

func SuccessResult(data *interface{}) *ApiResult {
	return &ApiResult{
		Code:    0,
		Message: "OK",
		Data:    data,
	}
}

func (r *ApiResult) Error() string {
	return strconv.Itoa(r.Code) + ": " + r.Message
}

func (a *ApiResult) SetAppError(e *AppError) {
	a.Code = e.Code
	a.Message = e.Message
}
