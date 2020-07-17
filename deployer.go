package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"log"
)

func ping(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		log.Println("/ping")
		fmt.Fprintf(w, "pong")
		return
	}
	http.Error(w, "Unauthorized", 401)
}

func deploy(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" && req.Header.Get("secret") == "deploymenow" {
		account := "artisandigitalasia"
		project := req.Header.Get("project")
		branch := req.Header.Get("branch")
		log.Printf("/deploy %s:%s\n", project, branch)
		if project != "" && branch != "" {
			cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("docker contianer stop %s", project))
			stdout, _ := cmd.Output()
			log.Println(string(stdout))

			cmd = exec.Command("/bin/sh", "-c", fmt.Sprintf("docker pull %s/%s:%s", account, project, branch))
			stdout, err := cmd.Output()
			if err != nil {
				log.Println(err)
				return
			}
			log.Println((string(stdout)))

			cmd = exec.Command("/bin/sh", "-c", fmt.Sprintf("docker run -d --rm --name %s %s/%s:%s", project, account, project, branch))
			stdout, err = cmd.Output()
			if err != nil {
				log.Println(err)
				return
			}
			log.Println((string(stdout)))

			fmt.Fprintf(w, "Deployed")
			return
		}
	}
	http.Error(w, "Unauthorized", 401)
}

func main() {
	port := "8080"
	log.Printf("🚀 Deployer launched on port %s\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "Unauthorized", 401)
	})
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/deploy", deploy)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
