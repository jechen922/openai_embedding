package main

import (
	"fmt"
	"log"
	"openaigo/cmd"
	"openaigo/config"
	"openaigo/src/database/mysql"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	cfg := config.New()

	for dbName, dsn := range mysql.DSNMap(cfg) {
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
