package main

import (
	"io"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func webDavDownload() {
	if len(os.Args) < 4 {
		handleError("Remote Path/Filename missing!")
	}
	if len(os.Args) < 5 {
		handleError("Local Path/Filename missing!")
	}
	remoteReader, err := webDavInstance.ReadStream(os.Args[3])
	if err != nil {
		handleError("An Error Occurred!\nERROR: " + err.Error())
	}
	localFile, err := os.Create(os.Args[4])
	if err != nil {
		handleError("An Error Occurred!\nERROR: " + err.Error())
	}
	defer localFile.Close()

	io.Copy(localFile, remoteReader)
}

func webDavList() {
	if len(os.Args) < 4 {
		handleError("Remote Path/Filename missing!")
	}
	files, err := webDavInstance.ReadDir(os.Args[3])
	if err != nil {
		handleError("An Error Occurred!\nERROR: " + err.Error())
	}
	for _, file := range files {
		var size string
		// TODO: Simplify this chunk
		if file.Size() > 1000000000 {
			size = strconv.Itoa(int(file.Size()/1000000000)) + "Gb"
		} else if file.Size() > 1000000 {
			size = strconv.Itoa(int(file.Size()/1000000)) + "Mb"
		} else if file.Size() > 1000 {
			size = strconv.Itoa(int(file.Size()/1000)) + "Kb"
		} else {
			size = strconv.Itoa(int(file.Size())) + "B"
		}
		if file.IsDir() {
			// TODO: Somehow simplify this long line of code
			color.Blue(size + "	" + file.ModTime().Format("Jan 02 15:04:05 2006") + "	" + file.Name() + "/")
		} else {
			color.Blue(size + "	" + file.ModTime().Format("Jan 02 15:04:05 2006") + "	" + file.Name())
		}
	}
}
