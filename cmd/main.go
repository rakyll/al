package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/rakyll/al"
)

const size = 10

func main() {
	fmt.Println(al.Error())
	d := al.OpenDevice("")
	c := al.CreateContext(d, nil)
	fmt.Println(c)

	fmt.Println(al.MakeContextCurrent(c))

	sources := al.GenSources(1)
	bufs := al.GenBuffers(1)

	data, err := ioutil.ReadFile("8k16bitpcm.wav")
	if err != nil {
		log.Fatal(err)
	}

	bufs[0].BufferData(al.FORMAT_MONO16, data, 8000)
	sources[0].QueueBuffers(bufs[0])
	al.PlaySources(sources...)
	time.Sleep(10 * time.Second)

}
