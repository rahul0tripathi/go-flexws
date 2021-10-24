package gopool

import (
	"github.com/panjf2000/ants/v2"
)

var ConnHandlePool *ants.Pool
var poolErr error

func init() {
	ConnHandlePool, poolErr = ants.NewPool(5000)
	if poolErr != nil {
		panic(poolErr)
	}
}
