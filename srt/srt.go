package srt

// 字幕内容类型
type SrtContent int

var (
	SrtContentText SrtContent = 0 // SrtContentText 字幕内容类型：文本
	SrtContentTime SrtContent = 1 // SrtContentTime 字幕内容类型：时间
	SrtContentNum  SrtContent = 2 // SrtContentNum 字幕内容类型：序号
	SrtContentNone SrtContent = 3 // SrtContentNone 字幕内容类型：无
)

// const MatchString = `^\d{2}:\d{2}:\d{2}\[,.]\d{3} --> \d{2}:\d{2}:\d{2}\[,.]\d{3}.*$`
const MatchString = `\d{2}:\d{2}:\d{2}[\.,]\d{3}\s-->\s\d{2}:\d{2}:\d{2}[\.,]\d{3}(\s[A-Z]:\w+\sX1:\d+\.\d+\sX2:\d+\.\d+\sY1:\d+\.\d+\sY2:\d+\.\d+)?`
