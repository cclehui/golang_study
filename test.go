package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type MyFunc func(a int) int

func (f MyFunc) add(b int) int {
	//func add(b int) {
	fmt.Println("1111111")
	return f(b) + 10
}

func main() {
	fmt.Println(os.Getppid())
	fmt.Println(os.Getpid())

	pName := os.Args[0]
	absName, _ := filepath.Abs(pName)

	fmt.Println(absName)
	fmt.Println(filepath.Dir(absName))
	fmt.Println(filepath.Base(absName))

	tempPath := "../aaa/cclehui_test.txt"
	fmt.Println(filepath.Ext(tempPath))

	globResult, _ := filepath.Glob("tmpl/[*]*.tmpl")
	fmt.Println(globResult)

	filepath.Walk("tmpl", func(path string, info os.FileInfo, err error) error {
		fmt.Printf("path: %s, %v\n", path, info)

		return nil
	})

	/*

		file, _ := os.Open("tmpl/index.tmpl")

		data, _ := ioutil.ReadAll(file)

		fmt.Printf("%s\n", data)

			logger := bgf_log.GetLogger("log test")

			go func() {
				_, file, line, _ := runtime.Caller(0)

				logger.Debug(file + "\t" + strconv.Itoa(line))
				fmt.Println(file + "\t" + strconv.Itoa(line))

			}()


			time.Sleep(time.Second * 1)

			logger.Debug("xxxxxxxxx")
	*/
}
