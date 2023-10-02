package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	//"weather/util"
)

func TestHandleGet(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	res := httptest.NewRecorder()

	handleGet(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d , but got %d", http.StatusOK, res.Code)
	}
}

func TestFilerByQuery(t *testing.T) {
	//Test 1
	req1 := httptest.NewRequest("GET", "/", nil)
	res1 := httptest.NewRecorder()
	filerByQuery(res1, req1)

	//Test 2
	req2 := httptest.NewRequest("GET", "/query?limit=5", nil)
	res2 := httptest.NewRecorder()
	filerByQuery(res2, req2)

	//Test 3
	req3 := httptest.NewRequest("GET", "/query?weather=sun", nil)
	res3 := httptest.NewRecorder()
	filerByQuery(res3, req3)

	//Test 4
	req4 := httptest.NewRequest("GET", "/query?date=2015-12-31", nil)
	res4 := httptest.NewRecorder()
	filerByQuery(res4, req4)

	//Test 5
	req5 := httptest.NewRequest("GET", "/query?weather=sun&limit=3", nil)
	res5 := httptest.NewRecorder()
	filerByQuery(res5, req5)

}
