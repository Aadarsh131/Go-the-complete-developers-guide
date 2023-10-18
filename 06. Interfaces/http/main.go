package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter []byte

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		//os.Exit(1)
	}
	//io.Writer()
	lw := logWriter{} //cutom type which implements the Writer interface (can be found in io.Writer())

	length, err := lw.Write([]byte{11, 22, 55, 98})
	//OR
	//length, err := lw.Write(logWriter{11, 22, 55, 98})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("size of []byte(logWriter) is: ", length)

	//io.Copy(os.Stdout, resp.Body) //how does it output on the terminal without having to call normal fmt.Println()
	_, _ = io.Copy(lw, resp.Body)

	//OR

	//bs := make([]byte, 4000)
	//n, err := resp.Body.Read(bs)
	//if err != nil {
	//	fmt.Println("erererer: ", err)
	//}
	//fmt.Println(n, string(bs))
}
func (logWriter) Write(bs []byte) (int, error) { //we can also create our own custom Write() method, notice how interfaces does let us do any logic, interface just had the role to moudlarirze our code structures
	return len(bs), nil
}
