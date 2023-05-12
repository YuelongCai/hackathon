package rds

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Rds is a wrapper of gorm connection
type Rds struct {
	db *gorm.DB
}

// New Rds
func New(dsn string) (*Rds, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)  // max idle connections
	sqlDB.SetMaxOpenConns(100) // max open connections
	sqlDB.SetConnMaxLifetime(time.Hour)
	return &Rds{db: db}, nil
}

// Begin a transaction, recommend using Transaction
func (d *Rds) Begin() *Rds {
	return &Rds{db: d.db.Begin()}
}

// Commit a transaction, recommend using Transaction
func (d *Rds) Commit() {
	d.db.Commit()
}

// Rollback a transaction, recommend using Transaction
func (d *Rds) Rollback() {
	d.db.Rollback()
}

// Transaction runs a function in a transaction
func (d *Rds) Transaction(f func(tx *gorm.DB) error) error {
	return d.db.Transaction(f)
}
