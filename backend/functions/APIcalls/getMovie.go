package apiCall

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//************************************************************************************************************************************************************************

//chnage movie struct if needed ,,,its different for queries
type FilmId struct {
	Poster_path       string
	Adult             bool
	Overview          string
	Release_date      string
	Genres            []Genre
	Id                int32
	Original_title    string
	Original_language string
	Vote_average      float32
	Title             string
	Backdrop_path     string
	Popularity        float64
	Vote_count        int32
	Video             bool
}

type Genre struct {
	Id   int32
	Name string
}

var filmId FilmId

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("7 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************

func GetMovie(c *gin.Context) {

	//Communication
	movieID := c.Param("movie")

	response, err := http.Get("https://api.themoviedb.org/3/movie/" + movieID + "?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(responseData), &filmId)

	c.JSON(http.StatusOK, gin.H{
		"movie": filmId,
	})

	for i := 0; i < len(filmId.Genres); i++ {
		GetMoviesOfAGenra(c, filmId.Genres[i])
	}

}

func GetMoviesOfAGenra(c *gin.Context, genre Genre) {
	//Communication
	genId := genre.Name
	response, err := http.Get("https://api.themoviedb.org/3/discover/movie?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US&sort_by=popularity.desc&page=1&with_genres=" + genId + "&with_watch_monetization_types=flatrate")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(responseData), &api)

	c.JSON(http.StatusOK, gin.H{
		genId: api,
	})
}

//************************************************************************************************************************************************************************
