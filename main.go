// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func rebootHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("/usr/sbin/reboot")
	if err := cmd.Start(); err != nil {
		http.Error(w, "failed to execute reboot: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "rebooting")
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("/usr/sbin/shutdown", "-h", "now")
	if err := cmd.Start(); err != nil {
		http.Error(w, "failed to execute shutdown: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "shutting down")
}

func main() {
	http.HandleFunc("/reboot", rebootHandler)
	http.HandleFunc("/shutdown", shutdownHandler)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
