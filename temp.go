package main

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
	"time"
)

func cclehuiListen() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// 服务器push通知，需要更新配置
			w.Write([]byte("xxxxxxxxxxx\n"))
		})

		srv := &http.Server{
			// TODO:端口待定
			Addr:    ":8082",
			Handler: mux,
		}
		fmt.Println("11111111111")

		addr := srv.Addr
		if addr == "" {
			addr = ":http"
		}
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			return
		}

		fmt.Println(ln)

		srv.Serve(ln)
		//server.ListenAndServe()
	}()

}

type Base struct {
}

func (this *Base) Name() {
	fmt.Println("this is Base, " + this.Sex())
}

func (this *Base) Sex() string {

	return "sex is from base"
}

type Child struct {
	Base
}

func (this *Child) Sex() string {

	return "sex is from child"
}

func main() {

	child := Child{}

	child.Name()

	os.Exit(0)
	cclehuiListen()
	cclehuiListen()

	for {
		time.Sleep(time.Second * 1)
	}

	if matched, _ := regexp.MatchString("^\\d+", "1sss11222"); matched {
		fmt.Println("aaaaaaaaaaaa")
	}

	r, _ := regexp.Compile("p([a-z]+)(ch)")
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

	re := regexp.MustCompile("a(x*?)b")
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))

}
