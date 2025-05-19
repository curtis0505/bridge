package cwutil

// Expiration represents a point in time when some event happens.
// It can compare with a BlockInfo and will return is_expired() == true
// once the condition is hit (and for every block in the future)
// https://github.com/CosmWasm/cw-utils/blob/b30db1e1bf230f71f0d1f77757439e2d74512eb9/src/expiration.rs#L12
type Expiration struct {
	AtHeight *int64 `json:"at_height,omitempty"`
	AtTime   *int64 `json:"at_time,omitempty"`

	// Never will never expire. Used to express the empty variant
	Never *Never `json:"never,omitempty"`
}

type Never struct {
}

func NewExpiration() Expiration {
	return Expiration{
		Never: &Never{},
	}
}

func NewExpirationAtHeight(height int64) Expiration {
	h := height
	return Expiration{
		AtHeight: &h,
	}
}

func NewExpirationAtTime(time int64) Expiration {
	t := time
	return Expiration{
		AtTime: &t,
	}
}
