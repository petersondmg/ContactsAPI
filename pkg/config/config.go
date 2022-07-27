package config

type Config struct {
	API struct {
		Addr      string
		JWTSecret string
	}

	VarejaoDB struct {
		User     string
		Pass     string
		Database string
		Addr     string
	}

	MacapaDB struct {
		User     string
		Pass     string
		Database string
		Addr     string
	}
}

func Load() (*Config, error) {
	// parse config from files / envs
	c := new(Config)

	c.API.Addr = ":8082"
	c.API.JWTSecret = "jwt_secret_key"

	c.VarejaoDB.Database = "varejao"
	c.VarejaoDB.User = "admin"
	c.VarejaoDB.Pass = "admin"
	c.VarejaoDB.Addr = "postgres:5432"

	c.MacapaDB.Database = "macapa"
	c.MacapaDB.User = "admin"
	c.MacapaDB.Pass = "admin"
	c.MacapaDB.Addr = "mysql:3306"

	return c, nil
}
