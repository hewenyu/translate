package srt

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

// 解析SRT行文本
func (s *Srt) ParseSrtLine(line string) error {
	return nil
}

