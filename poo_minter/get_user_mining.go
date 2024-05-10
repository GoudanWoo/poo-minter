package poo_minter

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

type getUserMiningResponse struct {
	StorageSize     float64 `json:"storage_size"`      // 马桶容量
	StorageLevel    uint    `json:"storage_level"`     // 马桶等级
	MiningRate      float64 `json:"mining_rate"`       // 挖屎效率, 单位: 秒
	PointsInStorage float64 `json:"points_in_storage"` // 马桶暂存屎量
	SpendablePoints float64 `json:"spendable_points"`  // 可用屎量
	CurrentLeague   uint    `json:"current_league"`    // 皮搋子段位
}

func (minter *PooMinter) getUserMining(ctx g.Ctx) (userMining getUserMiningResponse, err error) {
	stepName := "获取用户数据"

	request, err := minter.client.Get(ctx, "/user/mining")
	if err != nil {
		err = gerror.Wrapf(err, "%s 请求发送失败", stepName)
		return
	}
	defer request.Close()
	if request.StatusCode != http.StatusOK {
		err = gerror.Newf("%s 请求发送失败, HTTP %d, %s", stepName, request.StatusCode, request.Status)
		return
	}

	var response getUserMiningResponse
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&response)
	if err != nil {
		err = gerror.Wrapf(err, "%s 响应解析失败", stepName)
		return
	}

	userMining = response

	return
}
