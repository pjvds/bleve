package foundationdb

import (
	"encoding/json"
	"errors"

	"github.com/blevesearch/bleve/index/store"
	"github.com/blevesearch/bleve/registry"
)

const (
	Name                    = "foundationdb"
	defaultCompactBatchSize = 100
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

type Store struct {
	db            fdb.Database
	space         subspace.Subspace
	readonly      bool
	mergeOperator store.MergeOperator
}

func New(mo store.MergeOperator, config map[string]interface{}) (store.KVStore, error) {
	// Open the default database from the system cluster
	db := fdb.MustOpenDefault()
	space := subspace.Sub("blave")

	readonly, _ := config["read_only"].(bool)

	return &Store{
		db:            db,
		space:         space,
		readonly:      readonly,
		mergeOperator: mo,
	}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (bs *Store) Reader() (store.KVReader, error) {
	return nil, ErrNotImplemented
}

func (bs *Store) Writer() (store.KVWriter, error) {
	return nil, ErrNotImplemented
}

func (bs *Store) Stats() json.Marshaler {
	return nil
}

func init() {
	registry.RegisterKVStore(Name, New)
}
