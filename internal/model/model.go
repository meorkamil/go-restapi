package model

import "time"

type Employee struct {
	Id        uint      `gorm:"primaryKey"`
	Empid     string    `json:"empid"`
	Name      string    `json:"name"`
	Dept      string    `json:"dept"`
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
}
