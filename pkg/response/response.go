package response

import (
	"gas-detect/pkg/response/status"
	kratosError "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"time"
)

// R 通用分页返回对象
type R struct {
	Code      status.Code `json:"code"`
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	RequestId string      `json:"requestId"`
	Timestamp string      `json:"timestamp"`
	Exception string      `json:"exception"`
}

func Result(code status.Code, success bool, data interface{}, message string, exception string) R {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return R{
		code,
		success,
		data,
		message,
		uuid.NewString(),
		currentTime,
		exception,
	}
}

func Success() R {
	return Result(status.Success, true, map[string]interface{}{}, status.Message(status.Success), "")
}

func SuccessWithMessage(message string) R {
	return Result(status.Success, true, map[string]interface{}{}, message, "")
}

func SuccessWithData(data interface{}) R {
	return Result(status.Success, true, data, status.Message(status.Success), "")
}

func SuccessWithDetailed(data interface{}, message string) R {
	return Result(status.Success, true, data, message, "")
}

func Fail(code status.Code, exception string) R {
	return Result(code, false, map[string]interface{}{}, status.Message(code), exception)
}

func FailWithMessage(code status.Code, message string, exception string) R {
	return Result(code, false, map[string]interface{}{}, message, exception)
}

func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	var resp R
	if v != nil {
		resp = SuccessWithData(v)
	} else {
		resp = Success()
	}

	if err := ResponseWrite(w, r, resp, int(status.Success)); err != nil {
		return err
	}
	return nil
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {

	var (
		code    int
		message string
		reason  string
		resp    R
	)

	if se := kratosError.FromError(err); se != nil {
		code = int(se.Code)
		message = se.Message
		reason = se.Reason
	} else {
		code = int(status.InternalServerError)
	}
	resp = FailWithMessage(status.Code(code), message, reason)

	ResponseWrite(w, r, resp, code)
}

func ResponseWrite(w http.ResponseWriter, r *http.Request, rsp R, code int) (err error) {
	// 通过Request Header的Accept中提取出对应的编码器
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(rsp)
	if err != nil {
		w.WriteHeader(int(status.InternalServerError))
		return err
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	// 设置HTTP Status Code
	w.WriteHeader(code)
	w.Write(body)
	return nil
}
