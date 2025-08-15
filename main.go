package main

import (
	"fmt"

	gopkgin "gopkg.in/yaml.v3"
	ghodss "github.com/ghodss/yaml"
)

const (
	first = `---
example: &example
  - "Hello"
  - "World!"

name: "hello"
commands:
  - << *example
  - "Oho!"
`
	second = `---
example: &example
  - "Hello"
  - "World!"

name: "hello"
commands:
  - <<: *example
  - "Oho!"
`
	third = `---
example: &example
  - "Hello"
  - "World!"

name: "hello"
commands:
  - *example
  - "Oho!"
`
	fourth = `---
example: &example
  - "Hello"
  - "World!"

name: "hello"
commands:
  - <<: *example
  - "Oho!"
`
	fifth = `---
example: &example
  - "Hello"
  - "World!"

name: "hello"
commands:
  *example
  - "Oho!"
`
)

type Unmarshaller func ([]byte, any) error

type Step struct {
	Name string `json:"string"`
	Commands []string `json:"commands"`
}

func tryUnmarshal(unmarshal Unmarshaller) {
	for i, d := range []string{first, second, third, fourth, fifth} {
		var step Step
		err := unmarshal([]byte(d), &step)
		if err != nil {
			fmt.Printf("unmarshal err: %d -> %v\n", i, err)
			continue
		}
		fmt.Printf("%d -> %#v\n", i, step)
	}
}


func main() {
	fmt.Println("===== gopkging =====")
	tryUnmarshal(gopkgin.Unmarshal)

	fmt.Println("===== ghodss =====")
	tryUnmarshal(ghodss.Unmarshal)
}
