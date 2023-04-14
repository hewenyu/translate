package srt

import (
	"context"

	"github.com/hewenyu/translate"
)

// 字幕翻译器
type SrtTranslator struct {
	filePath    string               // 字幕文件路径
	from        string               // 源语言
	to          string               // 目标语言
	translation translate.Translator // 翻译器
	srtSwitch   bool                 // 是否输出双语字幕 true: 输出双语字幕 false: 只输出翻译后的字幕
}

// NewSrtTranslator 创建一个新的 Srt 翻译器
func NewSrtTranslator(filePath string, from string, to string, t translate.Translator, srtSwitch bool) *SrtTranslator {
	return &SrtTranslator{
		filePath:    filePath,
		from:        from,
		to:          to,
		translation: t,
		srtSwitch:   srtSwitch,
	}
}

func (t *SrtTranslator) Translate(ctx context.Context) error {
	//分析、解析字幕
	srtParser := NewSrtParser(t.filePath)

	err := srtParser.Parse()
	if err != nil {
		return err
	}

	for k := range srtParser.data {
		//翻译字幕
		translateText, err := t.translation.Translate(ctx, srtParser.data[k].Text[0], t.from, t.to)
		if err != nil {
			return err
		}

		//输出字幕
		if t.srtSwitch {
			srtParser.data[k].Text = append(srtParser.data[k].Text, translateText)
		} else {
			srtParser.data[k].Text = []string{translateText}
		}
	}
	return nil
}
