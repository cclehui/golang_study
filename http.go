package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"
)

func main() {
	go func() {
		for {
			log.Println(time.Now().Format("2006-01-02 15:04:05"), ", ", GetDeployRootPath(true))

			time.Sleep(time.Second * 1)

		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}

func GetDeployRootPath(evalSymlinks bool) string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	if evalSymlinks {
		ex, err = filepath.EvalSymlinks(ex)
		if err != nil {
			panic(err)
		}
	}

	return filepath.Dir(ex)
}
