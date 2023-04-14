package srt

// SrtLine 字幕行结构
type SrtLine struct {
	ID                   int      //字幕自然行id
	Number               string   //字幕序号
	TimeStart            string   //字幕开始时间戳
	TimeStartSecond      float64  //字幕开始（秒）
	TimeStartMilliSecond int64    //字幕开始（毫秒）
	TimeEnd              string   //字幕结束时间戳
	TimeEndSecond        float64  //字幕结束（秒）
	TimeEndMilliSecond   int64    //字幕结束（毫秒）
	Text                 []string //字幕文本
}

// 字幕解析器
type SrtParser struct {
}

func ParseSrt(filePath string) (*Srt, error) {
	// read file
	// parse srt
	// return srt
	return nil, nil
}
