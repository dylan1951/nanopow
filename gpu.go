//go:build !cl
// +build !cl

package nanopow

type noneGPUWorker struct{}

func NewWorkerGPU(_ interface{}) (*noneGPUWorker, error) {
	return NewWorkerGPUThread(0, nil)
}

func NewWorkerGPUThread(_ uint64, _ interface{}) (*noneGPUWorker, error) {
	return nil, ErrNotSupported
}

func (w *noneGPUWorker) GenerateWork(ctx *Context, root []byte, difficulty uint64) (err error) {
	return ErrNotSupported
}

func GetDevice() (interface{}, error) {
	return nil, ErrNotSupported
}
