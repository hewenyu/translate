package srt

import (
	"testing"

	"github.com/hewenyu/translate/auzer"
)

// test read srt file

const filePath = "/workspaces/translate/example/100.srt"

const (
	// Azure 资源密钥
	resourceKey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	// Azure 翻译服务的终结点
	endpoint = "https://api.cognitive.microsofttranslator.com"
	// Azure 翻译服务的区域
	region = "eastus"
)

// test read srt file
func TestParseSrt(t *testing.T) {
	translateFactor := auzer.NewAuzerTranslator(resourceKey, endpoint, region)
	srtTranslator := NewSrtTranslator(filePath, "zh", "en", translateFactor, true)

	srtTranslator.Translate()

}
