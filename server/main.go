package main
const (
	port = ":50051"
)

type Config struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}