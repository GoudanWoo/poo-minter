package poo_minter

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

type getUserMeResponse struct {
	Id       string `json:"id"`
	Icon     string `json:"icon"`
	Username string `json:"username"`
}

func (minter *PooMinter) getUserMe(ctx g.Ctx) (userMe getUserMeResponse, err error) {
	stepName := "获取用户信息"

	request, err := minter.client.Get(ctx, "/user/me")
	if err != nil {
		err = gerror.Wrapf(err, "%s 请求发送失败", stepName)
		return
	}
	defer request.Close()
	if request.StatusCode != http.StatusOK {
		err = gerror.Newf("%s 请求发送失败, HTTP %d, %s", stepName, request.StatusCode, request.Status)
		return
	}

	var response getUserMeResponse
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&response)
	if err != nil {
		err = gerror.Wrapf(err, "%s 响应解析失败", stepName)
		return
	}

	userMe = response

	return
}
