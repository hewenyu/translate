package srt

import "github.com/hewenyu/translate"

// 字幕文件的结构
type Srt struct {
	Id                   int     //字幕自然行id
	Number               string  //字幕序号
	TimeStart            string  //字幕开始时间戳
	TimeStartSecond      float64 //字幕开始（秒）
	TimeStartMilliSecond int64   //字幕开始（毫秒）
	TimeEnd              string  //字幕结束时间戳
	TimeEndSecond        float64 //字幕结束（秒）
	TimeEndMilliSecond   int64   //字幕结束（毫秒）
	Text                 string  //字幕文本
	TranslateText        string  //翻译字幕文本
}

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
	// srt, err := ParseSrt(t.filePath)

	return nil
}
