package singleton

type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func getInstanct() *Singleton {
	return singleton
}
