package model

import "time"

type Employee struct {
	Empid     string    `gorm:"primaryKey"`
	Name      string    `json:"name"`
	Dept      string    `json:"dept"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdateAt  time.Time `gorm:"autoCreateTime"`
	DeletedAt time.Time `gorm:"autoCreateTime"`
}

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price float64
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
		Port    string `yaml:"Port"`
		DBName  string `yaml:"DBName"`
		DBFlags string `yaml:"DBFlags"`
		Type    string `yaml:"Type"`
	} `yaml:"database"`
	Jwt struct {
		SecretKey string `yaml:"SecretKey"`
	} `yaml:"jwt"`
	AppAdmin struct {
		Enable bool   `yaml:"enable"`
		User   string `yaml:"enable"`
		Pass   string `yaml:"pass"`
	} `yaml:"appadmin"`
}

type Token struct {
	Raw string `json:"token"`
}
