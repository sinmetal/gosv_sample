package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	metadatabox "github.com/sinmetalcraft/gcpbox/metadata"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは\n")

	runRevision := os.Getenv("K_REVISION")
	if runRevision != "" {
		fmt.Fprintf(w, "Cloud Run Revision: %s\n", runRevision)
	}

	instanceID, err := metadatabox.InstanceID()
	if err != nil {
		fmt.Printf("failed metadatabox.InstanceID() err=%s", err)
	}
	if instanceID != "" {
		fmt.Fprintf(w, "Instance ID: %s\n", instanceID)
	}
}
