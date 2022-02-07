package controller

import (
	"echoauth/src/service"
	"fmt"
	"log"
	"net/http"

	tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"

	"github.com/labstack/echo"
)

// 사용자의 namespace 목록 조회
func GetNameSpaceList(c echo.Context) error {
	fmt.Println("====== GET NAMESPACE LIST ========")

	optionParam := c.QueryParam("option")

	if optionParam == "id" {
		nameSpaceInfoList, nsStatus := service.GetNameSpaceListByOptionID(optionParam)
		if nsStatus.StatusCode == 500 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": nsStatus.Message,
				"status":  nsStatus.StatusCode,
			})
		}

		return c.JSON(http.StatusOK, nameSpaceInfoList)
	} else {
		nameSpaceInfoList, nsStatus := service.GetNameSpaceListByOption(optionParam)
		if nsStatus.StatusCode == 500 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": nsStatus.Message,
				"status":  nsStatus.StatusCode,
			})
		}
		setNameSpaceErr := service.SetStoreNameSpaceList(c, nameSpaceInfoList)
		if setNameSpaceErr != nil {
			fmt.Println("error", setNameSpaceErr)
		}
		return c.JSON(http.StatusOK, nameSpaceInfoList)
	}

}

// namespace 등록 처리
func NameSpaceRegProc(c echo.Context) error {

	// loginInfo := service.CallLoginInfo(c)
	// if loginInfo.UserID == "" {

	// 	// Login 정보가 없으므로 login화면으로
	// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
	// }

	// namespace := c.FormValue("name")
	// description := c.FormValue("description")
	// fmt.Println("namespace : " + namespace + " , description :" + description)
	// nameSpaceInfo := new(model.NameSpaceInfo)
	// nameSpaceInfo.Name = namespace
	// nameSpaceInfo.Description = description

	nameSpaceInfo := new(tbcommon.TbNsInfo)
	if err := c.Bind(nameSpaceInfo); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// if err = c.Validate(nameSpaceInfo); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }
	fmt.Println("nameSpaceInfo : ", nameSpaceInfo)

	// Tubblebug 호출하여 namespace 생성

	// person := Person{"Alex", 10}
	// pbytes, _ := json.Marshal(person)
	respBody, respStatus := service.RegNameSpace(nameSpaceInfo)
	fmt.Println("=============respBody =============", respBody)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	// 저장 성공하면 namespace 목록 조회
	nameSpaceList, nsStatus := service.GetNameSpaceList()
	if nsStatus.StatusCode == 500 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":       respStatus.Message,
			"status":        respStatus.StatusCode,
			"NameSpaceList": nil,
		})
	}
	storeNameSpaceErr := service.SetStoreNameSpaceList(c, nameSpaceList)
	if storeNameSpaceErr != nil {
		log.Println("Store NameSpace Err")
	}
	// return namespace 목록
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        "200",
		"NameSpaceList": nameSpaceList,
	})
}

// NameSpace 삭제
func NameSpaceDelProc(c echo.Context) error {
	log.Println("NameSpaceDelProc : ")

	paramNameSpaceID := c.Param("nameSpaceID")

	respBody, respStatus := service.DelNameSpace(paramNameSpaceID)
	fmt.Println("=============respBody =============", respBody)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	// 저장 성공하면 namespace 목록 조회
	nameSpaceList, nsStatus := service.GetNameSpaceList()
	if nsStatus.StatusCode == 500 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":       nsStatus.Message,
			"status":        nsStatus.StatusCode,
			"NameSpaceList": nil,
		})
	}

	storeNameSpaceErr := service.SetStoreNameSpaceList(c, nameSpaceList)
	if storeNameSpaceErr != nil {
		log.Println("Store NameSpace Err")
	}

	// return namespace 목록
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        respStatus,
		"NameSpaceList": nameSpaceList,
	})

	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"message": "success",
	// 	"status":  "200",
	// })
}
