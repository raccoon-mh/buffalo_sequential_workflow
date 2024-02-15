package actions

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

func WorkflowHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("flow/index.html"))
}

func TablerMainHandler(c buffalo.Context) error {
	targetDir := "./templates/tabler"
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return err
	}
	list := []string{}
	for _, file := range files {
		if file.Name() == "main.html" {
			continue
		}
		list = append(list, strings.TrimRight(file.Name(), ".html"))
	}
	c.Set("files", list)
	return c.Render(http.StatusOK, r.HTML("tabler/main.html"))
}

func TablerHandler(c buffalo.Context) error {
	target := c.Param("target")
	target = "tabler/" + target + ".html"
	return c.Render(http.StatusOK, r.HTML(target))
}
