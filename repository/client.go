package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type apiRepository struct {
	URL   string
	AppId string
}

func NewApiRepository(url, appId string) *apiRepository {
	return &apiRepository{URL: url,
		AppId: appId,
	}
}
func ExternalStoreUrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	store := r.URL.Query()["q"]
	fmt.Println(store)
	if len(store) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - No store was found!"))
		return
	}

	url := fmt.Sprintf("https://localhost:4000/order")

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to read url", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Unable to read body", err)
	}

	var data map[string]interface{}

	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(data)

}
