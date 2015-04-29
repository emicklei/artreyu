package model

import (
	"testing"

	"github.com/emicklei/assert"
	"gopkg.in/yaml.v2"
)

func Test_ReadYaml(t *testing.T) {

	d, err := LoadAssembly("test-assembly.yaml")
	assert.That(t, "load err", err).IsNil()

	out, _ := yaml.Marshal(d)
	t.Log(string(out))

	assert.That(t, "deps", d.Parts).Len(2)
}
