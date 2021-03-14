// crc calculation and verification
package protocol

import "errors"

// verify crc od data, It is true when the verification is passed, otherwise it is false.
func VerifyCrc(dataSegment []byte, crc uint) (bool, error) {
	if len(dataSegment) == 0 {
		return false, errors.New("dataSegment empty")
	}

	calCrc, err := CalCrc(dataSegment)
	if err != nil {
		return false, err
	}

	if calCrc != crc {
		return false, nil
	}

	return true, nil
}

// calculate the crc of the packet dataSegment
func CalCrc(dataSegment []byte) (uint, error) {
	if len(dataSegment) == 0 {
		return 0, errors.New("dataSegment empty")
	}

	defer func() {
		if err := recover(); err != nil {
			//  exception catch
		}
	}()

	var crcResult, check uint
	crcResult = 0xFFFF
	for _, data := range dataSegment {
		crcResult = (crcResult >> 8) ^ uint(data)
		for j := 0; j < 8; j++ {
			check = crcResult & 0x0001
			crcResult >>= 1
			if check == 0x0001 {
				crcResult ^= 0xA001
			}
		}
	}
	return crcResult, nil
}
