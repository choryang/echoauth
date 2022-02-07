package service

import (
	// "encoding/base64"
	"fmt"

	// "log"
	// "io"
	// "net/http"

	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"

	// "github.com/cloud-barista/cb-webtool/src/model/tumblebug"
	tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"
	// tbmcir "github.com/cloud-barista/cb-webtool/src/model/tumblebug/mcir"
	// tbmcis "github.com/cloud-barista/cb-webtool/src/model/tumblebug/mcis"
	util "github.com/cloud-barista/cb-webtool/src/util"
)

func SetStoreNameSpaceList(c echo.Context, nameSpaceList []tbcommon.TbNsInfo) error {
	fmt.Println("====== SET NAME SPACE ========")
	fmt.Println(c)
	store := echosession.FromContext(c)
	store.Set(util.STORE_NAMESPACELIST, nameSpaceList)
	err := store.Save()
	return err
}
