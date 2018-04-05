package main

import "fmt"

type Coster interface{
  Cost() float64
}

func display(c Coster){
  fmt.Println("cost : ", c.Cost())
}

type Item struct{
  name string
  price float64
  quantity int
}

func (t Item) Cost() float64{
  return t.price * float64(t.quantity)
}

type DiscountItem struct{
  Item
  discountRate float64
}

func (t DiscountItem) Cost() float64{
  return t.Item.Cost() * (1.0 - t.discountRate/100)
}

type Items []Coster

func (ts Items) Cost() (c float64){
  for _, t := range ts{
    c += t.Cost()
  }
  return
}

func main(){
  shoes := Item{"shoes",10000,1}
  eventShoes := DiscountItem{Item{"eventShoes", 10000, 1},30.0}
  display(shoes)
  display(eventShoes)

  items := Items{shoes, eventShoes}
  display(items)
}
