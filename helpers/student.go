package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const USERNAME, PASSWORD = "batman", "password"

type Student struct {
	ID    string
	Name  string
	Grade int32
}

var students []Student

func ActionStudent(ctx *gin.Context) {
	ctx.Request.BasicAuth()
	if !Auth(ctx) {
		return
	}
	if !AllowOnlyGET(ctx) {
		return
	}
	if id := ctx.Param("id"); id != "" {
		student := GetStudentByID(id)
		ctx.JSON(http.StatusOK, gin.H{
			"student": student,
		})
		return
	}
	allStudents := GetAllStudent()
	ctx.JSON(http.StatusOK, gin.H{
		"allStudents": allStudents,
	})

}

func Auth(ctx *gin.Context) bool {
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "something went wrong",
		})
		return false
	}
	if isValid := (username == USERNAME) && (password == PASSWORD); !isValid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "wrong username/password",
		})
		return false
	}
	return true
}

func AllowOnlyGET(ctx *gin.Context) bool {
	if ctx.Request.Method != "GET" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "only allow GET method",
		})
		return false
	}
	return true
}

func init() {
	// append(students, Student{ID: 1, Name: "Rosa", Grade: 1})
	// append(students, Student{ID: 2, Name: "Maria", Grade: 2})
}

func GetAllStudent() []Student {
	return students
}

func GetStudentByID(id string) Student {
	for _, student := range students {
		if student.ID == id {
			return student
		}
	}
	return Student{}
}
