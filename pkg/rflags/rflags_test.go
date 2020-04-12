package rflags

import (
	"reflect"
	"testing"
)

// source="./data" debug output=out count=555 count=666 v=aaa v=bbb v=ccc

type OutFlags struct {
	Source string `rflag:"source,s,src"`
	Debug  bool   `rflag:"debug,d"`
	Output string
	Count  int
	Values []string `rflag:"value,val,v"`
}

func TestParseFlags(t *testing.T) {
	f := OutFlags{}
	args := []string{`source="./data"`, `debug`, `output=out`, `count=555`, `count=666`, `v=aaa`, `v=bbb`, `v=ccc`}
	if err := ParseFlags(&f, args); err != nil {
		t.Error(err)
	}
	if f.Source != "./data" {
		t.Errorf("Source should be: %s, got: %s", "./data", f.Source)
	}
	if f.Debug != true {
		t.Errorf("Debug should be: %v, got: %v", true, f.Debug)
	}
	if f.Output != "out" {
		t.Errorf("Output should be: %s, got: %s", "out", f.Output)
	}
	if f.Count != 666 {
		t.Errorf("Count should be: %d, got: %d", 666, f.Count)
	}

	values := []string{"aaa", "bbb", "ccc"}
	if !reflect.DeepEqual(f.Values, values) {
		t.Errorf("Values should be: %v, got: %v", values, f.Values)
	}
}
