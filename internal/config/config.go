package config

import "flag"

const NetProtocol = "http"

// const HTTPHost = "127.0.0.1"
// const HTTPPort = ":8080"
const IDLength = 8

var Addr = flag.String("a", "127.0.0.1:8080", "адрес запуска HTTP-сервера")
var Base = flag.String("b", "http://127.0.0.1:8080", "базовый адрес результирующего сокращённого URL")
