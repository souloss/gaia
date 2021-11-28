package bboltkvstorage

import (
	"encoding/json"
	"reflect"

	"github.com/witchc/gaia/interfaces"
	"go.etcd.io/bbolt"
)

var NODE_BUCKET_KEY string = "nodes"

// NodeStorage bbolt repo
type BBoltNodeStorage struct {
	db *bbolt.DB
}

// NewNodeStorage create new repository
func NewNodeStorage(db *bbolt.DB) *BBoltNodeStorage {
	db.Update(func(t *bbolt.Tx) error {
		if _, err := t.CreateBucketIfNotExists([]byte(NODE_BUCKET_KEY)); err != nil {
			return err
		}
		return nil
	})
	return &BBoltNodeStorage{
		db: db,
	}
}

// Put a Node
func (r *BBoltNodeStorage) Put(e interfaces.Node) error {
	err := r.db.Update(func(t *bbolt.Tx) error {
		buf, err := json.Marshal(e)
		if err != nil {
			return err
		}
		b := t.Bucket([]byte(NODE_BUCKET_KEY))
		b.Put([]byte(e.GetHostName()), buf)
		return nil
	})
	return err
}

// Get a Node
func (r *BBoltNodeStorage) Get(hostname string, nodetype reflect.Type) (interface{}, error) {
	rv := reflect.New(nodetype).Interface()
	err := r.db.View(func(t *bbolt.Tx) error {
		b := t.Bucket([]byte(NODE_BUCKET_KEY))
		buf := b.Get([]byte(hostname))
		err := json.Unmarshal(buf, &rv)
		return err
	})
	return rv, err
}

// List a Node
func (r *BBoltNodeStorage) List(nodetype reflect.Type) ([]interface{}, error) {
	result := []interface{}{}
	err := r.db.View(func(t *bbolt.Tx) error {
		b := t.Bucket([]byte(NODE_BUCKET_KEY))
		b.ForEach(func(k, v []byte) error {
			rv := reflect.New(nodetype).Interface()
			if err := json.Unmarshal(v, &rv); err != nil {
				return err
			}
			result = append(result, rv)
			return nil
		})
		return nil
	})
	return result, err
}
