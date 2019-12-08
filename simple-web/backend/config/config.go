package config

var DbName string
var DbArg string
var PostPerPage int

func SetDbArg(host, port string) {
	DbArg = `host=${host} port=${port} user=postgres dbname=fmb sslmode=disable password=postgres`
}

func init() {
	DbName = "postgres"
	DbArg = "host=db port=5432 user=postgres dbname=fmb sslmode=disable password=postgres"
	PostPerPage = 5
}