package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetNumberDays(c *gin.Context) {
	year, err := strconv.Atoi(c.Params.ByName("year"))
	if err != nil {
	  c.String(http.StatusBadRequest, err.Error())
	  log.Println(err)
	  return
	}

	today := time.Now()
    inputYaer := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

    days := int(today.Sub(inputYaer).Hours())/24
	if days >=0 {
		c.JSON(http.StatusOK, "Days gone:"+ fmt.Sprint(days))
	} else {
		c.JSON(http.StatusOK, "Days left:"+ fmt.Sprint(-days))
	}
  
  }