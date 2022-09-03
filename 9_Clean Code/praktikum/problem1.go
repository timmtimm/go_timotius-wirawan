package main

/*
- penamaan struct menggunakan Kapital
*/
type User struct { 
   id       int
   username int
   password int
}

/*
- penamaan struct yang lebih dari satu kata menggunakan PascalCase
- t tidak deskriptif lebih baik ganti customers karena menampung banyak user
*/
type UserService struct { 
   customers []user
}

/*
- u tidak deskriptif lebih baik diganti user
- nama function menggunakan camelCase
*/
func (user UserService) getAllUsers() []user{ 
   return user.customer
}

/*
- u tidak deskriptif lebih baik diganti user
- nama function menggunakan camelCase
- r diganti dengan customer karena mewakili satu user
*/
func (user UserService) getUserByID(id int) user{ 
   for _, customer := range user.customers {
       if id == customer.id {
           return customer
       }
   }
   return user{}
}