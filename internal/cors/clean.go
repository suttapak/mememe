package cors

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sync/errgroup"
)

func Clean() error {
	var (
		paths = []Item{
			{
				FilePath:     fmt.Sprintf("./internal/controller"),
				TemplatePath: "tmpl/controller.tmpl",
			},
			{
				FilePath:     fmt.Sprintf("./internal/service"),
				TemplatePath: "tmpl/service.tmpl",
			},
			{
				FilePath:     fmt.Sprintf("./internal/repository"),
				TemplatePath: "tmpl/repository.tmpl",
			},
			{
				FilePath:     fmt.Sprintf("./internal/route"),
				TemplatePath: "tmpl/router.tmpl",
			},
			{
				FilePath:     fmt.Sprintf("./internal/model"),
				TemplatePath: "tmpl/model.tmpl",
			},
			{
				FilePath:     fmt.Sprintf("./internal/filter"),
				TemplatePath: "tmpl/filter.tmpl",
			},
		}
	)
	egp := errgroup.Group{}
	for _, path := range paths {
		path := path
		egp.Go(func() error {
			// check file exists
			if _, err := os.Stat(path.FilePath); err != nil {
				log.Printf("file %s not found", path.FilePath)
				return nil
			}
			// remove file directory
			if err := os.RemoveAll(path.FilePath); err != nil {
				log.Printf("failed to remove file %s", path.FilePath)
				return nil
			}
			log.Printf("file %s removed", path.FilePath)

			return nil
		})
	}
	return egp.Wait()
}
