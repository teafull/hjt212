// 将环境因子数据编码为呆传输的数据
package protocol

import (
	"bytes"
	"errors"
)

type HjtEncoder struct {
}

// Encoder Say hjt212 data packet is converted into a collection of data bytes to be sent.
func (h *HjtEncoder) Encoder(hjt212Cmd Hjt212Cmd) ([]byte, error) {
	// check data
	if !h.checkHjt212Package(hjt212Cmd) {
		return nil, errors.New("hjt212 data package check failed")
	}

	hjt212Byte := make([]byte, 4096)
	hjt212Buf := bytes.NewBuffer(hjt212Byte)
	hjt212Buf.WriteString(Head) // package of head

	// data pack

	hjt212Buf.WriteString(Tail) // package of tail

	return hjt212Byte, nil
}

// check data
func (h *HjtEncoder) checkHjt212Package(hjt212Cmd Hjt212Cmd) bool {
	if len(hjt212Cmd.QN) != 0 &&
		hjt212Cmd.ST != 0 &&
		hjt212Cmd.CN != 0 &&
		len(hjt212Cmd.MN) != 0 &&
		len(hjt212Cmd.PW) != 0 {
		return true
	}

	return false
}

// make data cmd package
func (h *HjtEncoder) makeCmdDataPkg(hjt212Cmd Hjt212Cmd) []byte {
	cmdDataPkgByte := make([]byte, 1024)
	cmdDataPkgBuf := bytes.NewBuffer(cmdDataPkgByte)

	cmdDataPkgBuf.WriteString("QN=")
	// TODO make hjt212Cmd to cmdDataPkgBuf

	return cmdDataPkgByte
}

// make cp data
func (h *HjtEncoder) makeCmdCp(params map[string]string) []byte {
	cmdCpByte := make([]byte, 512)
	cmdCpBuf := bytes.NewBuffer(cmdCpByte)

	cmdCpBuf.WriteString("CP=&&")
	// TODO make params to cmdCpBuf
	cmdCpBuf.WriteString("&&")

	return cmdCpByte
}
