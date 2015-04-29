package main

import (
	"fmt"

	"github.com/emicklei/typhoon/local"
	"github.com/emicklei/typhoon/model"
)

func main() {
	ar, _ := model.LoadArtifact("./model/test-artifact.yaml")
	re := local.NewRepository("/tmp")
	fmt.Println(ar, re)

	err := re.Store(ar, "README.md")
	fmt.Println(err)
}
