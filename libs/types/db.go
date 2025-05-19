package types

type UpdateResult struct {
	MatchedCount  int64 // The number of documents matched by the filter.
	ModifiedCount int64 // The number of documents modified by the operation.
}
