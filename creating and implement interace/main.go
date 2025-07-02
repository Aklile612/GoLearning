package main
//Rectange Interface
type Rectangle interface{
	Permeter() float64
	Area() float64
}

type Calculate struct{
	length float64
	width float64
}

func (c Calculate) Permeter()float64{
	return (c.length+c.width)*2
}

func (c Calculate) Area() float64{
	return (c.length*c.width)
}
func main() {
	var r Rectangle
	
}