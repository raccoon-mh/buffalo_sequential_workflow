package actions

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
)

func DEBUGRouteHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("_debug/buffaloRoute/index.html"))
}

func DEBUGWorkflowHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("flow/index.html"))
}

func DEBUGTablerMainHandler(c buffalo.Context) error {
	targetDir := "./templates/_debug/tabler"
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
	return c.Render(http.StatusOK, r.HTML("_debug/tabler/main.html"))
}

func DEBUGTablerHandler(c buffalo.Context) error {
	target := c.Param("target")
	target = "tabler/" + target + ".html"
	return c.Render(http.StatusOK, r.HTML(target))
}

func DEBUGSamplePageHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("_debug/tabler/main.html"))
}
