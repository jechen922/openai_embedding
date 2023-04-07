package main

import (
	"fmt"
	"log"
	"openai_golang/cmd"
	"openai_golang/config"
	"openai_golang/src/database/mysql"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	if err := config.EnvInit(); err != nil {
		log.Fatal(err)
	}

	for dbName, dsn := range mysql.DSNMap() {
		app := &cli.App{
			Description: "manage database migrations",
			ArgsUsage:   "",
			Action: func(c *cli.Context) error {
				tables := mysql.Tables(dbName)
				if len(tables) == 0 {
					return fmt.Errorf("DB %s not sepcify any table", dbName)
				}
				dbName = strings.ToLower(dbName)
				_, filename, _, _ := runtime.Caller(0)
				basePath := filepath.Join(filepath.Dir(filename), "../..")
				outPath := path.Join(basePath, "/src/model/po/"+dbName)
				if err := cmd.Exec("gentool",
					"-db", "mysql",
					"-dsn", dsn,
					"-modelPkgName", dbName,
					"-outPath", outPath,
					"-tables", strings.Join(tables, ","),
					"-fieldNullable",
					"-fieldWithTypeTag",
					"-fieldWithIndexTag",
					"-fieldSignable",
					"-onlyModel",
				); err != nil {
					log.Fatal(err)
				}
				return nil
			},
		}
		if err := app.Run(os.Args); err != nil {
			fmt.Println(err.Error())
		}
	}
}
