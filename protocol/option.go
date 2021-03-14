// 定义数据结构,
/*
-------------------------------------------------------------------------------------
| 包头(2byte,##) | 数据包长度(4byte) | 数据包(0-1024byte) | CRC(4byte) | 包尾(2byte\r\n) |
-------------------------------------------------------------------------------------

// 数据包
---------------------------------------
|  字段       | 标识 | 字节数  | 描述     |
---------------------------------------
| 请求编号     | QN  | 20byte | ms时间戳 |
---------------------------------------
| 系统编号     | ST  | 5byte  |         |
---------------------------------------
| 命令编号     | CN  | 7byte  |         |
---------------------------------------
| 设备唯一标识  | MN  | 14byte |         |
---------------------------------------
| 密码        | PW   | 6byte |         |
---------------------------------------
| 总包号      | PNUM | 4byte  |        |
---------------------------------------
| 包号        | PNO  | 4byte |         |
---------------------------------------
| 数据标识     | Flag | 3byte |         |
---------------------------------------
*/
package protocol

import (
	"bytes"
)

// Packet head and tail
const (
	Head = "##"
	Tail = "\r\n"
)

// hjt212 data package
type Hjt212Package struct {
	QN      []byte                    // 请求编号, ms级时间戳
	ST      int                       // 系统编号
	CN      int                       // 命令编号
	MN      []byte                    // 设备唯一标识
	PW      []byte                    // 密码
	PNum    int                       // 总包号
	Flag    int                       // Flag标志
	Package map[int]map[string]string //  data pack，format is <PNo, <dataKey, dataValue>>, there is only one packet by default.
}

// Hjt212Encoder related interfaces for performing coding
type Hjt212Encoder interface {
	Encoder(hjt212Package Hjt212Package) ([]byte, error)
}

// Hjt212Decoder perform decoding related interfaces
type Hjt212Decoder interface {
	Decoder(dadaPack []byte) (Hjt212Package, error)
}

func ExtractDataByHeadAndTail(dataSegment, head, tail []byte) []byte {
	if len(head) == 0 && len(tail) == 0 {
		return dataSegment
	}

	if len(head) == 0 {
		tailIndex := bytes.Index(dataSegment, tail)
		if tailIndex < 0 {
			tailIndex = len(dataSegment)
		}
		return dataSegment[:tailIndex]
	}

	if len(tail) == 0 {
		headIndex := bytes.Index(dataSegment, head)
		if headIndex < 0 {
			headIndex = 0 // not found begin index is 0
		} else {
			headIndex = headIndex + len(head)
		}
		return dataSegment[headIndex:]
	}

	headIndex := bytes.Index(dataSegment, head)
	if headIndex < 0 {
		headIndex = 0 // not found begin index is 0
	} else {
		headIndex = headIndex + len(head)
	}
	startHead := dataSegment[headIndex:]

	tailIndex := bytes.Index(startHead, tail)
	if tailIndex < 0 {
		tailIndex = len(startHead)
	}
	return startHead[:tailIndex]
}
