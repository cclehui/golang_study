package main

import (
	"fmt"
	htmltemplate "html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("xxxxxxxxxx")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("11111111111")

	})

	//tmpl_files := make([]string)
	var tmpl_files []string
	filepath.Walk("tmpl", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".tmpl") {
			//absPath, _ := filepath.Abs(path)
			//tmpl_files = append(tmpl_files, absPath)
			tmpl_files = append(tmpl_files, path)
		}

		return nil
	})
	fmt.Printf("111111111, %v\n", tmpl_files)

	testFunc := func(filenames ...string) {
		for _, filename := range filenames {
			_, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Printf("eeeeeee, %v\n", err)
			}
		}
	}

	testFunc(tmpl_files...)

	r.GET("/template_raw", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "index.html", gin.H{"title": "ddddddddd"})

		//filenames, err := filepath.Glob("tmpl/i*?.tmpl")
		//fmt.Println(filenames)

		//t, err := template.ParseFiles("tmpl/index.tmpl")
		//t, err := template.ParseGlob("tmpl/*.tmpl")
		t, err := template.ParseFiles(tmpl_files...)
		//t, err := template.ParseFiles("tmpl/index.tmpl", "tmpl/user/index.tmpl")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(t)

		err = t.ExecuteTemplate(c.Writer, "index.tmpl", &map[string]string{"title": "33333333333"})
		//err = t.ExecuteTemplate(c.Writer, "user/index.tmpl", &map[string]string{"title": "33333333333"})

		fmt.Println(err)
		//t.Execute(c.Writer, &map[string]string{"title": "222222222222"})
		//t.Execute(c.Writer, "aaaaaaaaaaaaaa")
	})

	r.GET("/template_gin", func(c *gin.Context) {
		r.LoadHTMLGlob("tmpl/*.tmpl")
		//r.LoadHTMLGlob("tmpl/index.tmpl")
		//r.LoadHTMLFiles("./tmpl/user/index.tmpl")

		//c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "ddddddddd"})

		//c.HTML(http.StatusOK, "tmpl/user/index.tmpl ", gin.H{"title": "ddddddddd"})
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "ddddddddd"})
	})

	r.GET("/template_gin_extend", func(c *gin.Context) {

		DefaultConfig := gintemplate.TemplateConfig{
			Root:         "tmpl",
			Extension:    ".tmpl",
			Master:       "layouts/master",
			Partials:     []string{},
			Funcs:        make(htmltemplate.FuncMap),
			DisableCache: false,
			Delims:       gintemplate.Delims{Left: "{{", Right: "}}"},
		}

		r.HTMLRender = gintemplate.New(DefaultConfig)

		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "ddddddddd"})
	})

	//r.Static("/static", "/home/cclehui/go_workspace/src/babytree.com/test/static")
	r.Static("/static", "static")
	r.Run()
}
