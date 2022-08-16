package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	source := "./examples"
	pattern := "/go-optimization-by-examples/"
	port := 3000
	fs := http.FileServer(http.Dir(source))

	http.Handle(pattern, http.StripPrefix(pattern, fs))

	log.Print("Listening on :3000...")

	go openbrowser(fmt.Sprintf("http://localhost:%d/%s", port, pattern))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
