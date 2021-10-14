//API RESTFUL IN GO WITH MYSQL
package main

import (
  "net/http"
  "fmt"
  "second-crud-go/Employe"
  "second-crud-go/account"
)

func main(){
  fmt.Printf("Server listening at port 3000 \n")
  http.HandleFunc("/",Employe.Home)
  http.HandleFunc("/api/getempleados",Employe.GetEmployees)
  http.HandleFunc("/api/setempleados",Employe.SetEmployees)
  http.HandleFunc("/api/updateempleados",Employe.UpdateEmployees)
  http.HandleFunc("/api/deleteempleados",Employe.DeleteEmployee)
  http.HandleFunc("/api/register",account.Register)
  http.HandleFunc("/api/auth",account.Auth)
  http.ListenAndServe(":3000",nil)
}
