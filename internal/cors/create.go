package cors

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/sync/errgroup"
)

type (
	Item struct {
		FilePath     string
		TemplatePath string
	}
	Data struct {
		UpperName string
		LowerName string
	}
)

func Create(name string) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	var (
		paths = []Item{
			{
				FilePath:     fmt.Sprintf("%s/internal/controller/%s.go", workingDir, name),
				TemplatePath: fmt.Sprintf("%s/internal/tmpl/controller.tmpl", workingDir),
			},
			{
				FilePath:     fmt.Sprintf("%s/internal/service/%s.go", workingDir, name),
				TemplatePath: fmt.Sprintf("%s/internal/tmpl/service.tmpl", workingDir),
			},
			{
				FilePath:     fmt.Sprintf("%s/internal/repository/%s.go", workingDir, name),
				TemplatePath: fmt.Sprintf("%s/internal/tmpl/repository.tmpl", workingDir),
			},
		}
		data = Data{
			UpperName: strings.Title(name),
			LowerName: strings.ToLower(name),
		}
	)
	egp := errgroup.Group{}
	for _, path := range paths {
		path := path
		egp.Go(func() error {
			// parse template
			t, err := template.ParseFiles(path.TemplatePath)
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
