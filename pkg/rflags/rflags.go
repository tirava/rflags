package rflags

import (
	"fmt"
	"reflect"
	"strings"
)

// source="./data" debug output=out

func ParseFlags(str interface{}, args []string) error {
	strT := reflect.TypeOf(str)
	if strT.Kind() != reflect.Ptr {
		return fmt.Errorf("pointer needed")
	}

	aliases, err := getAliases(str)
	if err != nil {
		return err
	}

	flags, err := getFlags(args)
	if err != nil {
		return err
	}

	for flag := range flags {
		fieldNum, exists := aliases[flag]
		if !exists {
			return fmt.Errorf("unexpected flag: %s", flag)
		}

		fieldT := reflect.TypeOf(str).Elem().Field(fieldNum)
		fieldV := reflect.ValueOf(str).Elem().Field(fieldNum)

		switch fieldT.Type.Kind() {
		case reflect.String:
			fieldV.SetString(flags[flag])
		case reflect.Bool:
			fieldV.SetBool(true)
		default:
			return fmt.Errorf("unexpected field type: %s", fieldT.Type.Kind().String())
		}
	}

	return nil
}

type Aliases map[string]int

func getAliases(str interface{}) (Aliases, error) {
	aliases := Aliases{}
	strT := reflect.TypeOf(str).Elem()

	for i := 0; i < strT.NumField(); i++ {
		fieldT := strT.Field(i)
		alternativesStr := fieldT.Tag.Get("rflag")
		if alternativesStr == "" {
			alternativesStr = strings.ToLower(fieldT.Name)
		}
		alternatives := strings.Split(alternativesStr, ",")

		for _, alt := range alternatives {
			if _, exists := aliases[alt]; exists {
				return nil, fmt.Errorf("duplicated alias %s on field %s", alt, fieldT.Name)
			}
			aliases[alt] = i
		}
	}

	return aliases, nil
}

type Flags map[string]string

func getFlags(args []string) (Flags, error) {
	flags := Flags{}

	for _, arg := range args {
		parts := strings.Split(arg, "=")
		name := parts[0]
		val := ""
		if len(parts) > 1 {
			val = parts[1]
		}

		if _, exists := flags[name]; exists {
			return nil, fmt.Errorf("dublicate flag: %s", name)
		}

		val = strings.Trim(val, `"`)
		flags[name] = val
	}

	return flags, nil
}
