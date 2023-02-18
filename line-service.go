package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	token := ""                            // Replace with Line Notify API token
	message := "Message From Line Notify!" // Replace with message

	// Build the HTTP request
	client := &http.Client{}
	values := url.Values{}
	values.Set("message", message)
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Println("Error building HTTP request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status code and body
	fmt.Println("Response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:", string(body))
	fmt.Println("Message :" + message)

}
