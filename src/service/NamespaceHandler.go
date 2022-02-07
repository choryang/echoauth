package service

import (
	"echoauth/src/model"
	"echoauth/src/util"
	"encoding/json"
	"fmt"
	"net/http"

	tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"
)

func GetNameSpaceListByOptionID(optionParam string) ([]string, model.WebStatus) {
	fmt.Println("GetNameSpaceList start")
	var originalUrl = "/ns"
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.TUMBLEBUG + urlParam
	if optionParam == "id" {
		url = url + "?option=" + optionParam
	} else {
		return nil, model.WebStatus{StatusCode: 500, Message: "option param is not ID"}
	}
	// url := util.TUMBLEBUG + "/ns"

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	//body := HttpGetHandler(url)

	if err != nil {
		// 	// Tumblebug 접속 확인하라고
		// fmt.Println(err)
		// panic(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	//nameSpaceInfoList := map[string][]string{}
	nameSpaceInfoList := tbcommon.TbIdList{}
	// defer body.Close()
	json.NewDecoder(respBody).Decode(&nameSpaceInfoList)
	//spew.Dump(body)
	//fmt.Println(nameSpaceInfoList["idList"])
	//
	//return nameSpaceInfoList["idList"], model.WebStatus{StatusCode: respStatus}
	//fmt.Println(nameSpaceInfoList["output"])
	//return nameSpaceInfoList["output"], model.WebStatus{StatusCode: respStatus}
	fmt.Println(nameSpaceInfoList.IDList)
	return nameSpaceInfoList.IDList, model.WebStatus{StatusCode: respStatus}
}

func GetNameSpaceListByOption(optionParam string) ([]tbcommon.TbNsInfo, model.WebStatus) {
	fmt.Println("GetNameSpaceList start")
	var originalUrl = "/ns"
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.TUMBLEBUG + urlParam
	if optionParam != "" {
		url = url + "?option=" + optionParam
	}
	// url := util.TUMBLEBUG + "/ns"

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	//body := HttpGetHandler(url)

	if err != nil {
		// 	// Tumblebug 접속 확인하라고
		// fmt.Println(err)
		// panic(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	nameSpaceInfoList := map[string][]tbcommon.TbNsInfo{}
	// defer body.Close()
	json.NewDecoder(respBody).Decode(&nameSpaceInfoList)
	//spew.Dump(body)
	fmt.Println(nameSpaceInfoList["ns"])

	return nameSpaceInfoList["ns"], model.WebStatus{StatusCode: respStatus}
}

// NameSpace 등록.  등록 후 생성된 Namespace 정보를 return
func RegNameSpace(nameSpaceInfo *tbcommon.TbNsInfo) (tbcommon.TbNsInfo, model.WebStatus) {
	// buff := bytes.NewBuffer(pbytes)
	var originalUrl = "/ns"
	urlParam := util.MappingUrlParameter(originalUrl, nil)
	url := util.TUMBLEBUG + urlParam
	// url := util.TUMBLEBUG + "/ns"

	//body, err := util.CommonHttpPost(url, nameSpaceInfo)
	pbytes, _ := json.Marshal(nameSpaceInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode

	resultNameSpaceInfo := tbcommon.TbNsInfo{}
	if err != nil {
		fmt.Println(err)
		failResultInfo := tbcommon.TbSimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return resultNameSpaceInfo, model.WebStatus{StatusCode: 500, Message: failResultInfo.Message}
	}

	json.NewDecoder(respBody).Decode(&resultNameSpaceInfo)
	return resultNameSpaceInfo, model.WebStatus{StatusCode: respStatus}
	//return respBody, model.WebStatus{StatusCode: respStatus}
}

// 사용자의 namespace 목록 조회
func GetNameSpaceList() ([]tbcommon.TbNsInfo, model.WebStatus) {
	fmt.Println("GetNameSpaceList start")
	var originalUrl = "/ns"
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.TUMBLEBUG + urlParam
	// url := util.TUMBLEBUG + "/ns"

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	//body := HttpGetHandler(url)

	if err != nil {
		// 	// Tumblebug 접속 확인하라고
		// fmt.Println(err)
		// panic(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	nameSpaceInfoList := map[string][]tbcommon.TbNsInfo{}
	// defer body.Close()
	json.NewDecoder(respBody).Decode(&nameSpaceInfoList)
	//spew.Dump(body)
	fmt.Println(nameSpaceInfoList["ns"])

	return nameSpaceInfoList["ns"], model.WebStatus{StatusCode: respStatus}
}

// NameSpace 삭제
func DelNameSpace(nameSpaceID string) (tbcommon.TbSimpleMsg, model.WebStatus) {
	var originalUrl = "/ns/{nsId}"
	var paramMapper = make(map[string]string)
	paramMapper["{nsId}"] = nameSpaceID
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)
	url := util.TUMBLEBUG + urlParam
	// url := util.TUMBLEBUG + "/ns/" + nameSpaceID

	// 경로안에 parameter가 있어 추가 param없이 호출 함.
	resp, err := util.CommonHttp(url, nil, http.MethodDelete)

	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode

	resultInfo := tbcommon.TbSimpleMsg{}
	json.NewDecoder(respBody).Decode(&resultInfo)
	if err != nil {
		fmt.Println(err)
		//return resultInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
		json.NewDecoder(respBody).Decode(&resultInfo)
		return resultInfo, model.WebStatus{StatusCode: 500, Message: resultInfo.Message}
	}

	return resultInfo, model.WebStatus{StatusCode: respStatus}
}
