package singleton

// Singleton type
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

// GetInstance return Singleton
func GetInstance() *Singleton {
	return singleton
}
