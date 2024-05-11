package updater

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"golang.org/x/net/http2"
	"net/http"
	"strconv"
)

var VersionStr string
var versionCode_ string
var VersionCode uint64

const RepoOwner = "GoudanWoo"
const Repo = "poo-minter"

func init() {
	VersionCode, _ = strconv.ParseUint(versionCode_, 10, 64)
}

func Check(ctx g.Ctx) (updatable bool, versionStr string, versionCode uint64, err error) {
	client := g.Client().
		Proxy(genv.GetWithCmd("PROXY", "").String()).
		SetBrowserMode(true)
	err = http2.ConfigureTransport(client.Transport.(*http.Transport))
	if err != nil {
		panic(err)
	}

	stepName := "检查更新"

	request, err := client.Get(ctx, fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", RepoOwner, Repo), g.MapStrAny{
		"per_page": 1,
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

	links := gstr.Split(request.Header.Get("Link"), ", ")
	if len(links) == 2 && gstr.HasSuffix(links[1], "; rel=\"last\"") {
		var matchString []string
		matchString, err = gregex.MatchString(`&page=(\d+)`, links[1])
		if err != nil {
			err = gerror.Wrapf(err, "未匹配到最新版本号")
			return
		}
		versionCode, err = strconv.ParseUint(matchString[1], 10, 64)
		if err != nil {
			err = gerror.Wrapf(err, "未解析出最新版本号")
			return
		}
	} else {
		err = gerror.Newf("未找到最新版本号")
		return
	}

	response, err := gjson.DecodeToJson(request.ReadAll())
	if err != nil {
		err = gerror.Wrapf(err, "未解析出最新版本")
		return
	}

	versionStr = response.Get("0.sha").String()[:7]

	updatable = versionCode > VersionCode

	return
}
