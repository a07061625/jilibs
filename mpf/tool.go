/**
 * 工具方法
 * User: 姜伟
 * Date: 2019/12/24 0024
 * Time: 9:20
 */
package mpf

import (
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"

    "github.com/a07061625/gompf/mpf/mplog"
    "github.com/vmihailenco/msgpack/v4"
)

var (
    toolCharTotal    []byte
    toolCharLower    []byte
    toolCharNumLower []byte
    toolSeededRand   *rand.Rand
)

func init() {
    toolCharTotal = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'm', 'n', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
    toolCharLower = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'm', 'n', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
    toolCharNumLower = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'm', 'n', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
    toolSeededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 生成随机字符串
//   length int 生成字符串的长度
//   dataType string 数据类型
//     lower: 小写字母
//     numlower: 数字和小写字母
//     total: 数字和大小写字母
func ToolCreateNonceStr(length int, dataType string) string {
    b := make([]byte, length)

    switch dataType {
    case "lower":
        for i := range b {
            b[i] = toolCharLower[toolSeededRand.Intn(24)]
        }
    case "numlower":
        for i := range b {
            b[i] = toolCharNumLower[toolSeededRand.Intn(32)]
        }
    default:
        for i := range b {
            b[i] = toolCharTotal[toolSeededRand.Intn(57)]
        }
    }

    return string(b)
}

// 压缩数据
func ToolPack(data interface{}) ([]byte, error) {
    res, err := msgpack.Marshal(data)
    if err != nil {
        mplog.LogError("pack data error: " + err.Error())
        return nil, err
    }
    return res, nil
}

// 解压数据
func ToolUnpack(data []byte, item *interface{}) error {
    err := msgpack.Unmarshal(data, item)
    if err != nil {
        mplog.LogError("unpack data error: " + err.Error())
        return err
    }
    return nil
}

// 生成随机整数
//   startNum int 起始值
//   maxNum int 最大值
func ToolCreateRandNum(startNum, maxNum int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(maxNum) + startNum
}

// 生成请求ID
func ToolCreateReqId(reqId string) string {
    trueId := reqId
    if len(trueId) != 32 {
        nowTime := time.Now().Unix()
        needStr := ToolCreateNonceStr(8, "total") + strconv.FormatInt(nowTime, 10)
        trueId = HashMd5(needStr, "")
    }
    os.Setenv("MP_REQ_ID", trueId)

    return trueId
}

// 获取请求ID
func ToolGetReqId() string {
    reqId := os.Getenv("MP_REQ_ID")
    if len(reqId) == 32 {
        return reqId
    } else {
        return ToolCreateReqId("")
    }
}

// 处理错误
func ToolHandleError() func() {
    return func() {
        if p := recover(); p != nil {
            err, ok := p.(error)
            if ok {
                mplog.LogError(err.Error())
                return
            }
            fmt.Printf("%#v\n", p)
        }
    }
}
