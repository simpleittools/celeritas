package main

import (
	"fmt"
	"github.com/fatih/color"
	"path/filepath"
	"time"
)

func doAuth() error {
	rootpath := filepath.ToSlash(cel.RootPath)
	// migrations
	dbType := cel.DB.DataType

	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())

	upFile := rootpath + "/migrations/" + fileName + ".up.sql"
	downFile := rootpath + "/migrations/" + fileName + ".down.sql"

	err := copyFilefromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens cascade;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	// copy files over
	err = copyFilefromTemplate("templates/data/user.go.txt", rootpath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/data/token.go.txt", rootpath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/data/remember_token.go.txt", rootpath+"/data/remember_token.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy over middleware
	err = copyFilefromTemplate("templates/middleware/auth.go.txt", rootpath+"/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/middleware/auth-token.go.txt", rootpath+"/middleware/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/middleware/remember.go.txt", rootpath+"/middleware/remember.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy over handlers
	err = copyFilefromTemplate("templates/handlers/auth-handlers.go.txt", rootpath+"/handlers/auth-handlers.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy over views
	err = copyFilefromTemplate("templates/mailer/password-reset.html.tmpl", rootpath+"/mail/password-reset.html.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/mailer/password-reset.plain.tmpl", rootpath+"/mail/password-reset.plain.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/views/login.jet", rootpath+"/views/login.jet")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/views/forgot.jet", rootpath+"/views/forgot.jet")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/views/reset-password.jet", rootpath+"/views/reset-password.jet")
	if err != nil {
		exitGracefully(err)
	}

	color.Yellow("  - users, tokens, and remember_tokens migrations crated and executed")
	color.Yellow("  - user and token models created")
	color.Yellow("  - auth middleware created")
	color.Yellow("")
	color.Yellow("Don't forget to add user and token models in data/models.go, and add appropriate middleware to your routes!")

	return nil
}
