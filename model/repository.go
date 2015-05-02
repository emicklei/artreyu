package model

type Repository interface {
	Fetch(a Artifact, destination string) error
	Store(a Artifact, source string) error
	Assemble(a Artifact, source string) error
}
