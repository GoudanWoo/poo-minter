package poo_minter

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

type getUserUpgradesResponse struct {
	Faucet  Prop `json:"faucet"`  // 水龙头
	Storage Prop `json:"storage"` // 马桶
}

func (minter *PooMinter) getUserUpgrades(ctx g.Ctx) (userUpgrades getUserUpgradesResponse, err error) {
	stepName := "获取用户装备"

	request, err := minter.client.Get(ctx, "/user/upgrades")
	if err != nil {
		err = gerror.Wrapf(err, "%s 请求发送失败", stepName)
		return
	}
	defer request.Close()
	if request.StatusCode != http.StatusOK {
		err = gerror.Newf("%s 请求发送失败, HTTP %d, %s", stepName, request.StatusCode, request.Status)
		return
	}

	var response getUserUpgradesResponse
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&response)
	if err != nil {
		err = gerror.Wrapf(err, "%s 响应解析失败", stepName)
		return
	}

	userUpgrades = response

	return
}
