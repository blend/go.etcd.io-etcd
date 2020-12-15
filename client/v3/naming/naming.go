package naming

// NamingOperation defines the corresponding operations for a name resolution change.
type NamingOperation uint8

const (
	// NamingAdd indicates a new address is added.
	NamingAdd NamingOperation = iota
	// NamingDelete indicates an existing address is deleted.
	NamingDelete
)

// NamingUpdate defines a name resolution update. Notice that it is not valid having both
// empty string Addr and nil Metadata in an Update.
type NamingUpdate struct {
	// Op indicates the operation of the update.
	Op NamingOperation
	// Addr is the updated address. It is empty string if there is no address update.
	Addr string
	// Metadata is the updated metadata. It is nil if there is no metadata update.
	// Metadata is not required for a custom naming implementation.
	Metadata interface{}
}

// NamingWatcher watches for the updates on the specified target.
type NamingWatcher interface {
	// Next blocks until an update or error happens. It may return one or more
	// updates. The first call should get the full set of the results. It should
	// return an error if and only if Watcher cannot recover.
	Next() ([]*NamingUpdate, error)
	// Close closes the Watcher.
	Close()
}
