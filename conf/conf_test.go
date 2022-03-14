package conf

import "testing"

func TestConf(t *testing.T) {
	conf := Config()
	t.Log(conf.DB.Addr)
}
