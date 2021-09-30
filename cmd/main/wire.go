package main

import (
	"github.com/google/wire"
	"hxj/initials"
)

func NewMigrationJob() error {
	wire.Build(initials.NewDB())
	return nil
}
