package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/libsql/go-libsql"
)

func run(err error) {
  dbName := os.Getenv("TURSO_NAME")
  dbUrl := os.Getenv("TRUSO_URL")
  authToken := os.Getenv("TURSO_AUTH")

  dir, err := os.MkdirTemp("", "libsql-*")

  if err != nil {
    fmt.Println("Error creating temp directory:", err)
    os.Exit(1)
  }
  defer os.RemoveAll(dir)

  dbPath := filepath.Join(dir, dbName)

  duration := time.Minute
  connector, err := libsql.NewEmbeddedReplicaConnectorWithAutoSync(dbPath, dbUrl, authToken, duration)

  if err != nil {
    fmt.Println("Error creating connector:", err)
    os.Exit(1)
  }
  defer connector.Close()

  db := sql.OpenDB(connector)
  defer db.Close()
}

