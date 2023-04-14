package srt

import "github.com/hewenyu/translate"

// 字幕翻译器
type SrtTranslator struct {
	filePath  string               // 字幕文件路径
	from      string               // 源语言
	to        string               // 目标语言
	t         translate.Translator // 翻译器
	srtSwitch bool                 // 是否输出双语字幕 true: 输出双语字幕 false: 只输出翻译后的字幕
}

// NewSrtTranslator 创建一个新的 Srt 翻译器
func NewSrtTranslator(filePath string, from string, to string, t translate.Translator, srtSwitch bool) *SrtTranslator {
	return &SrtTranslator{
		filePath:  filePath,
		from:      from,
		to:        to,
		t:         t,
		srtSwitch: srtSwitch,
	}
}

func (t *SrtTranslator) Translate() error {
	//分析、解析字幕
	srtParser := NewSrtParser(t.filePath)

	err := srtParser.Parse()
	if err != nil {
		return err
	}
	return nil
}
