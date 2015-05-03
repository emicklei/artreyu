package model

type CompositeRepository struct {
	repositories []Repository
}

func (r CompositeRepository) Exists(a Artifact) bool {
	for _, each := range r.repositories {
		if each.Exists(a) {
			return true
		}
	}
	return false
}

func (r CompositeRepository) Fetch(a Artifact, destination string) error {
	var err error
	for _, each := range r.repositories {
		if err = each.Fetch(a, destination); err == nil {
			return nil
		}
	}
	return err
}
func (r CompositeRepository) Store(a Artifact, source string) error {
	return nil
}
func (r CompositeRepository) Assemble(a Artifact, source string) error {
	return nil
}
