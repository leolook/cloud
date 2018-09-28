package goid

//extern runtime.getg
func getg() *g

func Get() int64 {
	return getg().goid
}
