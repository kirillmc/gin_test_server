package main

import (
	"log"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kirillmc/data_filler/pkg/filler"
	"github.com/kirillmc/gin_test_server/internal"
	"github.com/kirillmc/gin_test_server/internal/model"
)

func main() {

	r := gin.Default()

	r.GET("/programs/:n", getNPrograms)
	r.POST("/programs/", createPrograms)
	r.PUT("/programs/", updatePrograms)
	r.DELETE("/programs/:n", deleteProgramByID)

	//r.GET("/programs", getPrograms)
	log.Println("Server is running on: localhost:8081...")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

func getNPrograms(c *gin.Context) {
	n, err := getId(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid programs count"})
		return
	}

	programs := filler.CreateOwnSetOfPrograms(int(n))

	c.JSON(http.StatusOK, programs)
}

func createPrograms(c *gin.Context) {
	var programs internal.TrainPrograms
	if err := c.ShouldBindJSON(&programs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := model.Response{Message: "Данные были добавлены"}

	c.JSON(http.StatusCreated, message)
}

func updatePrograms(c *gin.Context) {
	var programs internal.TrainPrograms
	if err := c.ShouldBindJSON(&programs); err != nil {
		log.Println("1")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := model.Response{Message: "Данные были обновлены"}

	c.JSON(http.StatusOK, message)
}

func deleteProgramByID(c *gin.Context) {
	_, err := getId(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid programs ID"})
		return
	}

	message := model.Response{Message: "Данные были удалены"}

	c.JSON(http.StatusOK, message)
}

func getId(c *gin.Context) (int64, error) {
	idStr := c.Param("n")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

//
//func getPrograms(c *gin.Context) {
//	query := fmt.Sprintf("SELECT %s ,%s, %s, %s, %s, %s, %s, %s, %s, %s, %s FROM users", idColumn, name, surname, email, avatar, login, password, role, weight, height, locked)
//	rows, err := db.Query(query)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
//		return
//	}
//	defer rows.Close()
//
//	users := []UserToGet{}
//	for rows.Next() {
//		var user UserToGet
//		err := rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.Avatar, &user.Login, &user.Password,
//			&user.Role, &user.Weight, &user.Height, &user.Locked)
//		if err != nil {
//			log.Println(err)
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
//			return
//		}
//		users = append(users, user)
//	}
//
//	c.JSON(http.StatusOK, users)
//}
