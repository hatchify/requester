package requester

// KVMap is a helper map for key/val pair objects
type KVMap map[string]string

func (k KVMap) add(entries ...KVInterface) {
	for _, entry := range entries {
		k[entry.getKey()] = entry.getVal()
	}
}

func (k KVMap) forEach(fn func(key, val string) error) (err error) {
	for key, val := range k {
		if err = fn(key, val); err != nil {
			return
		}
	}

	return
}

// KV is a helper struct for key/val pairs
type KV struct {
	Key string
	Val string
}

func (kv *KV) getKey() (key string) {
	return kv.Key
}

func (kv *KV) getVal() (key string) {
	return kv.Val
}

// KVInterface is the interface wrapper of KV
type KVInterface interface {
	getKey() (key string)
	getVal() (val string)
}

// KVS is the slice of KVS
type KVS []KV
