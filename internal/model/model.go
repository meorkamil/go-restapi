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
		Addr string `yaml:"addr"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host    string `yaml:"host"`
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		Port    string `yaml:"port"`
		DBName  string `yaml:"dbname"`
		DBFlags string `yaml:"dbflags"`
		Type    string `yaml:"type"`
	} `yaml:"database"`
	Jwt struct {
		SecretKey string `yaml:"secretkey"`
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
