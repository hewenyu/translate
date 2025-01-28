package srt

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/hewenyu/translate/auzer"
)

// test read srt file

const filePath = "../test/test_data/test.srt"

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
	translateFactor, nil := auzer.NewTranslator(&auzer.AuzerConfig{
		ResourceKey: resourceKey,
		Endpoint:    endpoint,
		Region:      region,
	})
	srtTranslator := NewSrtTranslator(filePath, "en", "zh", translateFactor, true)

	err := srtTranslator.Translate(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestMathTime(t *testing.T) {
	// format like : 00:01:57.816 --> 00:02:01.576 A:middle align	X1:0.000000	X2:0.000000	Y1:0.000000	Y2:0.000000
	// format likt 00:00:00,000 --> 00:00:08,160

	var (
		tmp1 = "00:01:57.816 --> 00:02:01.576 A:middle align	X1:0.000000	X2:0.000000	Y1:0.000000	Y2:0.000000"
		tmp2 = "00:22:08,920 --> 00:22:09,560"
	)

	if ok, _ := regexp.MatchString(MatchString, tmp1); ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	if ok, _ := regexp.MatchString(MatchString, tmp2); ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

}

func TestFileName(t *testing.T) {
	name := strings.TrimSuffix(filePath, filepath.Ext(filePath))
	fmt.Println(name) // 输出：example
}
