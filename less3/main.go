package main

import (
	"fmt"
)

func main() {
	// fmt.Println("hello")
	var zzBox myBox
	zzBox.Put("firstSting", 3, 4)
	zzBox.Put("secondSting")
	zzBox.Put(true)
	zzBox.Put(nil)
	zzBox.PrintAll()
	zzBox.Drop(3)
	zzBox.PrintAll()
	zzBox.Drop("firstSting")
	zzBox.PrintAll()
	zzBox.Drop(false)
	zzBox.PrintAll()
	zzBox.Drop(true)
	zzBox.PrintAll()
	zzBox.Drop(nil)
	zzBox.PrintAll()

}

// write generic in go with interface{}

type Box interface {
	Put(...interface{})
	Drop(interface{})

	PrintAll()
}

type myBox struct {
	a []interface{}
}

func (c *myBox) Put(b ...interface{}) {
	for _, val := range b {
		c.a = append(c.a, val)
	}
}

func (c *myBox) PrintAll() {
	fmt.Println(c.a)
}

func (c *myBox) Drop(d interface{}) {

	b := c.a

	for i, val := range b {
		switch v := val.(type) {
		case int:
			s, ok := d.(int)
			if ok && s == v {
				c.a = append(b[:i], b[i+1:]...)
			}
		case bool:
			s, ok := d.(bool)
			if ok && s == v {
				c.a = append(b[:i], b[i+1:]...)
			}
		case string:
			s, ok := d.(string)
			if ok && s == v {
				c.a = append(b[:i], b[i+1:]...)
			}
		case nil:
			if v == nil {
				c.a = append(b[:i], b[i+1:]...)
			}
		default:
			fmt.Println("type unknown")
		}
	}
}
