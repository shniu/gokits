package bitcask

// BitcaskKVEngine represents the implementation of Bitcask Paper.
type BitcaskKVEngine struct {
}

func (kve *BitcaskKVEngine) Get(key string) []byte {
	return nil
}

func (kve *BitcaskKVEngine) Put(key string, value []byte) error {
	return nil
}

func (kve *BitcaskKVEngine) Delete(key string) error {
	return nil
}

func (kve *BitcaskKVEngine) Merge() error {
	return nil
}

func (kve *BitcaskKVEngine) Close() {
}
