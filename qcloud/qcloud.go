package qcloud

import (
	"context"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

/*
 * 腾讯云翻译
 */

// Translator 腾讯云翻译器
type Translator struct {
	client    *tmt.Client
	projectId int64
}

// NewTranslator 创建腾讯云翻译器
func NewTranslator(cfg *QCloudConfig) (*Translator, error) {
	credential := common.NewCredential(
		cfg.SecretId,
		cfg.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tmt.tencentcloudapi.com"

	client, err := tmt.NewClient(credential, cfg.Region, cpf)
	if err != nil {
		return nil, fmt.Errorf("create tencent cloud client: %w", err)
	}

	return &Translator{
		client:    client,
		projectId: cfg.ProjectId,
	}, nil
}

// Translate 实现翻译接口
func (t *Translator) Translate(ctx context.Context, text string, sourceLang string, targetLang string) (string, error) {
	request := tmt.NewTextTranslateRequest()
	request.SourceText = common.StringPtr(text)
	request.Source = common.StringPtr(sourceLang)
	request.Target = common.StringPtr(targetLang)
	request.ProjectId = common.Int64Ptr(t.projectId)

	response, err := t.client.TextTranslateWithContext(ctx, request)
	if err != nil {
		return "", fmt.Errorf("translate text: %w", err)
	}

	return *response.Response.TargetText, nil
}
