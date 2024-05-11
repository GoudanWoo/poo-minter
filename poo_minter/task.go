package poo_minter

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

type taskResponse struct {
	Message string `json:"message"`
}

func (minter *PooMinter) task(ctx g.Ctx, task string) (err error) {
	stepName := "提交任务"

	request, err := minter.client.ContentJson().Post(ctx, "/user/tasks", g.MapStrAny{
		"task_type": task,
	})
	if err != nil {
		err = gerror.Wrapf(err, "%s 请求发送失败", stepName)
		return
	}
	defer request.Close()
	if request.StatusCode != http.StatusOK {
		err = gerror.Newf("%s 请求发送失败, HTTP %d, %s", stepName, request.StatusCode, request.Status)
		return
	}

	var response taskResponse
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&response)
	if err != nil {
		err = gerror.Wrapf(err, "%s 响应解析失败", stepName)
		return
	}

	if response.Message != "Ok" {
		err = gerror.Wrapf(err, "%s 失败, %s", stepName, response.Message)
		return
	}

	return
}
