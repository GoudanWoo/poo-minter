package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/glog"
	"poo/poo_minter"
)

func init() {
	g.Log().SetLevel(glog.LEVEL_ALL ^ glog.LEVEL_DEBU)
}

func main() {
	ctx := context.Background()
	minter := poo_minter.NewPooMinter(genv.GetWithCmd("INIT_DATA").String())
	err := minter.Mint(ctx)
	if err != nil {
		g.Log().Fatalf(ctx, "%+v", err)
	}
}
