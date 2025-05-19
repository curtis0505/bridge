package multicall

type CallOptions struct {
	chunkSize int
	parallel  bool
}

func NewCallOption() *CallOptions {
	return &CallOptions{
		chunkSize: DefaultBatchSize,
		parallel:  false,
	}
}

func (o CallOptions) GetChunkSize() int { return o.chunkSize }
func (o CallOptions) GetParallel() bool { return o.parallel }

type CallOptionFunc func(*CallOptions)

func WithChunkSize(size int) CallOptionFunc {
	return func(o *CallOptions) {
		o.chunkSize = size
	}
}

func WithCallParallel() CallOptionFunc {
	return func(o *CallOptions) {
		o.parallel = true
	}
}
