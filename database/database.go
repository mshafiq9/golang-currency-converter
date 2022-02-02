package DatabaseManager

import (
    "database/sql"
    "os"

    "github.com/lib/pq"
)


func OpenDatabase() (db *sql.DB, err error) {
    url := os.Getenv("DATABASE_URL")
    connection, _ := pq.ParseURL(url)
    connection += " sslmode=require"

    db, err = sql.Open("postgres", connection)

    if err != nil {
      panic(err)
    }

    return db, err
}

func CloseDatabase(db *sql.DB){
    defer db.Close()
}
