package database

import (
	"time"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	CreateUserTable(db)
	CreateProjectsTable(db)
	CreateEndpointsTable(db)
}

type User struct {
	Id        uint      `gorm:"unique;primaryKey;autoIncrement"`
	Name      string    `gorm:"size:150;notnull"`
	Email     string    `gorm:"size:100;unique;notnull"`
	Password  string    `gorm:"size:60;notnull"`
	CreatedAt time.Time `gorm:"notnull"`
	UpdatedAt time.Time `gorm:"notnull"`
}

type Project struct {
	Id        string    `gorm:"size:32;unique;primaryKey"`
	UserId    uint      `gorm:"notnull"`
	User      User      `gorm:"foreignKey:UserId;references:Id"`
	Name      string    `gorm:"size:30;notnull"`
	CreatedAt time.Time `gorm:"notnull"`
	UpdatedAt time.Time `gorm:"notnull"`
}

type Endpoint struct {
	Id        uint      `gorm:"primaryKey;autoIncrement"`
	Path      string    `gorm:"size:50;notnull"`
	Method    string    `gorm:"size:10;notnull"`
	Schema    string    `gorm:"size:300;notnull"`
	ProjectId string    `gorm:"size:32;notnull"`
	Project   Project   `gorm:"foreignKey:ProjectId;references:Id"`
	CreatedAt time.Time `gorm:"notnull"`
	UpdatedAt time.Time `gorm:"notnull"`
}

func CreateUserTable(db *gorm.DB) {

	if !Exists(db, &User{}) {
		db.Migrator().CreateTable(&User{})
	}
}

func CreateProjectsTable(db *gorm.DB) {
	if !Exists(db, &Project{}) {
		db.Migrator().CreateTable(&Project{})
	}
}

func CreateEndpointsTable(db *gorm.DB) {
	if !Exists(db, &Endpoint{}) {
		db.Migrator().CreateTable(&Endpoint{})
	}
}
