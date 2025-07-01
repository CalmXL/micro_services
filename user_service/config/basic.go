package config

import "flag"

const (
	DBName = "micro_services"
)

var (
	IP   = flag.String("ip", "0.0.0.0", "IP Address")
	PORT = flag.String("port", "50001", "IP Port")
)

const (
	REGPHONENUMBER = `^1[3456789]\d{9}$`
)
