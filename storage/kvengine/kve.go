package kvengine

// KVE is the top interface of key-value engine, also-is key value engine.
type KVE interface {
	Get(key string) []byte
	Put(key string, value []byte) error
	Delete(key string) error
	Merge() error

	Close()
}

func Open() (KVE, error) {
	return nil, nil
}
