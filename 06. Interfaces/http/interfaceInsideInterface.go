/*Chatgpt result for example of interfaces inside interfaces*/

package main

import "fmt"

// 2D Shape interface

type Shape2D interface {
	Area() float64
}

// 3D Shape interface

type Shape3D interface {
	Volume() float64
}

// Renderable shape interface

type Renderable interface {
	Render() string
}

// Interface for 2D shapes that can be rendered

type Renderable2D interface {
	Shape2D
	Renderable
}

// Interface for 3D shapes that can be rendered
// any type that satisfies/implements both Shape3D and Renderable interface is a valid Renderable3D interface

type Renderable3D interface {
	Shape3D
	Renderable
	getStats() string
}

// Circle type implementing 2D and renderable interfaces, i.e it also satisfies Renderable2D interface

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Render() string {
	return fmt.Sprintf("Rendering a 2D circle with radius %.2f", c.Radius)
}

// Sphere type implementing 3D and renderable interfaces, i.e, it is a valid Renderable3D interface

type Sphere struct {
	Radius float64
}

func (s Sphere) Volume() float64 {
	return (4.0 / 3.0) * 3.14159 * s.Radius * s.Radius * s.Radius
}

func (s Sphere) Render() string {
	return fmt.Sprintf("Rendering a 3D sphere with radius %.2f", s.Radius)
}

func main() {
	circle := Circle{Radius: 2.5}
	sphere := Sphere{Radius: 3.0}

	var render2D Renderable2D = circle
	//var render3D Renderable3D = sphere
	render3D := sphere // The actual difference here and the previous line (`var render3D Renderable3D = sphere`) is that sphere must stasfies Renderable3D interface in the previous case

	fmt.Println(render2D.Render())
	fmt.Println("Area:", render2D.Area())

	fmt.Println(render3D.Render())
	fmt.Println("Volume:", render3D.Volume())

}

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
//	//Read() string
//	//OR
//	Reader
//
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
//	var reader Reader = MyReader{Data: "Hello, Reader!"}
//	var fileReader FileReader= MyFileReader{FileData: "Hello, FileReader!"}
//
//	fmt.Println(reader.Read())         // Output: Hello, Reader!
//	fmt.Println(fileReader.Read())     // Output: Hello, FileReader!
//	fmt.Println(fileReader.ReadFile()) // Output: Hello, FileReader! (from file)
//}
