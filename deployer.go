package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"os"
)

func ping(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		log.Println("/ping")
		fmt.Fprintf(w, "pong")
		return
	}
	http.Error(w, "Unauthorized", 401)
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func deploy(w http.ResponseWriter, req *http.Request) {
	secret := getEnv("SECRET", "deploymenow") 
	if req.Method == "POST" && req.Header.Get("secret") == secret {
		account := getEnv("DOCKERHUB_ACCOUNT", "artisandigitalasia")
		project := req.Header.Get("project")
		branch := req.Header.Get("branch")
		log.Printf("/deploy %s:%s\n", project, branch)
		if project != "" && branch != "" {
			cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo docker contianer stop %s", project))
			stdout, _ := cmd.Output()
			log.Println(string(stdout))

			cmd = exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo docker pull %s/%s:%s", account, project, branch))
			stdout, err := cmd.Output()
			if err != nil {
				log.Println(err)
				return
			}
			log.Println((string(stdout)))

			cmd = exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo docker run -d --rm --name %s %s/%s:%s", project, account, project, branch))
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
	port := getEnv("PORT", "8080")
	log.Println("SECRET:", getEnv("SECRET", "deploymenow"))
	log.Println("PORT:", port)
	log.Printf("🚀 Deployer launched on port %s\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "Unauthorized", 401)
	})
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/deploy", deploy)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
