package auzer

import (
	"context"
	"fmt"
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
	cli := NewAuzerTranslator(resourceKey, endpoint, region)
	msg, err := cli.Translate(context.Background(), "hello", "en", "zh")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg)
}
