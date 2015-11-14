package model

import (
	"testing"

	"github.com/emicklei/assert"
)

func TestLoadConfig(t *testing.T) {
	c, err := LoadConfig("test-config.yaml")
	assert.That(t, "loaded?", err).IsNil()
	assert.That(t, "api", c.Api).Equals(1)

	nexus := c.Repositories[1]
	assert.That(t, "url", nexus.URL).Equals("https://here.com/nexus")
	assert.That(t, "path", nexus.Path).Equals("/content/repositories")
	assert.That(t, "user", nexus.User).Equals("!")
	assert.That(t, "plugin", nexus.Plugin).Equals("nexus")
	assert.That(t, "password", nexus.Password).Equals("*")
}
