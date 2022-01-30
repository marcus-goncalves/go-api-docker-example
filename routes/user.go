package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type ResponseBody struct {
	Status  int
	Message string
	Body    interface{}
}

var (
	Data = []User{
		{Id: 1, Name: "Mark", Age: 37},
		{Id: 2, Name: "Rafa", Age: 30},
		{Id: 3, Name: "John", Age: 22},
		{Id: 4, Name: "Jesus", Age: 33},
	}
)

// @Router /user/del/{id} [delete]
// @Tags User
// @Summary Delete an user
// @Description Delete an user by its id
// @Param id path string true "ID User"
// @Success 200 {object} ResponseBody
// @Failure 400 {object} ResponseBody
func (u *User) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 0)
	if err != nil {
		res := ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Invalid Id",
		}

		c.JSON(res.Status, res)
		return
	}

	newData := []User{}
	for _, user := range Data {
		if user.Id != id {
			newData = append(newData, user)
		}
	}
	Data = newData

	res := ResponseBody{
		Status:  http.StatusOK,
		Message: "success",
	}
	c.JSON(res.Status, res)
}

// @Router /user/add [post]
// @Tags User
// @Summary Create a new User
// @Param user body User true "Add user"
// @Success 200 {object} ResponseBody
// @Failure 400 {object} ResponseBody
func (u *User) Add(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		res := ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Invalid body",
		}

		c.JSON(res.Status, res)
		return
	}

	Data = append(Data, newUser)

	res := ResponseBody{
		Status:  http.StatusCreated,
		Message: "success",
		Body:    newUser,
	}
	c.JSON(res.Status, res)
}

// @Router /user/list [get]
// @Tags User
// @Summary Return all users
// @Description Return a list of user objects
// @Success 200 {object} ResponseBody
// @Failure 400 {object} ResponseBody
func (u *User) List(c *gin.Context) {
	if len(Data) == 0 {
		res := ResponseBody{
			Status:  http.StatusNotFound,
			Message: "Query did not return any data",
		}

		c.JSON(res.Status, res)
		return
	}

	res := ResponseBody{
		Status:  http.StatusOK,
		Message: "success",
		Body:    Data,
	}

	c.JSON(res.Status, res)
}

// @Router /user/find [get]
// @Tags User
// @Summary Return a user by id, name or age
// @Description Return a user by its id, name or age as query params
// @Param id query string false "search by id"
// @Param name query string false "search by name"
// @Param age query string false "search by age"
// @Success 200 {object} ResponseBody
// @Failure 400 {object} ResponseBody
func (u *User) Find(c *gin.Context) {
	params := c.Request.URL.Query()

	switch {
	case params.Has("id"):
		id := params.Get("id")
		res := parseId(id)

		c.JSON(res.Status, res)

	case params.Has("name"):
		name := params.Get("name")
		res := parseName(name)

		c.JSON(res.Status, res)

	case params.Has("age"):
		age := params.Get("age")
		res := parseAge(age)

		c.JSON(res.Status, res)

	default:
		res := ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Invalid parameter",
		}
		c.JSON(res.Status, res)
		return
	}

}

func parseId(idParam string) ResponseBody {
	id, err := strconv.ParseInt(idParam, 10, 0)
	if err != nil {
		res := ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Invalid id",
		}
		return res
	}

	for _, user := range Data {
		if id == user.Id {
			res := ResponseBody{
				Status:  http.StatusOK,
				Message: "success",
				Body:    user,
			}
			return res
		}
	}

	res := ResponseBody{
		Status:  http.StatusNotFound,
		Message: "Id not found",
	}

	return res
}

func parseName(name string) ResponseBody {
	users := []User{}
	for _, user := range Data {
		if strings.Contains(strings.ToLower(user.Name), name) {
			fmt.Println(name, user.Name)
			users = append(users, user)
		}
	}

	res := ResponseBody{
		Status:  http.StatusOK,
		Message: "success",
		Body:    users,
	}
	return res
}

func parseAge(ageParam string) ResponseBody {
	age, err := strconv.ParseInt(ageParam, 10, 0)
	if err != nil {
		res := ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Invalid age",
		}

		return res
	}

	users := []User{}
	for _, user := range Data {
		if age == user.Age {
			users = append(users, user)
		}
	}

	res := ResponseBody{
		Status:  http.StatusOK,
		Message: "success",
		Body:    users,
	}
	return res
}
