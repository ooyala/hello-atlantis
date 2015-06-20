package hello

import (
	"log"
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"github.com/ooyala/go-dogstatsd"
)

var dCli *dogstatsd.Client

func InitDataDog() {
	var err error
	dCli, err = dogstatsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
	dCli.Namespace = "hello-go."
}

func Talk() string {
	return "Hello!"
}

func MsgHandler(msg string, printPath bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		err := dCli.Count("request.normal", 1, nil, 1)
		if err != nil {
			log.Printf("Error counting normal: %v", err)
		}
		output := msg
		if printPath {
			output += " " + r.URL.RequestURI()
		}

		w.Write([]byte("<!DOCTYPE html><html><head><title>hello-go</title></head><body><pre>\n"))
		w.Write([]byte(output + "\n"))
		w.Write([]byte("</pre></body></html>\n"))
	}
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	err := dCli.Count("request.healthz", 1, nil, 1)
	if err != nil {
		log.Printf("Error counting healthz: %v", err)
	}
	_, err = os.Stat("/etc/maint")
	if err == nil {
		w.Header().Add("Server-Status", "MAINTENANCE")
		http.Error(w, "maintenance mode", http.StatusNotFound)
		return
	}
	w.Header().Add("Server-Status", "OK")
	w.Write([]byte("OK\n"))
}

func StatuszHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	err := dCli.Count("request.statusz", 1, nil, 1)
	if err != nil {
		log.Printf("Error counting statusz: %v", err)
	}
	// gather output
	branchBytes, err := ioutil.ReadFile("/etc/atlantis/build/branch")
	if err != nil {
		http.Error(w, "Can't Read Branch File", 500)
		return
	}
	timeBytes, err := ioutil.ReadFile("/etc/atlantis/build/time")
	if err != nil {
		http.Error(w, "Can't Read Time File", 500)
		return
	}
	revlistBytes, err := ioutil.ReadFile("/etc/atlantis/build/revlist")
	if err != nil {
		http.Error(w, "Can't Read Revlist File", 500)
		return
	}
	output := "hello-go statusz\n"
	output += fmt.Sprintf("build branch: %s\n", branchBytes)
	output += fmt.Sprintf("build time: %s\n", timeBytes)
	if sha := r.FormValue("commit"); sha != "" {
		// check if commit exists
		if strings.Contains(fmt.Sprintf("\n%s\n", revlistBytes), fmt.Sprintf("\n%s", sha)) {
			output += fmt.Sprintf("YES - %s does exist.\n", sha)
		} else {
			output += fmt.Sprintf("NO - %s DOES NOT exist.\n", sha)
		}
	}
	output += "to check for a commit, set the query string parameter 'commit' to the sha to search for.\n"
	w.Write([]byte(output+"\n"))
}
