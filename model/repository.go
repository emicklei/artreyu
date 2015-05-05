package model

import "errors"

var (
	ErrArtifactNotFound = errors.New("artifact not found in repository")
)

type Repository interface {
	ID() string
	Exists(a Artifact) bool
	// destination cannot be a directory
	Fetch(a Artifact, destination string) error
	// source cannot be a directory
	Store(a Artifact, source string) error
}

type EmptyRepository struct{}

func (r EmptyRepository) Exists(a Artifact) bool { return false }
func (r EmptyRepository) Fetch(a Artifact, destination string) error {
	return ErrArtifactNotFound
}
func (r EmptyRepository) Store(a Artifact, source string) error {
	return nil
}
