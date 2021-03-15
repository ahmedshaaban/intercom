package customer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Customer struct {
	ID        int    `json:"user_id"`
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type client interface {
	Get(url string) (resp *http.Response, err error)
}

type Repo struct {
	url       string
	client    client
	customers []Customer
}

func New(url string, client client) (*Repo, error) {
	r := &Repo{client: client, customers: []Customer{}, url: url}
	err := r.fillCustomers()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Repo) GetCustomers() []Customer {
	return r.customers
}

func (r *Repo) fillCustomers() error {
	resp, err := r.client.Get("https://s3.amazonaws.com/intercom-take-home-test/customers.txt#")
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	customersString := strings.Split(string(body), "\n")
	for _, v := range customersString {
		var customer Customer
		err := json.Unmarshal([]byte(v), &customer)
		if err != nil {
			return err
		}
		r.customers = append(r.customers, customer)
	}

	return nil
}
