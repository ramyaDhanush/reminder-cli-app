package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type idsFlag []string

func (ids idsFlag) String() string {
	return strings.Join(ids, ", ")
}

func (ids *idsFlag) Set(id string) error {
	*ids = append(*ids, id)
	return nil
}

type person struct {
	name string
	born time.Time
}

func (p person) String() string {
	return fmt.Sprintf("My name is %s, and I am %s", p.name, p.born)
}

func (p *person) Set(name string) error {
	p.name = name
	p.born = time.Now()
	return nil
}


func main() {
	var ids idsFlag
	var p person
	
	flag.Var(&ids, "id", "id to be appended to the list")
	flag.Var(&p, "name", "name of the person")
	flag.Parse()
	
	fmt.Println(ids)
	fmt.Println(p)
}