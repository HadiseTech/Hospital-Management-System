package Admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/web1_group_project/hospital_client/entity"
)

var baseURLPhrmacist = "http://localhost:8182/admin/pharmasists/"

func GetPharmacist(id uint) (*entity.Pharmacist, error) {
	fmt.Println("Enteritn")
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURLPhrmacist, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	presc := &entity.Pharmacist{}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, presc)
	if err != nil {
		return nil, err
	}
	fmt.Println("pres", presc)
	return presc, nil
}
func Pharmacists() ([]entity.Pharmacist, error) {
	client := &http.Client{}
	sp := ""
	URL := fmt.Sprintf("%s%s", baseURLPhrmacist, sp)
	req, _ := http.NewRequest("GET", URL, nil)
	res, _ := client.Do(req)
	var prescs = make([]entity.Pharmacist, 0)
	body, _ := ioutil.ReadAll(res.Body)

	_ = json.Unmarshal(body, &prescs)

	return prescs, nil
}

func PutPharmacist(presc *entity.Pharmacist) {
	client := &http.Client{}
	id := presc.Uuid
	t := strconv.Itoa(int(id))
	json, err := json.Marshal(presc)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", baseURLPhrmacist+t, bytes.NewBuffer(json))
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	fmt.Println("id put", id)
	fmt.Println("put done")
}

func DeletePharmacist(ID uint) {
	client := &http.Client{}
	t := strconv.Itoa(int(ID))
	// Create request
	req, err := http.NewRequest("DELETE", baseURLPhrmacist+t, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

}

func PostPharmacist(presc *entity.Pharmacist) {
	client := &http.Client{}

	json, err := json.Marshal(presc)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", baseURLPhrmacist, bytes.NewBuffer(json))
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	fmt.Println("post done with", presc)
}