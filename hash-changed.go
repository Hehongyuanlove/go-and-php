package  main

import (
	"encoding/binary"
	"fmt"
	"math"
	"runtime"
	"time"
)

func main()  {
	start := time.Now()
	var bb []byte
	m:= make(map[string]int64,1e6)

	buf := make([]byte, binary.MaxVarintLen64)

	for i:=int64(math.MaxInt32);i<1e6+math.MaxInt32;i++ {

		n := binary.PutVarint(buf, i)
		bn := buf[:n]
		bb = append(bn,'_')

		value := time.Now().Unix()
		nm :=  binary.PutVarint(buf, value)
		bnm := buf[:nm]
		bb = append(bb,bnm...)

		key :=  string(bb)
		m[key] = value
	}
	fmt.Println(time.Since(start))
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.TotalAlloc , "B")
}

/**
* 1e6 次
* 135.1354ms
* 1,0627,1344 B
*/

/**
* 2e6 次
* 248.9743ms
* 1,5427,8416 B
*/