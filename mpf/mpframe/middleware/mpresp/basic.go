/**
 * Created by GoLand.
 * User: 姜伟
 * Date: 2020/2/8 0008
 * Time: 12:47
 */
package mpresp

import (
    "os"

    "github.com/a07061625/gompf/mpf"
    "github.com/a07061625/gompf/mpf/mpconstant/errorcode"
    "github.com/a07061625/gompf/mpf/mpconstant/project"
    "github.com/a07061625/gompf/mpf/mpresponse"
    "github.com/kataras/iris/v12/context"
)

// 发送响应数据
func NewBasicSend() context.Handler {
    return func(ctx context.Context) {
        respData, ok := ctx.Values().GetEntry(project.DataParamKeyRespData)
        if ok {
            data, ok := respData.ValueRaw.(string)
            if ok {
                ctx.Recorder().Header().Set(project.HttpHeadKeyContentType, project.HttpContentTypeText)
                ctx.Recorder().SetBodyString(data)
            } else {
                result := mpresponse.NewResultBasic()
                result.Data = data
                ctx.Recorder().Header().Set(project.HttpHeadKeyContentType, project.HttpContentTypeJson)
                ctx.Recorder().SetBodyString(mpf.JsonMarshal(result))
            }
        } else {
            result := mpresponse.NewResultBasic()
            result.Code = errorcode.CommonBaseServer
            result.Msg = "响应数据不能为空"
            ctx.Recorder().Header().Set(project.HttpHeadKeyContentType, project.HttpContentTypeJson)
            ctx.Recorder().SetBodyString(mpf.JsonMarshal(result))
        }
    }
}

// 请求响应结束
func NewBasicEnd() context.Handler {
    return func(ctx context.Context) {
        os.Unsetenv(project.DataParamKeyReqId)
        ctx.Values().Remove(project.DataParamKeyReqUrl)
        ctx.Values().Remove(project.DataParamKeyRespData)
        ctx.EndRequest()
    }
}
