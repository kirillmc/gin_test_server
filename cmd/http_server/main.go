package main

import (
	"log"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kirillmc/data_filler/pkg/filler"
)

func main() {

	r := gin.Default()

	r.GET("/programs/:n", getNPrograms)

	//r.POST("/programs", createProgram)
	//r.GET("/programs", getPrograms)
	//r.PUT("/programs/:n", updateProgramById)
	//r.DELETE("/programs/:n", deleteProgramByID)

	log.Println("Server is running on: localhost:80808...")
	if err := r.Run(":8080"); err != nil {
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

func getId(c *gin.Context) (int64, error) {
	idStr := c.Param("n")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

//func createProgram(c *gin.Context) {
//	//var user User
//	//if err := c.ShouldBindJSON(&user); err != nil {
//	//	log.Println("1")
//	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//	//	return
//	//}
//	//log.Println(user)
//	//var id int64
//	//
//	//if err != nil {
//	//	log.Println("2")
//	//	log.Println(err)
//	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
//	//	return
//	//}
//	//
//	//ID := CreateResponse{
//	//	Id: id,
//	//}
//	//
//	//c.JSON(http.StatusCreated, ID)
//}
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
//
//
//
//func updateProgramById(c *gin.Context) {
//	id, err := getId(c)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
//		return
//	}
//
//	var updateUser User
//	if err := c.ShouldBindJSON(&updateUser); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	query := fmt.Sprintf("UPDATE users SET %s=$1, %s=$2, %s=$3, %s=$4, %s=$5, %s=$6, %s=$7, %s=$8, %s=$9, %s=$10 WHERE id=$11", name, surname, email, avatar, login, password, role, weight, height, locked)
//	_, err = db.Exec(query,
//		updateUser.Name, updateUser.Surname, updateUser.Email, updateUser.Avatar, updateUser.Login, updateUser.Password,
//		updateUser.Role, updateUser.Weight, updateUser.Height, updateUser.Locked, id)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
//		return
//	}
//
//	c.Status(http.StatusOK)
//}
//
//func deleteProgramByID(c *gin.Context) {
//	id, err := getId(c)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
//		return
//	}
//
//	result, err := db.Exec("DELETE FROM users WHERE id = $1", id)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
//		return
//	}
//
//	rowsAffected, err := result.RowsAffected()
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
//		return
//	}
//
//	if rowsAffected == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//		return
//	}
//
//	c.Status(http.StatusOK)
//}
