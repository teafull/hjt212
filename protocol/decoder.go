// 解析环境因子数据到内存
package protocol

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type HjtDecoder struct {
}

// Decoder decoder bytes
func (h *HjtDecoder) Decoder(dadaPack []byte) (Hjt212Package, error) {
	if len(dadaPack) == 0 {
		return Hjt212Package{}, errors.New("data package empty")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	hjt212Package := Hjt212Package{Package: make(map[int]map[string]string)}

	headDataPacket := dadaPack[:2]
	tailDataPacket := dadaPack[len(dadaPack)-2 : len(dadaPack)]
	lenDataPacket, _ := strconv.Atoi(string(dadaPack[2:6]))
	dadaSegment := dadaPack[6 : len(dadaPack)-6]
	crcPackage := dadaPack[len(dadaPack)-6 : len(dadaPack)-2]
	dadaSegmentCrc, _ := strconv.ParseUint(string(crcPackage), 16, 0)

	crc, _ := VerifyCrc(dadaSegment, uint(dadaSegmentCrc))
	if string(headDataPacket) == Head && string(tailDataPacket) == Tail && lenDataPacket == len(dadaSegment) && crc {
		// op data segment
		hjt212Package.QN = ExtractDataByHeadAndTail(dadaSegment, []byte("QN="), []byte(";"))
		hjt212Package.ST, _ = strconv.Atoi(string(ExtractDataByHeadAndTail(dadaSegment, []byte("ST="), []byte(";"))))
		hjt212Package.CN, _ = strconv.Atoi(string(ExtractDataByHeadAndTail(dadaSegment, []byte("CN="), []byte(";"))))
		hjt212Package.PW = ExtractDataByHeadAndTail(dadaSegment, []byte("PW="), []byte(";"))
		hjt212Package.MN = ExtractDataByHeadAndTail(dadaSegment, []byte("MN="), []byte(";"))
		hjt212Package.Flag, _ = strconv.Atoi(string(ExtractDataByHeadAndTail(dadaSegment, []byte("Flag="), []byte(";"))))

		// PNum
		if bytes.Contains(dadaSegment, []byte("PNum=")) {
			hjt212Package.PNum, _ = strconv.Atoi(string(ExtractDataByHeadAndTail(dadaSegment, []byte("PNum="), []byte(";"))))
		} else {
			hjt212Package.PNum = 1
		}

		// get cp data, and get cp data package od multi
		cpData := ExtractDataByHeadAndTail(dadaSegment, []byte("CP=&&"), []byte("&&"))
		fmt.Println(string(cpData))
		//Package map[int] map[string] string  //  data pack，format is <PNo, <dataKey, dataValue>>, there is only one packet by default.
		for i := 0; i < hjt212Package.PNum; i++ {
			mapKey := make(map[string]string)
			cpField := bytes.Split(cpData, []byte(";"))
			for _, fieldKeyValue := range cpField {
				if len(fieldKeyValue) <= 0 {
					continue
				}

				if bytes.Contains(fieldKeyValue, []byte(",")) {
					// LA-td=50.1,LA-Flag=N
					cpSubField := bytes.Split(fieldKeyValue, []byte(","))
					for _, subFieldKeyValue := range cpSubField {
						if len(subFieldKeyValue) <= 0 {
							continue
						}

						key := subFieldKeyValue[:bytes.Index(subFieldKeyValue, []byte("="))]
						value := subFieldKeyValue[bytes.Index(subFieldKeyValue, []byte("="))+1:]
						mapKey[string(key)] = string(value)
					}
				} else {
					key := fieldKeyValue[:bytes.Index(fieldKeyValue, []byte("="))]
					value := fieldKeyValue[bytes.Index(fieldKeyValue, []byte("="))+1:]
					mapKey[string(key)] = string(value)
				}
			}

			//mapKey sort
			hjt212Package.Package[i+1] = mapKey
		}
	} else {
		return hjt212Package, errors.New("data package check err")
	}

	return hjt212Package, nil
}
