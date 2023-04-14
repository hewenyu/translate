package srt

import (
	"bufio"
	"fmt"
	"os"
)

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

// SrtAttr 字幕属性
type SrtAttr struct {
	line int // 字幕行数
}

// 字幕解析器
type SrtParser struct {
	filePath string    // 字幕文件路径
	data     []SrtLine // 字幕数据
}

// NewSrtParser 创建一个新的 Srt 解析器
func NewSrtParser(f string) *SrtParser {
	return &SrtParser{
		filePath: f,
	}
}

// Parse 解析字幕
func (p *SrtParser) Parse() error {
	// check .srt file is valid
	if !ValidSrtFile(p.filePath) {
		return os.ErrNotExist
	}
	// open file
	data, err := os.Open(p.filePath)
	if err != nil {
		return err
	}
	// close file
	defer data.Close()
	// read data
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}
