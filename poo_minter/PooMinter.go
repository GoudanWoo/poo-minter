package poo_minter

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/genv"
	"golang.org/x/net/http2"
	"net/http"
)

type PooMinter struct {
	initData string
	client   *gclient.Client
}

func NewPooMinter(InitData string) *PooMinter {
	client := g.Client().
		Header(g.MapStrStr{
			"User-Agent": genv.GetWithCmd("USER_AGENT", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Mobile Safari/537.36").String(),
			"Referer":    "https://telemine-app.vercel.app/",
		}).
		SetPrefix("https://telemine-app.vercel.app/api").
		Proxy(genv.GetWithCmd("PROXY", "").String()).
		SetBrowserMode(true)
	err := http2.ConfigureTransport(client.Transport.(*http.Transport))
	if err != nil {
		panic(err)
	}
	return &PooMinter{
		initData: InitData,
		client:   client,
	}
}
