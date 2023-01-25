package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router
}
func TestGetNumberDays(t *testing.T) {
	router := SetUpRouter()
	router.GET("/when/:year", GetNumberDays)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/when/2000", nil)
	router.ServeHTTP(w, r)
	
	str, _ := ioutil.ReadAll(w.Body)
    assert.Equal(t, "Days gone:8425", string(str))
    assert.Equal(t, http.StatusOK, w.Code)
}