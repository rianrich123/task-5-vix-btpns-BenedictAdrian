package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Status string `json:"status"`
	Info string `json:"info"`
}

func handlePage(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	var message Message
	err := json.NewDecoder(request.Body).Decode(&message)
	if err != nil {
		return
	}
	err = json.NewEncoder(writer).Encode(message)
	if err != nil {
		return
	}
}

func generateJWT()(string, error){
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"
}

tokenString, err:= token.SignedString(sampleSecretKey)
if err != nil{
	return "", err
}

return tokenString, nil

func verifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {

}

return http.HandlerFunc(func(writer http.ResponseWriter,request *http.Request){

})

if request.Header["Token"] != nil {

}

token, err := jwt.Parse(request.Header\["Token"]\)[0], func(token *jwt.Token) (interface{}, error){
	_, ok := token.Method.(*jwt.SigningMethodECDSA)

}

if token.Valid {
	endpointHandler(writer, requestd)
} else {
	writer.WriteHeader(http.StatusUnauthorized)
	_, err := writer.Write([]byte ("You're Unauthorized due to invalid token"))
	if err != nil{
		return
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
		_, err := writer.Write([]byte("You're Unauthorized due to No token in the header"))
		if err != nil {
			return
		}
	}
}


type user struct {
	UserID        string `json:"userid"`
	PhotoRelation string `json:"photorelation"`
	CreateAt      string `json:"createat"`
	UpdatedAt     string `json:"updateat"`
}

var (
	password = os.Getenv("MSSQL_DB_PASSWORD")
	email = os.Getenv("MSSQL_DB_EMAIL")
	username = os.Getenv("MSSQL_DB_USER")
	port = os.Getenv("MSSQL_DB_PORT")
	database = os.Getenv("MSSQL_DB_DATABASE")
)

connectionString := fmt.Println("user id=%s;password=%s;port=%s;database=%s", password, email, username, port, database)

sqlObj, connectionError := sql.Open("mssql", database.connectionString); if connectionError != nil {
	fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
}

data := database.Database{
	SqlDb: sqlObj,
}

fmt.Println("-> Welcome to reminders Console App, build using Golang and Microsoft SQL Server")
fmt.Println("-> Select a numeric option; \n [1] Create a new Reminder \n [2] Get a reminder \n [3] Delete a reminder")

consoleReader := bufio.NewScanner(os.Stdin)
consoleReader.Scan()
userChoice := consoleReader.Text()

switch userChoice {
case "1":
	var (
		titleInput,
		descriptionInput,
		aliasInput string
	)
}

type photo struct {
	PhotoID      string `json:"photoid"`
	Title        string `json:"title"`
	Caption      string `json:"caption"`
	PhotoUrl     string `json:"photourl"`
	UserRelation string `json:"userrelation"`
}

var photos = []photo{
	{PhotoID: "1", Title: "Pas Foto", Caption: "Pas Foto Kerja Ben", PhotoUrl: "bit.ly/PasFotoBen", UserRelation: "hehehe"},
}

func getPhoto(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, photos)
}

func createPhoto(c *gin.Context) {
	var newPhoto photo

	if err := c.BindJSON(&newPhoto); err != nil {
		return
	}

	photos = append(photos, newPhoto)
	c.IndentedJSON(http.StatusCreated, newPhoto)
}

func Update(response http.ResponseWriter, request *http.Request) {
	var pic entities.photo
	err := json.NewDecoder(request.Body).Decode(&pic)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		fmt.Println("Pic Info - Updated")
		fmt.Println("id:		", pic.PhotoId)
		fmt.Println("title:		", pic.Title)
		fmt.Println("caption:	", pic.Caption)
		fmt.Println("photo url :", pic.PhotoUrl)
		fmt.Println("user relation :", pic.UserRelation)
		respondWithJSON(response, http.StatusOK, pic)
	}
}

func respondWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func respondWithError(response http.ResponseWriter, statusCode int, msg string) {
	respondWithJSON(response, statusCode, map[string]string{"error": msg})
}

func deletePhoto(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	photoid := vars["photoid"]
	respondWithJSON(response, http.StatusOK, photo{PhotoId: photoid})
	title := vars["title"]
	respondWithJSON(response, http.StatusOK, photo{Title: title})
	caption := vars["caption"]
	respondWithJSON(response, http.StatusOK, photo{Caption: caption})
	photourl := vars["photourl"]
	respondWithJSON(response, http.StatusOK, photo{PhotoUrl: photourl})
	userrelation := vars["userrelation"]
	respondWithJSON(response, http.StatusOK, photo{UserRelation: userrelation})
}

func main() {

	router := gin.Default()
	router.GET("/photo", getPhoto)
	router.POST("/photo", createPhoto)
	router.DELETE("/photo", deletePhoto)
	router.Run("localhost:8080")
}
