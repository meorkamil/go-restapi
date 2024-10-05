package model

type Employee struct {
	Empid string `json:"empid"`
	Name  string `json:"name"`
	Dept  string `json:"dept"`
}

type Config struct {
	Server struct {
		Addr string `yaml:"Addr"`
		Port int    `yaml:"Port"`
	} `yaml:"server"`
	Database struct {
		Host    string `yaml:"Host"`
		User    string `yaml:"User"`
		Pass    string `yaml:"Pass"`
		DBName  string `yaml:"DBName"`
		DBFlags string `yaml:"DBFlags"`
	} `yaml:"database"`
}
