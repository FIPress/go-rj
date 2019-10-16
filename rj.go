package rj

import (
	"errors"
	"io/ioutil"
	"log"
	"reflect"
)

// Load loads a file of given path and parse it into a node
func Load(path string) (node *Node, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	node, err = Parse(bytes)

	return
}

// Parse parses given bytes into a node
func Parse(input []byte) (node *Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			err = errors.New("parse RJ failed")
		}
	}()

	scanner := newScanner(input)
	scanner.scan()

	return scanner.root, scanner.error
}

// ParseString parses a string to a node
func ParseString(input string) (node *Node, err error) {
	return Parse([]byte(input))
}

// Marshal encodes a value to RJ bytes
func Marshal(v interface{}) []byte {
	e := newEncoder(v)
	return e.encode()
}

// Unmarshal decode RJ bytes to struct value
func Unmarshal(data []byte, v interface{}) (err error) {
	node, err := Parse(data)
	if err != nil {
		return
	}

	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errValueNotAssignable
	}
	return decodeNode(node, rv)
}
