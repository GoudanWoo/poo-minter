package poo_minter

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/grand"
	"time"
)

func (minter *PooMinter) Mint(ctx g.Ctx) (err error) {
	ctx = context.WithValue(ctx, gctx.StrKey(`MiddlewareClientTracingHandled`), 1)
	err = minter.login(ctx)
	if err != nil {
		err = gerror.Wrapf(err, "%s 失败", "上号")
		return
	}

	user, err := minter.getUserMe(ctx)
	if err != nil {
		err = gerror.Wrapf(err, "%s 失败", "获取用户信息")
		return
	}

	g.Log().Infof(ctx, "你好, %s(%s)", user.Username, user.Id)

	for {
		userMining, err := minter.getUserMining(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "%s 失败: %+v", "获取用户数据", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		switch userMining.CurrentLeague {
		case -1:
			g.Log().Infof(ctx, "当前段位: %s", "没皮搋子")
		case 0:
			g.Log().Infof(ctx, "当前段位: %s", "木皮搋子")
		case 1:
			g.Log().Infof(ctx, "当前段位: %s", "铜皮搋子")
		case 2:
			g.Log().Infof(ctx, "当前段位: %s", "银皮搋子")
		case 3:
			g.Log().Infof(ctx, "当前段位: %s", "金皮搋子")
		case 4:
			g.Log().Infof(ctx, "当前段位: %s", "白金皮搋子")
		}

		userUpgrades, err := minter.getUserUpgrades(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "%s 失败: %+v", "获取用户装备", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		g.Log().Infof(ctx, "水龙头 等级: Lv.%d, 效率: %.2f 屎/h -> %.2f 屎/h, 需花费 %.0f 屎", userUpgrades.Faucet.NextLevel-1, userUpgrades.Faucet.CurrentValue*3600, userUpgrades.Faucet.NextValue*3600, userUpgrades.Faucet.Cost)
		g.Log().Infof(ctx, "马桶 等级: Lv.%d, 容量: %.0f 屎 -> %.0f 屎, 需花费 %.0f 屎", userUpgrades.Storage.NextLevel-1, userUpgrades.Storage.CurrentValue, userUpgrades.Storage.NextValue, userUpgrades.Storage.Cost)
		faucetUpgradable := userUpgrades.Faucet.NextLevel < userUpgrades.Faucet.Cap
		if !faucetUpgradable {
			g.Log().Infof(ctx, "%s 已到最大等级限制", "水龙头")
		}
		storageUpgradable := userUpgrades.Storage.NextLevel < userUpgrades.Storage.Cap
		if !storageUpgradable {
			g.Log().Infof(ctx, "%s 已到最大等级限制", "马桶")
		}

		balance := userMining.PointsInStorage + userMining.SpendablePoints
		percentage := userMining.PointsInStorage / userMining.StorageSize
		g.Log().Infof(ctx, "总共已有 %.2f 屎, 其中 马桶 已存 %.2f / %.0f 屎, 达到容量的 %.2f%%", balance, userMining.PointsInStorage, userMining.StorageSize, percentage*100)

		if faucetUpgradable && userMining.SpendablePoints >= userUpgrades.Faucet.Cost {
			g.Log().Infof(ctx, "准备升级 %s", "水龙头")
			err := minter.upgrade(ctx, "faucet")
			if err != nil {
				g.Log().Errorf(ctx, "%s 失败: %+v", "升级道具", err)
				time.Sleep(1 * time.Minute)
				continue
			}
			g.Log().Infof(ctx, "升级 %s 成功", "水龙头")
			continue
		}

		if faucetUpgradable && balance >= userUpgrades.Faucet.Cost || percentage >= 0.9 {
			g.Log().Infof(ctx, "准备掏屎")
			err := minter.claim(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "%s 失败: %+v", "掏屎", err)
				time.Sleep(1 * time.Minute)
				continue
			}
			g.Log().Infof(ctx, "掏屎成功")
			continue
		} else {
			var waitShit float64 = 0
			waitShit1 := userUpgrades.Faucet.Cost - userMining.SpendablePoints
			waitShit2 := userMining.StorageSize * 0.9
			if faucetUpgradable && waitShit1 < waitShit2 {
				g.Log().Infof(ctx, "下次掏屎后可升级 %s", "水龙头")
				waitShit = waitShit1
			} else {
				waitShit = waitShit2
			}
			waitShit -= userMining.PointsInStorage
			waitTime := time.Duration(waitShit/userMining.MiningRate) * time.Second
			g.Log().Infof(ctx, "大概等待 %.2f 屎 (≈ %s) 后掏屎", waitShit, waitTime)
			time.Sleep(min(waitTime, time.Duration(grand.N(30, 300))*time.Second))
			continue
		}
	}
}
