package model

import (
	"fmt"
	"log"
)

type CachingRepository struct {
	source Repository
	cache  Repository
}

func NewCachingRepository(main, cache Repository) Repository {
	return CachingRepository{main, cache}
}

func (r CachingRepository) ID() string {
	return fmt.Sprintf("%s with %s as cache", r.source.ID(), r.cache.ID())
}

func (c CachingRepository) Exists(a Artifact) bool {
	if a.IsSnapshot() {
		// TODO fetch timestamp of latest and compare with local
		return false
	}
	return c.cache.Exists(a)
}

func (c CachingRepository) Fetch(a Artifact, destination string) error {
	if a.IsSnapshot() {
		return c.source.Fetch(a, destination)
	}
	err := c.cache.Fetch(a, destination)
	if err != nil {
		log.Printf("[%s] not found on [%s], try fetching from [%s]\n", a.StorageBase(), c.cache.ID(), c.source.ID())
		return c.source.Fetch(a, destination)
	}
	return nil
}

func (c CachingRepository) Store(a Artifact, source string) error {
	log.Printf("storing [%s] in [%s]\n", a.StorageBase(), c.source.ID())
	if err := c.source.Store(a, source); err != nil {
		return err
	}
	if a.IsSnapshot() {
		return nil
	}
	log.Printf("storing [%s] in [%s]\n", a.StorageBase(), c.cache.ID())
	return c.cache.Store(a, source)
}
