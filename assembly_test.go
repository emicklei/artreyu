package main

import (
	"os"
	"testing"

	"io/ioutil"

	"github.com/emicklei/assert"
	"gopkg.in/yaml.v2"
)

func Test_ReadYaml(t *testing.T) {
	f, err := os.Open("test-assembly.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	d := new(Assembly)
	data, _ := ioutil.ReadAll(f)
	err = yaml.Unmarshal(data, d)
	if err != nil {
		t.Fatal(err)
	}
	d.PostRead()

	out, _ := yaml.Marshal(d)
	t.Log(string(out))

	assert.That(t, "deps", d.Parts).Len(2)
}
