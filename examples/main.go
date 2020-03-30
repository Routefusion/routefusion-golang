package main

import (
	"fmt"

	routefusion "github.com/routefusion/routefusion-golang"
)

func main() {
	rf := routefusion.New(routefusion.
		Config{URL: "https://sandbox.api.routefusion.co",
		ClientID:  "0C270CB98CFD4DE4BBD45BA179B67307F263DA92C66572A2A8D043E55FC826B5",
		SecretKey: "723Ae7657b97722DA2D4a964EB1D137d21B4A81e88AF368F1e990B8b2F5727C0",
	})
	user, err := rf.GetUser()
	fmt.Println(user, err)

	updatedUser, err := rf.UpdateUser(&routefusion.User{UserName: "sammy1"})
	fmt.Println(updatedUser, err)

}
