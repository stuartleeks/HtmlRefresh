package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var backgroundColour string
var foregroundColour string
var fontSize string
var refreshRate string
var message string

func envVarOrDefault(envVarName string, defaultValue string) string {
	envValue, found := os.LookupEnv(envVarName)
	if found {
		return envValue
	}
	return defaultValue
}

func hello(w http.ResponseWriter, r *http.Request) {

	messageExpanded := os.ExpandEnv(message)
	html := fmt.Sprintf("<html><head><meta http-equiv=\"refresh\" content=\"%s\"></head><body style=\"background-color: %s;color: %s;font-size: %s;font-family: 'Open Sans',sans-serif; margin:0;padding:20px\">%s</body></head>", refreshRate, backgroundColour, foregroundColour, fontSize, messageExpanded)
	io.WriteString(w, html)
}

func main() {
	backgroundColour = envVarOrDefault("HTML_REFRESH_BGCOLOUR", "azure")
	foregroundColour = envVarOrDefault("HTML_REFRESH_FGCOLOUR", "black")
	fontSize = envVarOrDefault("HTML_REFRESH_FONTSIZE", "12vw")
	refreshRate = envVarOrDefault("HTML_REFRESH_RATE", "1")
	message = envVarOrDefault("HTML_REFRESH_MESSAGE", "$HOSTNAME")

	log.Printf("HTML_REFRESH_BGCOLOUR: %s\n", backgroundColour)
	log.Printf("HTML_REFRESH_BGCOLOUR: %s\n", foregroundColour)
	log.Printf("HTML_REFRESH_FONTSIZE: %s\n", fontSize)
	log.Printf("HTML_REFRESH_RATE: %s\n", refreshRate)
	log.Printf("HTML_REFRESH_MESSAGE: %s\n", message)

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
