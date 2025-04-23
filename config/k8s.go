//go:build k8s

package config

var Config = &Config{
	DB: DBConfig{
		DNS: "root:root@tcp(localhost:30002)/webook",
	},
	Redis: RedisConfig{
		Addr: "localhost:30003",
	},
}
