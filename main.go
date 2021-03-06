package main

import (
	"errors"
	"fmt"
	"os"

	// Using postgres sql driver
	_ "github.com/lib/pq"

	"github.com/fengjh/gorm-default-value/callbacks"
	"github.com/jinzhu/gorm"
)

var (
	// DB returns a gorm.DB interface, it is used to access to database
	DB *gorm.DB
)

type Product struct {
	gorm.Model
	Name     string  `sql:"not null"`
	Price    float64 `sql:"not null;default:'0.01'"`
	Category string  `sql:"not null;default:'clothing'"`
}

func (p *Product) BeforeCreate(scope *gorm.Scope) {
	fmt.Printf("BeforeCreate --> %v\n", p.Price)
	fmt.Printf("BeforeCreate --> %v\n", p.Category)
}

func (p *Product) AfterCreate(scope *gorm.Scope) {
	fmt.Printf("AfterCreate --> %v\n", p.Price)
	fmt.Printf("AfterCreate --> %v\n", p.Category)
}

func (p *Product) BeforeCreateTransactionCommit(scope *gorm.Scope) {
	fmt.Printf("BeforeCreateTransactionCommit --> %v\n", p.Price)
	fmt.Printf("BeforeCreateTransactionCommit --> %v\n", p.Category)
}

func (p *Product) AfterCreateTransactionCommit(scope *gorm.Scope) {
	fmt.Printf("AfterCreateTransactionCommit --> %v\n", p.Price)
	fmt.Printf("AfterCreateTransactionCommit --> %v\n", p.Category)
}

func init() {
	initDB()
	migrate()
}

func initDB() {
	var err error
	var db *gorm.DB

	dbParams := os.Getenv("DB_PARAMS")
	if dbParams == "" {
		panic(errors.New("DB_PARAMS environment variable not set"))
	}

	db, err = gorm.Open("postgres", fmt.Sprintf(dbParams))
	if err == nil {
		DB = db
	} else {
		panic(err)
	}

	// Register custom Gorm callbacks
	callbacks.RegisterCallbacks(DB)
}

func migrate() {
	DB.DropTableIfExists(&Product{})
	DB.AutoMigrate(&Product{})
}
