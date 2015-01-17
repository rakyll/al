package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rakyll/al"
)

const size = 100000

func main() {
	fmt.Println(al.Error())
	d := al.OpenDevice("")
	c := al.CreateContext(d, []int32{})
	fmt.Println(c)

	fmt.Println(al.MakeContextCurrent(c))

	sources := al.GenSources(1)

	bufs := al.GenBuffers(1)
	fmt.Println(bufs)

	data := make([]int32, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int31() * 100
	}

	bufs[0].BufferData(al.FORMAT_STEREO16, data, 4000)
	sources[0].QueueBuffers(bufs[0])
	al.PlaySources(sources...)

	time.Sleep(10 * time.Second)

}
