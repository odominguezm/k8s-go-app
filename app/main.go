package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Ruta principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola Orlando! App corriendo en Kubernetes v1.0\n")
	})

	// Health Check (Vital para SRE/K8s)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	fmt.Println("Servidor iniciado en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}
