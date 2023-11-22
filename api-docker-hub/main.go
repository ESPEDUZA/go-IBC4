package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Webhook reçu")

	decoder := json.NewDecoder(r.Body)
	var payload map[string]interface{}
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "Erreur de décodage JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	payloadBytes, _ := json.MarshalIndent(payload, "", "    ")
	fmt.Println("Payload reçu:", string(payloadBytes))

	imageTag := payload["push_data"].(map[string]interface{})["tag"].(string)
	fmt.Println("Tag de l'image poussée:", imageTag)

	cmd := exec.Command("./script.sh", imageTag)

	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	fmt.Println("Sortie du script:", stdout.String())
	fmt.Println("Erreur du script:", stderr.String())

	if err != nil {
		http.Error(w, "Erreur d'exécution du script: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/webhook", handleWebhook)

	fmt.Println("Serveur démarré et écoute sur le port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}
}
