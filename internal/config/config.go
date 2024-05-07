package config

import (
	"flag"
	"os"
)

var initialized = false

func Initialized() bool {
	return initialized
}

var Config = struct {
	Addr     string
	Base     string
	IDLength int
}{
	// не передаваемые параметры
	IDLength: 8,
}

func Init() {
	getValues(
		[]value{
			{
				&Config.Addr,
				"SERVER_ADDRESS",
				"a",
				"127.0.0.1:8080",
				"адрес запуска HTTP-сервера",
			},
			{
				&Config.Base,
				"BASE_URL",
				"b",
				"http://127.0.0.1:8080",
				"базовый адрес результирующего сокращённого URL",
			},
		},
	)
	initialized = true
}

type value struct {
	p            *string
	envName      string
	flagName     string
	defaultValue string
	usage        string
}

func getValues(values []value) {
	for _, v := range values {
		var ok bool
		if *v.p, ok = os.LookupEnv(v.envName); !ok {
			flag.StringVar(v.p, v.flagName, v.defaultValue, v.usage)
		}
	}
	flag.Parse()
}
