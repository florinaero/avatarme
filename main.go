package main

import (
	"avatarme/avt"
	"crypto"
	"fmt"
)

// 1. Generate hash based on name
// 2. Generate image based on hash
func main() {
	fmt.Println("Start main")
	name := "Cristian"
	ptr_name := &name
	avt.Init()
	avt.GenerateHash(ptr_name, crypto.MD5)
	avt.GenerateHash(ptr_name, crypto.SHA256)
	
}