package auzer

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

const (
	// Azure 资源密钥
	resourceKey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	// Azure 翻译服务的终结点
	endpoint = "https://api.cognitive.microsofttranslator.com"
	// Azure 翻译服务的区域
	region = "eastus"
)

// create test function
func TestAuzerTranslator_Translate(t *testing.T) {
	cli := func() *AuzerTranslator {
		var config *AuzerConfig = &AuzerConfig{
			ResourceKey: resourceKey,
			Endpoint:    endpoint,
			Region:      region,
		}
		return &AuzerTranslator{resourceKey: config.ResourceKey, endpoint: config.Endpoint, region: config.Region, client: http.DefaultClient}
	}()
	msg, err := cli.Translate(context.Background(), "hello", "en", "zh")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg)
}
