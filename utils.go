package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

// Acts like a snippet so I dont have to write those two lines over and over
func handleError(errMsg string) {
	color.Red(errMsg)
	os.Exit(-1)
}

// TODO: Organize and optimize this chunk of text
func commandsHelp() {
	fmt.Println(`IServ CLI help
  Available Commands:
    help: ( shows this text )
    webdav:
      list [remote path] ( lists all files in the given remote repository. Similar to using 'ls' on a UNIX-based system )
      download [remote file] [local path + filename] ( downloads the given remote file at the given local path + filename )
	
	`)
}

func validateConfig(envMap map[string]string) {
	if envMap["IServInstanceHost"] == "" {
		handleError("IServInstanceHost is missing! Exiting...")
	}
	if envMap["Username"] == "" {
		reader := bufio.NewReader(os.Stdin)
		color.Blue("-- Username is missing! Please provide below --")
		for envMap["Username"] == "" {
			fmt.Print(": ")
			username, err := reader.ReadString('\n')
			if err != nil {
				handleError("Failed to read username\nERROR: " + err.Error())
			}
			envMap["Username"] = strings.Trim(username, "\n")
			if envMap["Username"] == "" {
				color.Red("Username cannot be blank!")
			} else {
				godotenv.Write(envMap, "./config.env")
			}
		}
	}

	if envMap["Password"] == "" {
		reader := bufio.NewReader(os.Stdin)
		color.Blue("-- Password is missing! Please provide below --")
		for envMap["Password"] == "" {
			fmt.Print(": ")
			password, err := reader.ReadString('\n')
			if err != nil {
				handleError("Failed to read password\nERROR: " + err.Error())
			}
			envMap["Password"] = strings.Trim(password, "\n")
			if envMap["Password"] == "" {
				color.Red("Password cannot be blank!")
			} else {
				godotenv.Write(envMap, "./config.env")
			}
		}
	}
}

// TODO: Same as the function above.. Oranize and optimize this
func handleArgs(envMap map[string]string) {
	if len(os.Args) <= 1 {
		commandsHelp()
		os.Exit(0)
	}
	switch strings.ToLower(os.Args[1]) {
	case "help":
		commandsHelp()
	case "webdav":
		switch strings.ToLower(os.Args[2]) {
		case "download":
			webDavDownload()
		case "list":
			webDavList()
		}
	case "email":
		switch strings.ToLower(os.Args[2]) {
		case "send":
			sendEmail(envMap)
		}
	}
}
