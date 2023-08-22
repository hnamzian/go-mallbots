package pg

type PGConfig struct {
	Conn string `mapstructure:"connection"`
}

func (c *PGConfig) Connection() string {
	return c.Conn
}