package connect

import (
      "database/sql"
    _"github.com/go-sql-driver/mysql"
  )

func Connect()(connection *sql.DB){
    conexion,err:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/sistema")
    if err!=nil{
      panic(err.Error())
    }
    return conexion
}
