// // This file is auto-generated, don't edit it. Thanks.
package main

//
//import (
//	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
//	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
//	util "github.com/alibabacloud-go/tea-utils/v2/service"
//	"github.com/alibabacloud-go/tea/tea"
//	"os"
//)
//
///**
//* 使用AK&SK初始化账号Client
//* @param accessKeyId
//* @param accessKeySecret
//* @return Client
//* @throws Exception
// */
//
//func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
//	config := &openapi.Config{
//		// 必填，您的 AccessKey ID
//		AccessKeyId: accessKeyId,
//		// 必填，您的 AccessKey Secret
//		AccessKeySecret: accessKeySecret,
//	}
//	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
//	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
//	_result = &dysmsapi20170525.Client{}
//	_result, _err = dysmsapi20170525.NewClient(config)
//	return _result, _err
//}
//
//func _main(args []*string) (_err error) {
//	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
//	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
//	client, _err := CreateClient(tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")), tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")))
//	if _err != nil {
//		return _err
//	}
//
//	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
//		SignName:      tea.String("马帅的博客"),
//		TemplateCode:  tea.String("SMS_462745353"),
//		PhoneNumbers:  tea.String("18739101871"),
//		TemplateParam: tea.String("{\"code\":\"6666\"}"),
//	}
//	runtime := &util.RuntimeOptions{}
//	tryErr := func() (_e error) {
//		defer func() {
//			if r := tea.Recover(recover()); r != nil {
//				_e = r
//			}
//		}()
//		// 复制代码运行请自行打印 API 的返回值
//		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
//		if _err != nil {
//			return _err
//		}
//
//		return nil
//	}()
//
//	if tryErr != nil {
//		var error = &tea.SDKError{}
//		if _t, ok := tryErr.(*tea.SDKError); ok {
//			error = _t
//		} else {
//			error.Message = tea.String(tryErr.Error())
//		}
//		// 如有需要，请打印 error
//		_, _err = util.AssertAsString(error.Message)
//		if _err != nil {
//			return _err
//		}
//	}
//	return _err
//}
//
//func main() {
//	err := _main(tea.StringSlice(os.Args[1:]))
//	if err != nil {
//		panic(err)
//	}
//}
