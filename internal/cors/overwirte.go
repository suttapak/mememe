// overwirte module.go file

/*
package logger

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(newAppLogger),
)

*/

package cors

import (
	"bytes"
	"fmt"
	"github.com/suttapak/mememe/internal/utils"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type (
	ModuleDto struct {
		PackageName string
		UpperName   string
	}
)

var (
	controller = "controller"
	service    = "service"
	repository = "repository"
)

func Overwirte(name string) error {
	if !utils.IsHaveControllerModuleFile() {
		// create controller module file
		if err := createModuleFile(name, controller); err != nil {
			return err
		}
	}
	if !utils.IsHaveServiceModuleFile() {
		// create service module file
		if err := createModuleFile(name, service); err != nil {
			return err
		}
	}
	if !utils.IsHaveRepositoryeModuleFile() {
		// create repository module file
		if err := createModuleFile(name, repository); err != nil {
			return err
		}
	}

	var (
		paths = []string{
			fmt.Sprintf("./internal/controller/module.go"),
			fmt.Sprintf("./internal/service/module.go"),
			fmt.Sprintf("./internal/repository/module.go"),
		}
	)
	var (
		newModule  = fmt.Sprintf("),\n\tfx.Provide(New%s),\n)", cases.Title(language.English).String(name))
		moduleName = fmt.Sprintf("fx.Provide(New%s)", cases.Title(language.English).String(name))
	)
	egp := errgroup.Group{}
	for _, path := range paths {
		path := path
		egp.Go(func() error {
			// read file
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			text := string(content)
			if strings.Contains(text, moduleName) {
				return nil
			}
			re := regexp.MustCompile(`\),\n\)`)
			text = re.ReplaceAllString(text, newModule)

			// write file
			if err := os.WriteFile(path, []byte(text), 0644); err != nil {
				return err
			}

			return nil
		})
	}
	if err := egp.Wait(); err != nil {
		return err
	}

	return nil
}

func createModuleFile(name, pkgname string) error {
	tmplContent, err := tmplFS.ReadFile("tmpl/module.tmpl")
	if err != nil {
		return err
	}
	// parse template
	t, err := template.New("tmpl/module.tmpl").Parse(string(tmplContent))
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, ModuleDto{
		PackageName: pkgname,
		UpperName:   cases.Title(language.English).String(name),
	}); err != nil {
		return err
	}
	// create file
	// create file
	f, err := os.Create(fmt.Sprintf("./internal/%s/module.go", pkgname))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}
