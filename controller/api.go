package controller

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"net/http"
	"io/ioutil"
	"strings"
)

const ApiUrl = "https://cdnjs.boiko.cn"
const LibrariesListUrl = "/libraries"

func LibrariesList(c *gin.Context) {
	req := struct {
		Search string `form:"search"`
		Fields string `form:"fields"`
	}{}
	if c.ShouldBind(&req) != nil {
		c.String(400, "Request parameter error")
		return
	}
	u, err := url.Parse(ApiUrl + LibrariesListUrl)

	if err != nil {
		c.String(500, "Server error")
		return
	}
	q := u.Query()
	if req.Search != "" {
		q.Set("search", req.Search)
	}
	if req.Fields != "" {
		q.Set("fields", req.Fields)
	}
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())

	if err != nil {
		c.String(500, "Server error")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(500, "Server error")
		return
	}
	content := Replace(string(body))
	c.String(200, content)
}
func LibrariesInfo(c *gin.Context) {
	req := struct {
		Fields string `form:"fields"`
	}{}

	if c.ShouldBind(&req) != nil {
		c.String(400, "Request parameter error")
		return
	}
	name := c.Param("name")
	u, err := url.Parse(ApiUrl + "/libraries/" + name)

	if err != nil {
		c.String(500, "Server error")
		return
	}
	q := u.Query()
	if req.Fields != "" {
		q.Set("fields", req.Fields)
	}
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())

	if err != nil {
		c.String(500, "Server error")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(500, "Server error")
		return
	}
	content := Replace(string(body))
	c.String(200, content)
}

func Replace(str string) string {
	content := strings.Replace(str, "cdnjs.cloudflare.com/ajax/libs", "cdn.boiko.cn", -1)
	return content
}
