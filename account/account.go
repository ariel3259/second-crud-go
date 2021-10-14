package account

 import (
   "fmt"
   "net/http"
   "encoding/json"
   "golang.org/x/crypto/bcrypt"
   "second-crud-go/connect"
 )
 //User struct for i/o data user
 type User struct{
   Id int
   User string
   Email string
   Password string
 }

  //this function that register a new user
 func Register(w http.ResponseWriter, r *http.Request){
   //user var storages the user's data
    var user User
    //decoding r.Body to
    json.NewDecoder(r.Body).Decode(&user)
    con:=connect.Connect()
    passwordHashed,err:=bcrypt.GenerateFromPassword([]byte(user.Password),14)
    if err!=nil{
      fmt.Fprintf(w,err.Error())
      return
    }
    register,err:=con.Prepare("insert into register(user,email,password) values(?,?,?)")
    if err!=nil{
      fmt.Fprintf(w,err.Error())
    }
    register.Exec(user.User,user.Email,passwordHashed)
    json.NewEncoder(w).Encode("A new user has been registred")
 }

 //a function that auths the user
 func Auth(w http.ResponseWriter,r *http.Request){
   //user var storage Auth  data user
   var user User
   json.NewDecoder(r.Body).Decode(&user)
   con:=connect.Connect()
   auth,err:=con.Query("select * from register where email=?",user.Email)
   if err!=nil{
     json.NewEncoder(w).Encode("Please insert a valid email")
     return
   }


var   userVerify User
//in this for, passwordHashed is captured by userVerify property
   for auth.Next(){
     var u User
     err:=auth.Scan(&u.Id,&u.User,&u.Email,&u.Password)
     if err!=nil{
       panic(err.Error())
     }
     userVerify.Id=u.Id
     userVerify.User=u.User
     userVerify.Email=u.Email
     userVerify.Password=u.Password
   }

//this password variable storages the password introduced by user
   password:=user.Password
//this password variable storages the hash of the password
  passwordHashed:=userVerify.Password

  if err!=nil{
      panic(err.Error())
 }

//the passwordHash is descripted and it is compares to password(for client)
invalidPassword:= bcrypt.CompareHashAndPassword([]byte(passwordHashed),[]byte(password))

//then, the result of invalidPassword validation is checked and his value is storaged at resultOfPasswordHash
resultOfPasswordHash:=invalidPassword==nil
//if resultOfPasswordHash is false, the password is wrong
 if !resultOfPasswordHash{
   json.NewEncoder(w).Encode("The password is wrong")
   return
 }
 //else, the user has been logged
  json.NewEncoder(w).Encode("Congratulations, you are logged")

//fmt.Fprintf(w,"%t",tryingShowResultOfPasswordHash)
 }
