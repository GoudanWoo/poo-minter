package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/glog"
	"poo/poo_minter"
	"poo/updater"
)

func init() {
	g.Log().SetLevel(glog.LEVEL_ALL ^ glog.LEVEL_DEBU)
}

func main() {
	ctx := context.Background()
	if updatable, versionStr, versionCode, err := updater.Check(ctx); err != nil {
		g.Log().Errorf(ctx, "检查更新失败: %+v", err)
	} else if updatable {
		g.Log().Warningf(ctx, "脚本有新版本 %s(%d) -> %s(%d), 请前往 https://github.com/%s/%s 更新", updater.VersionStr, updater.VersionCode, versionStr, versionCode, updater.RepoOwner, updater.Repo)
	}

	minter := poo_minter.NewPooMinter(genv.GetWithCmd("INIT_DATA").String())
	err := minter.Mint(ctx)
	if err != nil {
		g.Log().Fatalf(ctx, "%+v", err)
	}
}
