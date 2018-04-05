package main

import "fmt"

type shaper interface{
  area() float64
}

func describe(s shaper){
  fmt.Println("area : ", s.area())
}

type tri struct{ width, height float64}

func (t tri) area() float64{
  return t.width * t.height / 2.0
}

type rect struct{ width, height float64}

func (r rect) area() float64{
  return r.width * r.height
}

func main(){
  t := tri{3,4}
  describe(t)

  r := rect{3,4}
  describe(r)
}
