package srt

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	filePath  string     // 字幕文件路径
	data      []*SrtLine // 字幕数据
	count     int        // 计数器
	start     bool       // 是否开始
	dataCount int        // 字幕数据计数器
	hasNum    bool       // 是否有序号
}

// NewSrtParser 创建一个新的 Srt 解析器
func NewSrtParser(f string) *SrtParser {
	return &SrtParser{
		filePath:  f,
		data:      make([]*SrtLine, 0),
		count:     0,
		dataCount: 0,
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
		// get line
		text, dataType, err := p.ParseLine(scanner.Text())
		if err != nil {
			return err
		}
		switch dataType {
		case SrtContentNum:
			// 这个字幕有序号
			p.hasNum = true
			if !p.start {
				p.start = true
			}
			// check line is number
			if IsNumber(text) {
				// add new line
				p.data = append(p.data, &SrtLine{
					ID:     p.count,
					Number: text,
				})
			}
		case SrtContentTime:
			if !p.start && !p.hasNum {
				// 部分字幕没有序号，直接开始时间
				p.start = true
				p.data = append(p.data, &SrtLine{
					ID:     p.count,
					Number: strconv.Itoa(p.dataCount),
				})
				p.dataCount++
			}
			// check line is time
			if IsTime(text) {
				// get time
				start, end, err := GetTime(text)
				if err != nil {
					return err
				}
				// add time
				p.data[len(p.data)-1].TimeStart = start
				p.data[len(p.data)-1].TimeEnd = end
				fmt.Println(start, end)
			}
		case SrtContentText:
			if !p.start {
				continue
			}
			// check line is text
			if text != "" {
				// add text
				p.data[len(p.data)-1].Text = append(p.data[len(p.data)-1].Text, text)
			}
		default:
			// do nothing
		}
		p.count++
	}
	return nil
}

func (p *SrtParser) ParseLine(line string) (text string, dataType SrtContent, err error) {
	text = strings.TrimSpace(line)
	// check line is empty
	if text == "" {
		return text, SrtContentNone, nil
	}
	// check line is number
	if IsNumber(text) {
		return text, SrtContentNum, nil
	}
	// check line is time
	if IsTime(text) {
		return text, SrtContentTime, nil
	}
	// check line is text
	return text, SrtContentText, nil
}

// IsNumber 判断字符串是否为数字
func IsNumber(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// IsTime 判断字符串是否为时间
func IsTime(s string) bool {
	// format like : 00:01:57.816 --> 00:02:01.576 A:middle align	X1:0.000000	X2:0.000000	Y1:0.000000	Y2:0.000000
	pattern := `^\d{2}:\d{2}:\d{2}\.\d{3} --> \d{2}:\d{2}:\d{2}\.\d{3}.*$`
	// check time format by regexp
	if ok, _ := regexp.MatchString(pattern, s); ok {
		return true
	}

	return false
}

func GetTime(s string) (start, end string, err error) {
	// format like : 00:01:57.816 --> 00:02:01.576 A:middle align	X1:0.000000	X2:0.000000	Y1:0.000000	Y2:0.000000
	pattern := `^(\d{2}:\d{2}:\d{2}\.\d{3}) --> (\d{2}:\d{2}:\d{2}\.\d{3}).*$`
	// check time format by regexp
	if ok, _ := regexp.MatchString(pattern, s); ok {
		// get time
		start = strings.Split(s, " --> ")[0]
		end = strings.Split(s, " --> ")[1]
		return start, end, nil
	}

	return "", "", fmt.Errorf("time format error")
}
