// Package security_and_compliance code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larksecurity_and_compliance

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *SecurityAndComplianceService {
	s := &SecurityAndComplianceService{config: config}
	s.OpenapiLog = &openapiLog{service: s}
	return s
}

type SecurityAndComplianceService struct {
	config     *larkcore.Config
	OpenapiLog *openapiLog // openapi_log
}

type openapiLog struct {
	service *SecurityAndComplianceService
}

//
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=list_data&project=security_and_compliance&resource=openapi_log&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/security_and_compliancev1/listData_openapiLog.go
func (o *openapiLog) ListData(ctx context.Context, req *ListDataOpenapiLogReq, options ...larkcore.RequestOptionFunc) (*ListDataOpenapiLogResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/security_and_compliance/v1/openapi_logs/list_data"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, o.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListDataOpenapiLogResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, o.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}