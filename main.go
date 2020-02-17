package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"

// 	"gopkg.in/yaml.v2"
// )

// type conf struct {
// 	book    int `yaml:"book"`
// 	release int `yaml:"release"`
// }

// // type release struct {
// // 	releaseyear string `yaml:"releaseyear"`
// // }

// func (c *conf) getConf() *conf {
// 	yamlFile, error := ioutil.ReadFile("testing.yaml")
// 	if error != nil {
// 		log.Fatal(error)
// 	}

// 	errors := yaml.Unmarshal(yamlFile, c)
// 	if errors != nil {
// 		log.Fatalf("Unmarshal: %v", error)
// 	}

// 	return c
// }

// func main() {
// 	fmt.Println("hello world")

// 	// yamlFile, error := ioutil.ReadFile("testing.yaml")
// 	// if error != nil {
// 	// 	log.Fatal(error)
// 	// }
// 	// var output *conf
// 	// error = yaml.Unmarshal(yamlFile, output)
// 	// if error != nil {
// 	// 	fmt.Println("Can't Unmarshal !")
// 	// }
// 	var output conf
// 	output.getConf()

// 	fmt.Println(output)

// }

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Hits int   `yaml:"hits"`
	Time int64 `yaml:"time"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("testing.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c conf
	c.getConf()

	fmt.Println(c)
}
