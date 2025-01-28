package auzer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// AuzerTranslator Azure翻译器
type AuzerTranslator struct {
	resourceKey string
	endpoint    string
	region      string
	client      *http.Client
}

// NewTranslator 创建一个新的 Azure 翻译器
func NewTranslator(config *AuzerConfig) (*AuzerTranslator, error) {
	if config.ResourceKey == "" {
		return nil, fmt.Errorf("resource key is required")
	}
	if config.Endpoint == "" {
		return nil, fmt.Errorf("endpoint is required")
	}
	if config.Region == "" {
		return nil, fmt.Errorf("region is required")
	}

	return &AuzerTranslator{
		resourceKey: config.ResourceKey,
		endpoint:    config.Endpoint,
		region:      config.Region,
		client:      &http.Client{},
	}, nil
}

// Translate 实现翻译接口
func (t *AuzerTranslator) Translate(ctx context.Context, text string, sourceLang string, targetLang string) (string, error) {
	uri := t.endpoint + "/translate"
	u, err := url.Parse(uri)
	if err != nil {
		return "", fmt.Errorf("parse endpoint URL: %w", err)
	}

	// 设置查询参数
	q := u.Query()
	q.Add("api-version", "3.0")
	if sourceLang != "auto" {
		q.Add("from", sourceLang)
	}
	q.Add("to", targetLang)
	u.RawQuery = q.Encode()

	// 准备请求体
	body := []struct {
		Text string `json:"text"`
	}{{Text: text}}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("marshal request body: %w", err)
	}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewReader(jsonBody))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	// 设置请求头
	req.Header.Set("Ocp-Apim-Subscription-Key", t.resourceKey)
	req.Header.Set("Ocp-Apim-Subscription-Region", t.region)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := t.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("translation failed with status %d: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var result []struct {
		Translations []struct {
			Text string `json:"text"`
		} `json:"translations"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	if len(result) == 0 || len(result[0].Translations) == 0 {
		return "", fmt.Errorf("empty translation result")
	}

	return result[0].Translations[0].Text, nil
}
