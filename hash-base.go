package  main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main()  {
	start := time.Now()
	var bb []byte
	m:= make(map[string]int64,1e6)

	for i:=int64(math.MaxInt32);i<1e6+math.MaxInt32;i++ {
		value := time.Now().Unix()
		bb:=bb[:0]
		bb = AppendInt(bb,i)
		bb = append(bb,'_')
		bb= AppendInt(bb,value)
		key :=  string(bb)
		m[key] = value
	}
	fmt.Println(time.Since(start))
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.TotalAlloc , "B")
}

func AppendInt(dst []byte,n int64) []byte  {
	var b [20]byte
	buf := b[:]
	i:= len(buf)
	var q int64
	for n>=10 {
		i--
		q=n/10
		buf[i] = '0'+byte(n-q*10)
		n=q
	}
	i--
	buf[i] = '0'+byte(n)
	return append(dst,buf[i:]...)
}

/**
* 350.0016ms
* 1,3828,9792 B
*
* 1.0810163s
* 342131952 B
*/