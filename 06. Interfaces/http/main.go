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
		os.Exit(1)
	}

	lw := logWriter{} //cutom type which implements the Writer interface

	//io.Copy(os.Stdout, resp.Body) //how does it output on the terminal without having to call normal fmt.Println()
	io.Copy(lw, resp.Body)

	//OR

	//bs := make([]byte, 4000)
	//n, err := resp.Body.Read(bs)
	//if err != nil {
	//	fmt.Println("erererer: ", err)
	//}
	//fmt.Println(n, string(bs))
}
func (logWriter) Write(bs []byte) (int, error) { //we can also create our own custom Write() method, notice how interfaces does let us do any logic, interface just had the role to moudlarirze our code structures
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes: ", len(bs))
	return len(bs), nil
}

/*Chatgpt result for example of interfaces inside interfaces*/
// package main
//
// import "fmt"
//
// // 2D Shape interface
//
//	type Shape2D interface {
//		Area() float64
//	}
//
// // 3D Shape interface
//
//	type Shape3D interface {
//		Volume() float64
//	}
//
// // Renderable shape interface
//
//	type Renderable interface {
//		Render() string
//	}
//
// // Interface for 2D shapes that can be rendered
//
//	type Renderable2D interface {
//		Shape2D
//		Renderable
//	}
//
// // Interface for 3D shapes that can be rendered
//
//	type Renderable3D interface {
//		Shape3D
//		Renderable
//	}
//
// // Circle type implementing 2D and renderable interfaces
//
//	type Circle struct {
//		Radius float64
//	}
//
//	func (c Circle) Area() float64 {
//		return 3.14159 * c.Radius * c.Radius
//	}
//
//	func (c Circle) Render() string {
//		return fmt.Sprintf("Rendering a 2D circle with radius %.2f", c.Radius)
//	}
//
// // Sphere type implementing 3D and renderable interfaces
//
//	type Sphere struct {
//		Radius float64
//	}
//
//	func (s Sphere) Volume() float64 {
//		return (4.0 / 3.0) * 3.14159 * s.Radius * s.Radius * s.Radius
//	}
//
//	func (s Sphere) Render() string {
//		return fmt.Sprintf("Rendering a 3D sphere with radius %.2f", s.Radius)
//	}
//
//	func main() {
//		circle := Circle{Radius: 2.5}
//		sphere := Sphere{Radius: 3.0}
//
//		var render2D Renderable2D = circle
//		var render3D Renderable3D = sphere
//
//		fmt.Println(render2D.Render())
//		fmt.Println("Area:", render2D.Area())
//
//		fmt.Println(render3D.Render())
//		fmt.Println("Volume:", render3D.Volume())

//	}

/*another eg. Chatgpt result for example of interfaces inside interfaces*/
//package main
//
//import "fmt"
//
//// First interface
//type Reader interface {
//	Read() string
//}
//
//// Second interface that embeds the first interface
//type FileReader interface {
//	//Read()
//	Reader
//	ReadFile() string
//}
//
//// Struct implementing the Reader interface
//type MyReader struct {
//	Data string
//}
//
//func (r MyReader) Read() string {
//	return r.Data
//}
//
//// Struct implementing the FileReader interface
//type MyFileReader struct {
//	FileData string
//}
//
//func (fr MyFileReader) Read() string {
//	return fr.FileData
//}
//
//func (fr MyFileReader) ReadFile() string {
//	return fr.FileData + " (from file)"
//}
//
//func main() {
//	reader := MyReader{Data: "Hello, Reader!"}
//	fileReader := MyFileReader{FileData: "Hello, FileReader!"}
//
//	fmt.Println(reader.Read())         // Output: Hello, Reader!
//	fmt.Println(fileReader.Read())     // Output: Hello, FileReader!
//	fmt.Println(fileReader.ReadFile()) // Output: Hello, FileReader! (from file)
//}
