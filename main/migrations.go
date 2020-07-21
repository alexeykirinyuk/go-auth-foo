package main

import (
	"github.com/alexeykirinyuk/tech-task-go/components/auth"
	"github.com/alexeykirinyuk/tech-task-go/components/bar"
	"github.com/alexeykirinyuk/tech-task-go/components/foo"
	"github.com/alexeykirinyuk/tech-task-go/components/sigma"
	"github.com/alexeykirinyuk/tech-task-go/data"
)

func runAllMigrations(provider data.IDatabaseProvider) (err error) {
	db, err := provider.CreateConnection()
	if err != nil {
		return
	}

	err = auth.RunMigrations(db)
	if err != nil {
		return
	}

	err = bar.RunMigrations(db)
	if err != nil {
		return
	}

	err = foo.RunMigrations(db)
	if err != nil {
		return
	}

	err = sigma.RunMigrations(db)
	return
}
