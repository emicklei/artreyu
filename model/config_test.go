package model

import (
	"testing"

	. "github.com/emicklei/assert"
)

func TestLoadConfig(t *testing.T) {
	c, err := LoadConfig("test-config.yaml")
	Assert(t, "loaded?", err).IsNil()
	Assert(t, "api", c.Api).Equals(1)

	nexus := c.Repositories[1]
	Assert(t, "url", nexus.URL).Equals("https://here.com/nexus")
	Assert(t, "path", nexus.Path).Equals("/content/repositories")
	Assert(t, "user", nexus.User).Equals("!")
	Assert(t, "password", nexus.Password).Equals("*")
}
