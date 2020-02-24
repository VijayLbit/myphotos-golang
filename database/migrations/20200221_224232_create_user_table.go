package main

import (
	"github.com/astaxie/beego/migration"
)

// CreateUserTable_20200221_224232 ...
// DO NOT MODIFY
type CreateUserTable_20200221_224232 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUserTable_20200221_224232{}
	m.Created = "20200221_224232"

	migration.Register("CreateUserTable_20200221_224232", m)
}

// Up ...
// Run the migrations
func (m *CreateUserTable_20200221_224232) Up() {
	m.SQL("CREATE TABLE users(id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), email VARCHAR(255), password VARCHAR(100))")
}

// Down ...
// Reverse the migrations
func (m *CreateUserTable_20200221_224232) Down() {
	m.SQL("DROP TABLE users")
}
