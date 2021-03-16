// 将环境因子数据编码为呆传输的数据
package protocol

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type HjtEncoder struct {
}

// Encoder Say hjt212 data packet is converted into a collection of data bytes to be sent.
func (h *HjtEncoder) Encoder(hjt212Cmd Hjt212Cmd) ([]byte, error) {
	// check data
	if !h.checkHjt212Package(hjt212Cmd) {
		return nil, errors.New("hjt212 data package check failed")
	}

	hjt212Buf := bytes.NewBuffer(nil)
	hjt212Buf.WriteString(Head) // package of head

	// data pack
	dataPkg := h.makeCmdDataPkg(hjt212Cmd)
	hjt212Buf.WriteString(fmt.Sprintf("%04d", len(dataPkg))) // data package length
	hjt212Buf.Write(dataPkg)
	crcResult, _ := CalCrc(dataPkg)
	hjt212Buf.WriteString(fmt.Sprintf("%04X", crcResult)) // data package crc

	hjt212Buf.WriteString(Tail) // package of tail

	return hjt212Buf.Bytes(), nil
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
	cmdDataPkgBuf := bytes.NewBuffer(nil)

	// TODO make hjt212Cmd to cmdDataPkgBuf
	cmdDataPkgBuf.WriteString("QN=")
	cmdDataPkgBuf.Write(hjt212Cmd.QN)
	cmdDataPkgBuf.WriteString(";ST=")
	cmdDataPkgBuf.WriteString(strconv.Itoa(hjt212Cmd.ST))
	cmdDataPkgBuf.WriteString(";CN=")
	cmdDataPkgBuf.WriteString(strconv.Itoa(hjt212Cmd.CN))
	cmdDataPkgBuf.WriteString(";PW=")
	cmdDataPkgBuf.Write(hjt212Cmd.PW)
	cmdDataPkgBuf.WriteString(";MN=")
	cmdDataPkgBuf.Write(hjt212Cmd.MN)
	cmdDataPkgBuf.WriteString(";Flag=")
	cmdDataPkgBuf.WriteString(strconv.Itoa(hjt212Cmd.Flag))
	cmdDataPkgBuf.WriteString(";")
	cmdDataPkgBuf.Write(h.makeCmdCp(hjt212Cmd.Params))

	return cmdDataPkgBuf.Bytes()
}

// make cp data
func (h *HjtEncoder) makeCmdCp(params map[string]string) []byte {
	cpBuf := bytes.NewBuffer(nil)

	// TODO make params to cmdCpBuf
	i := 0
	cpBuf.WriteString("CP=&&")
	for pKey, pValue := range params {
		if len(params) > 1 && i != 0 {
			cpBuf.WriteString(",")
		}
		i++

		cpBuf.WriteString(pKey)
		cpBuf.WriteString("=")
		cpBuf.WriteString(pValue)
	}
	cpBuf.WriteString("&&")

	return cpBuf.Bytes()
}
