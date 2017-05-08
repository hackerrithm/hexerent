package models

import (
	"fmt"
	"hexerent/backend/database"

	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	UserID       int64
	UserName     string `json:"username"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Password     string `json:"password"`
	EmailAddress string `json:"emailaddress"`
	Country      string `json:"country"`
	City         string `json:"city"`
	ZipCode      string `json:"zipcode"`
	AreaCode     string `json:"areacode"`
	Number       string `json:"number"`
	Extention    string `json:"ext"`
}

// Student struct
type Student struct {
	*User
	School string
	Loan   float32
}

// Employee struct
type Employee struct {
	*User
	Company string
	Money   float32
}

// NewUser : Acts as a constructor
func NewUser(userid int64, username, firstname, lastname, password, emailaddress,
	number, areacode, extention,
	country, city, zipcode string) *User {
	user := new(User)
	user.UserName = username
	user.FirstName = firstname
	user.LastName = lastname
	user.EmailAddress = emailaddress
	user.Password = password
	user.Country = country
	user.City = city
	user.ZipCode = zipcode
	user.AreaCode = areacode
	user.Number = number
	user.Extention = extention

	return user
}

// NewStudent : Acts as a constructor
func NewStudent(usr *User, school string, loan float32) *Student {
	student := new(Student)
	u := usr
	student.School = school
	student.Loan = loan
	student.User = u

	return student
}

// NewEmployee : Acts as a constructor
func NewEmployee(usr *User, company string, money float32) *Employee {
	emp := new(Employee)
	u := usr
	emp.Company = company
	emp.Money = money
	emp.User = u

	return emp
}

// SayHi : Method
func (u User) SayHi() {
	fmt.Printf("Hi, I am %s you can call me Mr. %s\n", u.FirstName, u.LastName)
	fmt.Printf("You can call me at %s%s%s", u.Extention, u.AreaCode, u.Number)
	fmt.Printf(" or find me in %s.", u.Country)
}

// Sing : Method
func (u User) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// SayHi : Method
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.FirstName,
		e.Company, e.Number)
}

// Men implemented by User, Student and Employee
type Men interface {
	SayHi()
	Sing(lyrics string)
}

// InsertNewUser inserts a new user in the sql database
func InsertNewUser(userName, firstName, lastName, password, email string) Student {
	DB, err := database.NewOpen()

	var insertStatement = "INSERT user SET UserName=?,FirstName=?,LastName=?, Password=?, EmailAddress=?"
	stmt, err := DB.Prepare(insertStatement)
	//checkErr(err)

	if err != nil {
		fmt.Println(err)
	}

	hashedPassword, hashingError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if hashingError != nil {
		fmt.Println(hashingError)
	}

	res, err := stmt.Exec(userName, firstName, lastName, hashedPassword, email)

	//checkErr(err)
	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()
	//checkErr(err)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)

	account := NewStudent(NewUser(id, userName, firstName, lastName, password, email, "", "", "", "", "", ""), "", 0.00)

	return *account

}

// FindUser stuff
func FindUser(userName, password string) (Student, bool) {

	user := NewStudent(NewUser(1, userName, "", "", "", "", "", "", "", "", "", ""), "", 0.00)

	foundRecord := false

	DB, err := database.NewOpen()
	/*
		// query
		row := DB.QueryRow("SELECT * FROM user WHERE UserName=? AND Password=?", user.UserName, user.Password)
		err = row.Scan(&user.UserID, &user.UserName, &user.FirstName, &user.LastName, &user.Password, &user.EmailAddress)

		if err != nil {
			fmt.Println(err)
		}

	*/

	//var databaseUsername string
	var databasePassword string
	var id int64

	err = DB.QueryRow("SELECT * FROM user WHERE UserName=?", user.UserName).Scan(&id, &user.UserName, &user.FirstName, &user.LastName, &databasePassword, &user.EmailAddress)

	if err != nil {
		fmt.Println(err)
	}

	hashError := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))

	if hashError != nil {
		fmt.Println(err)
		//http.Redirect(w, r, "/login", http.StatusSeeOther)
		//return
		foundRecord = false
	} else {
		foundRecord = true
	}
	user.UserName = userName
	user.Password = string([]byte(databasePassword))
	user.UserID = id
	fmt.Println(user.UserName, user.Password)
	account := NewStudent(NewUser(user.UserID, user.UserName, "", "", user.Password, "", "", "", "", "", "", ""), "", 0.00)
	fmt.Println("account username: ", account.UserName)
	fmt.Println(account)

	return *account, foundRecord

}
