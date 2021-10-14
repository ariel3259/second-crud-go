package Employe

import (
  "second-crud-go/connect"
  "net/http"
  "fmt"
  "encoding/json"
)
type Employee struct{
  Id int
  Nombre string
  Correo string
}



func Home(w http.ResponseWriter,r *http.Request){
  fmt.Fprintf(w,"Hi word")
}

//set employes
func GetEmployees(w http.ResponseWriter,r *http.Request){
  //the conditional verifies that Method is Get, if it is Get, continue that function, else stops the function
  if r.Method!="GET"{
    fmt.Fprintf(w,"Method not suported")
    return
  }
  //con var it is created and initialized
  con:=connect.Connect()
  //prepare the sql sentences, this may be happend an error
  registro,err:=con.Query("Select * from empleados")
  //if there is an error, stops the function
  if err!=nil{
    panic(err.Error())
  }
  //else the data are saved at registro and continues
  //empleado var is a Employee object
  empleado:=Employee{}
  //arregloEmpleado var is an array of Employee object
  arregloEmpleado:=[]Employee{}
  //registro is route
  for registro.Next(){
//three variables gonna be create
//the first saves id into id variable
//the second saves the name into nombre variable
//the thirsd saves the address into correo variable
    var id int
    var nombre, correo string
    //verify if there is an error while registro is scaning
    err:=registro.Scan(&id,&nombre,&correo)
    //if there are an error, the program stops an show us an console alert
    if err!=nil{
      panic(err.Error())
    }
    //else, those variables are stored into  empleado object
    empleado.Id=id
    empleado.Nombre=nombre
    empleado.Correo=correo
    //after stored all the variables into the object
    //empleado is stored into  arregloEmpleado
    arregloEmpleado=append(arregloEmpleado,empleado)
  }
  //creates the header content-type
    w.Header().Set("Content-Type","application/json")
  //then writes an header status
    w.WriteHeader(http.StatusCreated)
    //Employees array is encode to json and send it to client
    json.NewEncoder(w).Encode(arregloEmpleado)
}

//create emoployes
func SetEmployees(w http.ResponseWriter, r *http.Request){
  if r.Method!="POST"{
  fmt.Fprintf(w,"Error, method not supported")
  return
  }
//this var gonna save the employee data
  var employee Employee

  //decode json data to employee data
  json.NewDecoder(r.Body).Decode(&employee)
//con var  it is created and initialized
  con:=connect.Connect()
  //prepare the sql sentences, this may be happen an error
  prepareInsert,err:=con.Prepare("insert into empleados(nombre,correo) values(?,?)")
//if there is an error, the program show us a panic message
  if err!=nil{
    panic(err.Error())
    return
  }
  //if there isn't an error, the function continues to execute the sql sentences
  prepareInsert.Exec(employee.Nombre,employee.Correo)
  //then writes an header status
  w.WriteHeader(http.StatusCreated)
  //the message is encoded by json and send it to the client
  json.NewEncoder(w).Encode("A new employee has been hired")
}

//update employe
func UpdateEmployees(w http.ResponseWriter, r *http.Request){
//the url must be http://localhost:3000/api/deleteempleados?id={id}
if r.Method!="PUT"{
  fmt.Fprintf(w,"Error, method not supported")
  return
}
var employee Employee
json.NewDecoder(r.Body).Decode(&employee)
con:=connect.Connect()
updateInsert,err:=con.Prepare("update empleados set nombre=?,correo=? where id=?")
if err!=nil{
  fmt.Fprintf(w,err.Error())
  return
}
updateInsert.Exec(employee.Nombre,employee.Correo,employee.Id)
json.NewEncoder(w).Encode("A employee has been updated")
}

//delete employe
func DeleteEmployee(w http.ResponseWriter,r *http.Request){
  if r.Method !="DELETE"{
    fmt.Fprintf(w,"Error, method not supported")
  }
  idEmpleado:=r.URL.Query().Get("id")
  con:=connect.Connect()
  deleteInsert,err:=con.Prepare("delete from empleados where id=?")
  if err!=nil{
    fmt.Fprintf(w,err.Error())
  }
  deleteInsert.Exec(idEmpleado)
  json.NewEncoder(w).Encode("A employee has been deleted  ")
}
