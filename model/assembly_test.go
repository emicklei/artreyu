package model

import (
	"testing"

	"github.com/emicklei/assert"
)

func Test_ReadYaml(t *testing.T) {
	d, err := LoadAssembly("test-assembly.yaml")
	assert.That(t, "load err", err).IsNil()
	assert.That(t, "deps", d.Parts).Len(2)
}

func Test_LoadArtifactFromAssemblyFile(t *testing.T) {
	a, err := LoadArtifact("test-assembly.yaml")
	assert.That(t, "load err", err).IsNil()
	assert.That(t, "artifact name", a.Name).Equals("company-linux-sdk")
}

func Test_LoadAssembleFromAssemblyFile(t *testing.T) {
	a, err := LoadAssembly("test-assembly.yaml")
	assert.That(t, "load err", err).IsNil()
	assert.That(t, "artifact name", a.Name).Equals("company-linux-sdk")
	assert.That(t, "0.repository", a.Parts[0].RepositoryName).Len(0)
	assert.That(t, "1.repository", a.Parts[1].RepositoryName).Equals("other")
}

func Test_LoadAssemblyFromArtifactFile(t *testing.T) {
	a, err := LoadAssembly("test-artifact.yaml")
	assert.That(t, "load err", err).IsNil()
	assert.That(t, "artifact name", a.Name).Equals("README")
}

func Test_LoadArtifact(t *testing.T) {
	a, err := LoadArtifact("test-artifact.yaml")
	assert.That(t, "load err", err).IsNil()
	assert.That(t, "any-os", a.AnyOS).IsTrue()
}

func Test_StorageLocation(t *testing.T) {
	a, err := LoadArtifact("test-artifact.yaml")
	assert.That(t, "load err", err).IsNil()

	loc := a.StorageLocation("Darwin", false)
	assert.That(t, "storage location", loc).Equals("com/company/README/2.0-SNAPSHOT/Darwin/README-2.0-SNAPSHOT.md")

	loc = a.StorageLocation("Darwin", true)
	assert.That(t, "storage any location", loc).Equals("com/company/README/2.0-SNAPSHOT/any-os/README-2.0-SNAPSHOT.md")

}
