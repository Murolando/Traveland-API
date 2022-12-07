package repository
type Config struct{
	host string
	port int
	user string
	password string
	dbname string
}

func NewConfig(host string,port int, user string,password string, dbname string)*Config{
	return &Config{
		host: host,
		port: port,
		user: user,
		password: password,
		dbname: dbname,
	}
}

type PostgresDB struct{
	cfg Config
}

func (p *PostgresDB) 
