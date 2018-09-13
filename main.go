package main

import (
	"github.com/hashicorp/go-rootcerts"
)


func main(){
	rootcerts.LoadCAFile("ca.crt")
	rootcerts.LoadSystemCAs()
}
