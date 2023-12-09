package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// ?
func (authenInfo authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + authenInfo.username + ":" + authenInfo.password
}

// don't touch below this line

func test5(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func ch4_6() {
	test5(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test5(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test5(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
