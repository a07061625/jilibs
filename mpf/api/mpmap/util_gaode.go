/**
 * Created by GoLand.
 * User: 姜伟
 * Date: 2020/1/17 0017
 * Time: 16:32
 */
package mpmap

import (
    "github.com/a07061625/gompf/mpf"
    "github.com/a07061625/gompf/mpf/api"
)

func (util *utilMap) SendGaoDeRequest(service IMapBase, errorCode uint) api.ApiResult {
    resp, result := util.SendOuter(service, errorCode)
    if result.Code > 0 {
        return result
    }

    respData, _ := mpf.JSONUnmarshalMap(resp.Content)
    nowStatus, ok := respData["status"]
    if ok && (nowStatus.(int) == 1) {
        result.Data = respData
        return result
    }

    result.Code = errorCode
    msg, ok := respData["info"]
    if ok {
        result.Msg = msg.(string)
    } else {
        result.Msg = "解析服务数据出错"
    }
    return result
}
