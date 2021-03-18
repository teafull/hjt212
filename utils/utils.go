package utils

import (
	"strconv"
	"time"
)

func GetQN() []byte {
	qN := time.Now().UnixNano() / 1e6
	return []byte(strconv.FormatInt(qN, 10))
}
