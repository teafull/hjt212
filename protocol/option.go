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

// Packet head and tail
const (
	Head = "##"
	Tail = "\r\n"
)
