// a program to grab memory over time

package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	//"io"
	"time"
)

var memsize, step, max int
var interval, now, then time.Duration

var myMemory = make([]byte, step)

//var myMemory bytes.Buffer

func usage() {
	fmt.Println("a program to suck up memory")
}

func init() {
	flag.IntVar(&memsize, "step", 1, "step size for memory increases")
	flag.IntVar(&max, "max", 5, "maximum memory to allcate")
	flag.DurationVar(&interval, "interval", 1, "time between memory steps")
	flag.Parse()
}

func fill(b bytes.Buffer) {
	b.Grow(step)
	slice := b.Bytes()
	bytesWritten, err := rand.Read(slice)
	if bytesWritten != len(slice) || err != nil {
		fmt.Println("error:", err)
		return
	}
	time.Sleep(interval)
}

func main() {
	//myMemory = make([]byte, step)
	myBuffer := bytes.NewBuffer(myMemory)
	for memsize <= max {
		//myMemory = append(myMemory, [step]byte)
		fill(*myBuffer)
		// the accounting is prolly wrong
	}
	usage()
}
