package main

import (
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

func main() {
  db, err := sql.Open("postgres", "user=postgres password=postgres dbname=maiddesk_development")
  if err != nil {
    fmt.Printf("connection error")
  }

  email := "sample_user@sample.com"
  rows, aerr := db.Query("SELECT id FROM users WHERE email = $1", email)
  
  if aerr != nil {
    fmt.Printf("error code:", aerr)
  } else {
    var company_name string
    for rows.Next() {
      err := rows.Scan(&company_name)
      if err != nil {
        fmt.Printf("another error")
      }
      fmt.Println(company_name)
    }  
  }
}
