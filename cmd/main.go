package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	customerrepo "github.com/ahmedshaaban/intercom/repo/customer"
	customerservice "github.com/ahmedshaaban/intercom/service/customer"

	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	customerUrl string `envconfig:"CUSTOMER_URL"`
}

func main() {
	var c envConfig
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	repo, err := customerrepo.New(c.customerUrl, http.DefaultClient)
	if err != nil {
		log.Fatal(err)
	}

	service := customerservice.New(repo)
	invitess := service.SortedInvitees()
	out := ""

	for _, v := range invitess {
		out += fmt.Sprintf("ID: %d, Name: %s\n", v.ID, v.Name)
	}

	err = os.WriteFile("output.txt", []byte(out), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
