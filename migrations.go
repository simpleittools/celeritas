package celeritas

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"path/filepath"
	_ "path/filepath"
)

func (c *Celeritas) MigrateUp(dsn string) error {
	rootpath := filepath.ToSlash(c.RootPath)
	m, err := migrate.New("file://"+rootpath+"/migrations", dsn)

	if err != nil {
		return err
	}

	defer m.Close()

	if err := m.Up(); err != nil {
		log.Println("Error running migration", err)
		return err
	}

	return nil

}

func (c *Celeritas) MigrateDownAll(dsn string) error {
	rootpath := filepath.ToSlash(c.RootPath)
	m, err := migrate.New("file://"+rootpath+"/migrations", dsn)

	if err != nil {
		return err
	}

	defer m.Close()

	if err := m.Down(); err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) Steps(n int, dsn string) error {
	rootpath := filepath.ToSlash(c.RootPath)
	m, err := migrate.New("file://"+rootpath+"/migrations", dsn)

	if err != nil {
		return err
	}

	defer m.Close()

	if err := m.Steps(n); err != nil {
		return err
	}

	return nil
}

func (c *Celeritas) MigrateForce(dsn string) error {
	rootpath := filepath.ToSlash(c.RootPath)
	m, err := migrate.New("file://"+rootpath+"/migrations", dsn)

	if err != nil {
		return err
	}

	defer m.Close()

	if err := m.Force(-1); err != nil {
		return err
	}

	return nil
}
