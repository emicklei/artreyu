package model

import (
	"testing"

	. "github.com/emicklei/assert"
)

func TestLoadConfig(t *testing.T) {
	c, err := LoadConfig("test-config.yaml")
	Assert(t, "load err", err).IsNil()
	Assert(t, "repo", c.Repository).Equals("nexus")
	Assert(t, "url", c.URL).Equals("https://here.com/nexus/content/repositories")
	Assert(t, "user", c.User).Equals("!")
	Assert(t, "password", c.Password).Equals("*")
	Assert(t, "osname", c.OSname).Equals("Darwin")
}
