package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"babytree.com/bgf/bgf_util/babytree/uidencoder"
)

func main() {

	//uidMap := uidencoder.GetUserIdUidMap()
	//fmt.Println(uidMap)
	//tempString := "abcde"
	//fmt.Println(tempString[2:3])

	//os.Exit(0)

	//for i := 1; i <= 10; i++ {
	rand.Seed(time.Now().UnixNano())
	//base := rand.Intn(4000000000)
	//base := 100000
	//for i := 0; i <= 4000000000; i++ {
	for i := 150000001; i <= 4000000000; i = i + 33 {
		//for i := base; i <= base+40; i++ {

		uid := uidencoder.Encode(i)
		userId := uidencoder.Decode(uid)

		if userId != i {
			fmt.Printf("bad, %v, %v, %v\n", i, uid, userId)
			os.Exit(1)
		}

		fmt.Printf("%v\t%v\n", i, uid)

	}

}
