package server

type Context map[string]string

func (c *Context) Get(key string) string {
	mp := map[string]string(*c)
	return mp[key]
}

func (c *Context) Put(key string, val string) {
	mp := map[string]string(*c)
	if mp == nil || len(mp) <= 0 {
		mp = make(map[string]string)
	}
	mp[key] = val
	*c = mp
}
