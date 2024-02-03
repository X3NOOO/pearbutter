package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"regexp"
)

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

func printStruct(s interface{}) {
	structType := reflect.TypeOf(s)

	if structType.Kind() != reflect.Struct {
		fmt.Println("Input is not a struct.")
		return
	}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fmt.Printf("Field Name: %s\n", field.Name)
		fmt.Printf("Field Type: %s\n", field.Type)
		fmt.Printf("Field Tag: %s\n", field.Tag)
		fmt.Println("--------")
	}
}

func printStructField(field reflect.StructField) {
	fmt.Printf("Field Name: %s\n", field.Name)
	fmt.Printf("Field Type: %s\n", field.Type)
	fmt.Printf("Field Tag: %s\n", field.Tag)
	fmt.Println("--------")
}

func printType(t reflect.Type) {
	fmt.Printf("Type Name: %s\n", t.Name())
	fmt.Printf("Type Kind: %s\n", t.Kind())

	// If it's a struct type, print the fields
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fmt.Printf("Field Name: %s\n", field.Name)
			fmt.Printf("Field Type: %s\n", field.Type)
			fmt.Printf("Field Tag: %s\n", field.Tag)
			fmt.Println("--------")
		}
	}
}

/*
Get all the new items since the last message from the RSS feed

Args:
	config: The bot configuration

Returns:
	[]string: The items in the RSS feed
*/
func ParseRss(config *BotConfig) ([]string, error) {
	rss, err := fetch(config.RssURL)
	if err != nil {
		return nil, err
	}

	var sf []reflect.StructField

	log.Println("formatting:", config.Formatting)
	formats := regexp.MustCompile(`%([\w>]*)%`).FindAllStringSubmatch(config.Formatting, -1)
	log.Println("formats:", formats)
	for i, f := range formats {
		field := reflect.StructField{Name: fmt.Sprintf("Field_%d", i), Type: reflect.TypeOf(""), Tag: reflect.StructTag(fmt.Sprintf("`xml:\"%s\"`", f[1]))}
		sf = append(sf, field)
	}

	for i, f := range sf {
		log.Println("field:", i)
		printStructField(f)
	}

	st := reflect.StructOf(sf)

	log.Println("type")
	printType(st)

	s := reflect.New(st)

	log.Println("struct")
	printStruct(s) // ERROR: what the fuck????

	err = xml.Unmarshal(rss, &s)
	if err != nil {
		return nil, err
	}

	log.Printf("struct: %v\n", s)

	return nil, nil		
}
