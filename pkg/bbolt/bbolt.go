package bbolt

import (
	"fmt"
	"io/fs"

	bolt "go.etcd.io/bbolt"
)

// singleton storing a connection to a BoltDB file.
// BoltDB is a key/value, binary file, database.
// It's pretty easy-to-use DB, fast, concurrency safe
// and can be used as a localhost DB.
// Basically a map[string]string (or map[string]map[string]string) stored as a binary file.
var db *bolt.DB

// Open opens a boltDB connection and stores it into a singleton.
func Open(path string, mode fs.FileMode, options *bolt.Options) error {
	var err error
	db, err = bolt.Open(path, mode, options)
	return err
}

// Close closes. :)
func Close() {
	db.Close()
}

// Write allows to write data (mostly marshaled) linked to a key
// inside a bolt bucket (equivalent of a SQL's table or NoSQL's collection)
func Write(bucket, key, content string) error {
	return db.Update(func(tx *bolt.Tx) error {
		// fetch or create a new bucket
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		// "insert"
		b.Put([]byte(key), []byte(content))

		return nil
	})
}

// Fetch retrieves content linked to a key from a bucket.
func Fetch(bucket, key string) ([]byte, error) {
	var content []byte

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s does not exist", bucket)
		}
		// actually retrieving content
		content = b.Get([]byte(key))
		return nil
	})

	return content, err
}

func Iterate(bucket string, action func(key, value []byte)) error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		if b == nil {
			return fmt.Errorf("bucket %s does not exist", bucket)
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			action(k, v)
		}

		return nil
	})
}
