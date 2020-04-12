package main

import (
	"fmt"
	"github.com/tirava/rflags/pkg/rflags"
	"os"
)

func main() {
	f := Flags{Values: []string{}}
	fmt.Println(rflags.ParseFlags(&f, os.Args[1:]))
	fmt.Println(f)
}

type Flags struct {
	Source string `rflag:"source,s,src"`
	Debug  bool   `rflag:"debug,d"`
	Output string
	Count  int
	Values []string `rflag:"value,val,v"`
}
