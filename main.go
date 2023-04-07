package main

import (
	"github.com/joho/godotenv"
	"github.com/studio-b12/gowebdav"
)

var webDavInstance *gowebdav.Client

func main() {
	envMap, err := godotenv.Read("./config.env")
	if err != nil {
		handleError("Failed to read config.env! Make sure this file exists!\nERROR: " + err.Error())
	}

	// Checking if the config is complete
	validateConfig(envMap)

	webDavInstance = gowebdav.NewClient("https://webdav."+envMap["IServInstanceHost"], envMap["Username"], envMap["Password"])

	handleArgs(envMap)

}
