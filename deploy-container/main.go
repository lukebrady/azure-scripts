package main

import (
	"github.com/Azure/azure-sdk-for-go/services/containerinstance/mgmt/2018-09-01/containerinstance"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	client := containerinstance.NewContainerClient("")
	authorize, err := auth.NewAuthorizerFromFile("client_credentials.json")
	if err == nil {
		client.Authorizer = authorize
	}
}
