package kv

// FamilyOption defines config items for family level
type FamilyOption struct {
	ID   int    `toml:"id"`
	Name string `toml:"name"`
}

// StoreOption defines config item for store level
type StoreOption struct {
	Path   string `toml:"-"` // ignore path field for INFO file
	Levels int    `toml:"levels"`
}

// DefaultStoreOption builds default store option
func DefaultStoreOption(path string) StoreOption {
	return StoreOption{
		Path:   path,
		Levels: 2,
	}
}

// storeInfo stores store config option, include all family's option in this kv store
type storeInfo struct {
	StoreOption StoreOption             `toml:"store"`
	Families    map[string]FamilyOption `toml:"families"`
}

// newStoreInfo creates store info instance for saving configs
func newStoreInfo(storeOption StoreOption) *storeInfo {
	return &storeInfo{
		StoreOption: storeOption,
		Families:    make(map[string]FamilyOption),
	}
}
