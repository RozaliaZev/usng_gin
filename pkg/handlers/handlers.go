package handlers

import (
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
	daysStr := strconv.Itoa(days)
	if days >=0 {
		c.JSON(http.StatusOK, "Days gone:"+ daysStr)
	} else {
		c.JSON(http.StatusOK, "Days left:"+ daysStr)
	}
  
  }