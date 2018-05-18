package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func envVarOrDefault(envVarName string, defaultValue string) string {
	envValue, found := os.LookupEnv(envVarName)
	if found {
		return envValue
	}
	return defaultValue
}
func hello(w http.ResponseWriter, r *http.Request) {

	backgroundColour := envVarOrDefault("HTML_REFRESH_BGCOLOUR", "azure")
	foregroundColour := envVarOrDefault("HTML_REFRESH_FGCOLOUR", "black")
	fontSize := envVarOrDefault("HTML_REFRESH_FONTSIZE", "12vw")
	refreshRate := envVarOrDefault("HTML_REFRESH_RATE", "1")
	message := envVarOrDefault("HTML_REFRESH_MESSAGE", "$HOSTNAME")

	messageExpanded := os.ExpandEnv(message)
	html := fmt.Sprintf("<html><head><meta http-equiv=\"refresh\" content=\"%s\"></head><body style=\"background-color: %s;color: %s;font-size: %s;font-family: 'Open Sans',sans-serif; margin:0;padding:20px\">%s</body></head>", refreshRate, backgroundColour, foregroundColour, fontSize, messageExpanded)
	io.WriteString(w, html)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
