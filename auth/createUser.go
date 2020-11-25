package auth

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	db "nusbi-server/config"
	"strings"
)

/*
Error message number:
-1: Success
1: Unable to parse body
2: Username should be at least 5 characters
3: Password should be at least 8 characters
4: Role should be a,t or s
*/

type createUserRequest struct {
	Username string
	Password string
	Role     string
}

type createUserResponse struct {
	Error int
}

func CreateUser(c *fiber.Ctx) error {
	var err error
	// TODO: Allow the addition user data
	request := createUserRequest{}
	err = c.BodyParser(&request)
	if err != nil {
		return c.JSON(createUserResponse{Error: 1})
	}

	// Data validation
	if len(request.Username) < 5 {
		return c.JSON(createUserResponse{
			Error: 2,
		})
	}
	if len(request.Password) < 8 {
		return c.JSON(createUserResponse{Error: 3})
	}
	if !(request.Role == "a" || request.Role == "t" || request.Role == "s") {
		return c.JSON(createUserResponse{Error: 4})
	}

	// Add user to config
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 7)
	if err != nil {
		return c.SendStatus(500)
	}
	_, err = db.Db.Exec(
		"insert into Users (user_id, password, role) VALUES (?,?,?)",
		strings.ToLower(request.Username),
		string(hash),
		request.Role,
	)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(createUserResponse{Error: -1})
}

/*
-1: Success
1: Invalid token role
2: Invalid request body
3: Username short
4: Password short
5: Duplicate user
*/

type createAdminRequest struct {
	Username string
	Password string
}

type createAdminResponse struct {
	Error int
}

// CreateAdmin create a new admin account
func CreateAdmin(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.JSON(createAdminResponse{Error: 1})
	}
	request := createAdminRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.JSON(createAdminResponse{Error: 2})
	}
	if len(request.Username) < 5 {
		return c.JSON(createAdminResponse{Error: 3})
	}
	if len(request.Password) < 8 {
		return c.JSON(createAdminResponse{Error: 4})
	}
	// Add user to config
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 7)
	if err != nil {
		return c.SendStatus(500)
	}
	_, err = db.Db.Exec(
		"insert into nusbiam.Users (user_id, password, role) VALUES (?,?,?)",
		strings.ToLower(request.Username),
		string(hash),
		"a",
	)
	if err != nil {
		if strings.Contains(err.Error(),"Error 1062"){
			return c.JSON(createUserResponse{Error: 5})
		}
		return c.SendStatus(500)
	}
	return c.JSON(createUserResponse{Error: -1})
}

type createStudentRequest struct {
	Username  string
	Password  string
	Batch     int
	FirstName string
	LastName  string
	Gender    string
	Dob       string
	Email     string
	Major string
}



/*
-1: Success
1: Invalid token role
2: Body parsing error
3: Short username
4: Short possword
5: Duplicate user
6: Invalid major
*/

func CreateStudent(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.JSON(createAdminResponse{Error: 1})
	}
	request := createStudentRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(createAdminResponse{Error: 2})
	}
	if len(request.Username) < 5 {
		return c.JSON(createAdminResponse{Error: 3})
	}
	if len(request.Password) < 8 {
		return c.JSON(createAdminResponse{Error: 4})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 7)
	if err != nil {
		return c.SendStatus(500)
	}
	_, err = db.Db.Exec(
		"insert into Users (user_id, password, role) VALUES (?,?,?)",
		strings.ToLower(request.Username),
		string(hash),
		"s",
	)
	if err != nil {
		if strings.Contains(err.Error(),"Error 1062"){
			return c.JSON(createAdminResponse{Error: 5})
		}
		return c.SendStatus(500)
	}
	_, err = db.Db.Exec("insert into Students (student_id,first_name,last_name,gender,dob,email,user_id,major_id,batch)"+
		" value (?,?,?,?,?,?,?,?,?)",
		uuid.New().String(), request.FirstName, request.LastName, request.Gender, request.Dob, request.Email, request.Username, request.Major, request.Batch)
	if err!=nil{
		if strings.Contains(err.Error(),"Error 1452"){
			return c.JSON(createAdminResponse{Error: 6})
		}
		return c.SendStatus(500)
	}
	return c.JSON(createUserResponse{Error: -1})
}

type createTeacherRequest struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
	Gender    string
	Dob       string
	Email     string
}

/*
-1: Success
1: Invalid token role
2: Body parsing error
3: Short username
4: Short possword
5: Duplicate user
*/

func CreateTeacher(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.JSON(createAdminResponse{Error: 1})
	}
	request := createTeacherRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(createAdminResponse{Error: 2})
	}
	if len(request.Username) < 5 {
		return c.JSON(createAdminResponse{Error: 3})
	}
	if len(request.Password) < 8 {
		return c.JSON(createAdminResponse{Error: 4})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 7)
	if err != nil {
		return c.SendStatus(500)
	}
	_, err = db.Db.Exec(
		"insert into Users (user_id, password, role) VALUES (?,?,?)",
		strings.ToLower(request.Username),
		string(hash),
		"t",
	)
	if err != nil {
		if strings.Contains(err.Error(),"Error 1062"){
			return c.JSON(createAdminResponse{Error: 5})
		}
		return c.SendStatus(500)
	}
	_, err = db.Db.Exec("insert into nusbiam.Lecturers (lecturer_id,first_name,last_name,gender,dob,email,user_id)"+
		" value (?,?,?,?,?,?,?)",
		uuid.New().String(), request.FirstName, request.LastName, request.Gender, request.Dob, request.Email, request.Username)
	if err!=nil{
		return c.SendStatus(500)
	}
	return c.JSON(createUserResponse{Error: -1})
}