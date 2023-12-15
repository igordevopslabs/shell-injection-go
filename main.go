package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type fixedResponse string

func ExecuteCommand(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query().Get("cmd")

	log.Println("Comando recebido:", cmd)

	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func (s fixedResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s)
}

func main() {
	http.HandleFunc("/inject", ExecuteCommand)
	http.Handle("/", fixedResponse("~> shell-injection-go is ready..."))
	http.Handle("/ready", fixedResponse("~> shell-injection-go is ready..."))
	http.Handle("/liveness", fixedResponse("~> shell-injection-go is live..."))
	log.Fatal(http.ListenAndServe(":9000", nil))
	log.Default()
}
