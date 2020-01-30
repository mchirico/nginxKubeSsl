package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func pvv(file string) string {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal("could not open file")
	}
	defer f.Close()

	f.WriteString("this is a test\n")
	b2 := make([]byte, 2000)
	f.Seek(0, 0)
	n, _ := f.Read(b2)
	if n < 200 {
		return fmt.Sprintf("v:%s\n", b2[:n])
	}
	return fmt.Sprintf("v:%s\n", b2[:1999])

}

func mount() string {
	cmd := exec.Command("sh", "-c", "mount; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s\n", stdoutStderr)
}

func main() {
	// use PORT environment variable, or default to 8080
	port := "2342"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/", hello)

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, " storage Hello, world!  Update 2\n")
	fmt.Fprintf(w, "Version: 2.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
	//	fmt.Fprintf(w, "\n\n%s\n", mount())
	fmt.Fprint(w, "\n\n %s\n", pvv("/pvv/junk.txt"))
}
