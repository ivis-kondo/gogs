package gin

import (
	"fmt"
	"net/http"

	"github.com/NII-DG/gogs/internal/conf"
	"github.com/NII-DG/gogs/internal/context"
)

type ServerInfo struct {
	Http string `json:"http"`
	Ssh  string `json:"ssh"`
}

func GetServerInfo(c *context.APIContext) {
	// "ginSsh": "ssh://git@dg02.dg.rcos.nii.ac.jp:3001"
	ssh_url := fmt.Sprintf("ssh://%s@%s:%d", conf.App.RunUser, conf.SSH.Domain, conf.SSH.Port)
	// "ginHttp": "http://dg02.dg.rcos.nii.ac.jp",
	http_url := fmt.Sprintf("%s", conf.Server.ExternalURL)
	c.JSON(http.StatusOK, ServerInfo{
		Http: http_url,
		Ssh:  ssh_url,
	})
}
