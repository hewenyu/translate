package auzer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type AuzerTranslator struct {
	resourceKey string       // Azure 资源密钥
	endpoint    string       // Azure 翻译服务的终结点
	region      string       // Azure 翻译服务的区域
	client      *http.Client // HTTP 客户端
}

// NewAuzerTranslator 创建一个新的 Auzer 翻译器
func NewAuzerTranslator(resourceKey string, endpoint string, region string) *AuzerTranslator {
	return &AuzerTranslator{
		resourceKey: resourceKey,
		endpoint:    endpoint,
		region:      region,
		client:      http.DefaultClient,
	}
}

func (t *AuzerTranslator) Translate(ctx context.Context, text string, sourceLang string, targetLang string) (string, error) {
	// TODO: 实现翻译逻辑
	uri := t.endpoint + "/translate?api-version=3.0"
	// create the query

	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	q := u.Query()

	q.Add("from", sourceLang)
	q.Add("to", targetLang)
	u.RawQuery = q.Encode()

	// Create an anonymous struct for your request body and encode it to JSON
	body := []struct {
		Text string
	}{
		{Text: text},
	}
	b, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	// Build the HTTP POST request
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	// Add required headers to the request
	req.Header.Add("Ocp-Apim-Subscription-Key", t.resourceKey)
	req.Header.Add("Ocp-Apim-Subscription-Region", t.region)
	req.Header.Add("Content-type", "application/json")

	// Call the Translator Text API
	resp, err := t.client.Do(req)
	if err != nil {
		return "", err
	}
	// Decode the JSON response
	var result []struct {
		Translations []struct {
			Text string
		}
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	// return the first translation
	if len(result) > 0 && len(result[0].Translations) > 0 {
		return result[0].Translations[0].Text, nil
	}

	return "", nil
}
