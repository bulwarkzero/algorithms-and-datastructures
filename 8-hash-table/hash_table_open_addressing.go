package hashtable

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

type Prober interface {
	Probe(key string, x int) int
	AdjustCapacity(capacity int) int
}

type opaEntity struct {
	key         string
	value       interface{}
	isTombStone bool
}

type HashTableOpenAddressing struct {
	bucket        []*opaEntity
	capacity      int
	loadFactor    float64
	resizeTrigger int
	size          int // total size, actual_keys + tombstones
	numKeys       int // number of actual keys stored in map
	hash          func(s string) uint64
	prober        Prober
}

func (h *HashTableOpenAddressing) adjustCapacity() {
	h.capacity = h.prober.AdjustCapacity(h.capacity)
}

func (h *HashTableOpenAddressing) probe(key string, x int) int {
	return h.prober.Probe(key, x)
}

func (h *HashTableOpenAddressing) normalizeIndex(idx uint64) uint64 {
	return idx % uint64(h.capacity)
}

func (h *HashTableOpenAddressing) keyIndex(key string, x int) int {
	return int(h.normalizeIndex(h.normalizeIndex(h.hash(key) + uint64(h.probe(key, x)))))
}

func (h *HashTableOpenAddressing) needsResize() bool {
	return h.size >= h.resizeTrigger
}

func (h *HashTableOpenAddressing) resize() {
	oldBucket := make([]*opaEntity, len(h.bucket))
	copy(oldBucket, h.bucket)

	h.capacity *= 2
	h.adjustCapacity()
	h.size = 0
	h.numKeys = 0
	h.resizeTrigger = int(float64(h.capacity) * h.loadFactor)

	h.bucket = make([]*opaEntity, h.capacity)

	for _, entity := range oldBucket {
		if entity == nil || entity.isTombStone {
			continue
		}

		h.Set(entity.key, entity.value)
	}
}

func (h *HashTableOpenAddressing) Set(key string, value interface{}) {
	if h.needsResize() {
		h.resize()
	}

	firstTombstoneIndex := -1

	for x := 0; ; x++ {
		index := h.keyIndex(key, x)
		entity := h.bucket[index]

		if entity != nil && entity.isTombStone {
			// it's a first tombstone index and could be used to shorten find loop
			if firstTombstoneIndex == -1 {
				firstTombstoneIndex = index
			}
		} else if entity != nil {
			// update
			if entity.key == key {
				// we already have a tombstone, so we can swap this key with tombstone
				// and shorten find path
				if firstTombstoneIndex != -1 {
					h.bucket[index] = &opaEntity{isTombStone: true}
					h.bucket[firstTombstoneIndex] = entity
				} else {
					entity.value = value
				}

				return
			}
		} else {
			// new key
			h.numKeys++

			if firstTombstoneIndex != -1 {
				h.bucket[firstTombstoneIndex] = &opaEntity{key: key, value: value}
			} else {
				h.bucket[index] = &opaEntity{key: key, value: value}
				// we used new empty spot, so size must be increased
				h.size++
			}

			break
		}
	}
}

func (h *HashTableOpenAddressing) Get(key string) interface{} {
	firstTombstoneIndex := -1

	for x := 0; ; x++ {
		index := h.keyIndex(key, x)
		entity := h.bucket[index]

		if entity != nil && entity.isTombStone {
			// it's a first tombstone index and could be used to shorten find loop
			if firstTombstoneIndex == -1 {
				firstTombstoneIndex = index
			}
		} else if entity != nil && entity.key == key {
			if firstTombstoneIndex != -1 {
				h.bucket[firstTombstoneIndex] = entity
				h.bucket[index] = &opaEntity{isTombStone: true}
			}

			return entity.value
		} else if entity == nil {
			break
		}
	}

	return nil
}

func (h *HashTableOpenAddressing) Unset(key string) {
	for x := 0; ; x++ {
		index := h.keyIndex(key, x)
		entity := h.bucket[index]

		if entity != nil && entity.isTombStone {
			continue
		} else if entity == nil {
			break
		} else if entity != nil && entity.key == key {
			entity.isTombStone = true
			h.numKeys--

			break
		}
	}
}

func (h *HashTableOpenAddressing) Size() int {
	return h.numKeys
}

func NewOpenAddressing(capacity int, prober Prober) HashTable {
	if capacity < DefaultCapacity {
		capacity = DefaultCapacity
	}

	ht := &HashTableOpenAddressing{
		capacity: capacity,
		hash:     defaultHash,
		prober:   prober,
	}

	ht.adjustCapacity()
	ht.loadFactor = DefaultLoadFactor
	ht.resizeTrigger = int(float64(ht.capacity) * DefaultLoadFactor)
	ht.bucket = make([]*opaEntity, ht.capacity)

	return ht
}
