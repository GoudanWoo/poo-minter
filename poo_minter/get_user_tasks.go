package poo_minter

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

type getUserTasksResponse struct {
	Uid    string `json:"uid"`
	Status struct {
		ActionInviteReferee uint `json:"ACTION_INVITE_REFEREE"`
		ActionJoinChannel   bool `json:"ACTION_JOIN_CHANNEL"`
		ActionReactMessage  bool `json:"ACTION_REACT_MESSAGE"`
		ActionSetAddress    bool `json:"ACTION_SET_ADDRESS"`
		ActionJoinTwitter   bool `json:"ACTION_JOIN_TWITTER"`
	} `json:"status"`
	Price struct {
		ActionInviteReferee float64 `json:"ACTION_INVITE_REFEREE"`
		ActionJoinChannel   float64 `json:"ACTION_JOIN_CHANNEL"`
		ActionReactMessage  float64 `json:"ACTION_REACT_MESSAGE"`
		ActionSetAddress    float64 `json:"ACTION_SET_ADDRESS"`
		ActionJoinTwitter   float64 `json:"ACTION_JOIN_TWITTER"`
	} `json:"price"`
	Address   *string `json:"addr"`
	BetaMedal any     `json:"beta_medal"`
	Rewards   struct {
		ActionInviteReferee float64 `json:"ACTION_INVITE_REFEREE"`
		ActionJoinChannel   bool    `json:"ACTION_JOIN_CHANNEL"`
		ActionReactMessage  bool    `json:"ACTION_REACT_MESSAGE"`
		ActionSetAddress    bool    `json:"ACTION_SET_ADDRESS"`
		ActionJoinTwitter   bool    `json:"ACTION_JOIN_TWITTER"`
	} `json:"rewards"`
}

func (minter *PooMinter) getUserTasks(ctx g.Ctx) (userTasks getUserTasksResponse, err error) {
	stepName := "获取用户任务"

	request, err := minter.client.Get(ctx, "/user/tasks")
	if err != nil {
		err = gerror.Wrapf(err, "%s 请求发送失败", stepName)
		return
	}
	defer request.Close()
	if request.StatusCode != http.StatusOK {
		err = gerror.Newf("%s 请求发送失败, HTTP %d, %s", stepName, request.StatusCode, request.Status)
		return
	}

	var response getUserTasksResponse
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&response)
	if err != nil {
		err = gerror.Wrapf(err, "%s 响应解析失败", stepName)
		return
	}

	userTasks = response

	return
}
