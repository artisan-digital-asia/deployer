package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func ping(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
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
		if project != "" && branch != "" {
			// cmd := exec.Command("docker", "login")
			// stdout, err := cmd.Output()
			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }
			// fmt.Println(string(stdout))

			cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("docker contianer stop %s", project))
			stdout, _ := cmd.Output()
			fmt.Println(string(stdout))

			cmd = exec.Command("/bin/sh", "-c", fmt.Sprintf("docker pull %s/%s:%s", account, project, branch))
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println((string(stdout)))

			cmd = exec.Command("/bin/sh", "-c", fmt.Sprintf("docker run -d --rm --name %s %s/%s:%s", project, account, project, branch))
			stdout, err = cmd.Output()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println((string(stdout)))

			fmt.Fprintf(w, "Deployed")
			return
		}
	}
	http.Error(w, "Unauthorized", 401)
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/deploy", deploy)
	http.ListenAndServe(":8080", nil)
}
