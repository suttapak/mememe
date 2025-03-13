package cors

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/sync/errgroup"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type (
	Item struct {
		FilePath, TemplatePath string
	}
	Data struct {
		RealName, UpperName, LowerName string
	}
)

//go:embed all:tmpl
var tmplFS embed.FS

func Create(name string) error {
	var (
		paths = []Item{
			{
				FilePath:     fmt.Sprintf("./internal/controller/%s.go", name),
				TemplatePath: "tmpl/controller.tmpl",
			},
			{
				FilePath:     fmt.Sprintf("./internal/service/%s.go", name),
				TemplatePath: "tmpl/service.tmpl",
			},
			{
				FilePath:     fmt.Sprintf("./internal/repository/%s.go", name),
				TemplatePath: "tmpl/repository.tmpl",
			},
			{
				FilePath:     fmt.Sprintf("./internal/route/%s.go", name),
				TemplatePath: "tmpl/router.tmpl",
			},
			{
				FilePath:     fmt.Sprintf("./internal/model/%s.go", name),
				TemplatePath: "tmpl/model.tmpl",
			},
			{
				FilePath:     fmt.Sprintf("./internal/filter/%s.go", name),
				TemplatePath: "tmpl/filter.tmpl",
			},
		}
		data = Data{
			RealName:  name,
			UpperName: convertSnakeToCamel(cases.Title(language.English).String(name)),
			LowerName: convertSnakeToCamel(strings.ToLower(name)),
		}
	)
	egp := errgroup.Group{}
	for _, path := range paths {
		path := path
		egp.Go(func() error {
			// check file exists
			if _, err := os.Stat(path.FilePath); err == nil {
				log.Printf("file %s already exists", path.FilePath)
				return nil
			}

			tmplContent, err := tmplFS.ReadFile(path.TemplatePath)
			if err != nil {
				return err
			}
			// parse template
			t, err := template.New(path.TemplatePath).Parse(string(tmplContent))
			if err != nil {
				return err
			}
			buf := new(bytes.Buffer)
			if err = t.Execute(buf, data); err != nil {
				return err
			}
			// create directory all
			dir := filepath.Dir(path.FilePath)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				if err := os.MkdirAll(dir, os.ModePerm); err != nil {
					return err
				}
			}

			// create file
			f, err := os.Create(path.FilePath)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = f.Write(buf.Bytes())
			if err != nil {
				return err
			}
			return nil
		})
	}
	return egp.Wait()
}
