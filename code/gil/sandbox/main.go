package main

func main() {
	x := "abc"
	y := &x
	z := &y
	println(y)
	println(&y)
	println(&z)
}
