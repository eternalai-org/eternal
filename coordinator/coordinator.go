package coordinator

type Coordinator struct {
	Workers map[string]*Worker
}

type Worker struct {
	WorkerID string
}

func NewCoordinator() *Coordinator {
	return &Coordinator{}
}
