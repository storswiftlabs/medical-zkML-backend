package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"medical-zkml-backend/pkg/config"
	"net/http"
)

func init() {
	config.NewConfig()
}

func IPFS(data []string) (string, error) {

	url := config.GetConfig().Get("ipfs.url").(string)
	var byteArr []byte
	for _, str := range data {
		byteArr = append(byteArr, []byte(str)...)
	}
	payload := bytes.NewBuffer(byteArr)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}
	// TODO: Put configuration information into the configuration file
	auth := "Bearer " + config.GetConfig().Get("ipfs.auth").(string)
	// Set the request header, including your NFT.Storage API key
	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "text/plain")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Check response
	if resp.Status != "200 OK" {
		fmt.Println("Error uploading data. Status:", resp.Status)
		return "", err
	}

	// Extract CID (Content Identifier) from the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return "", err
	}

	cid := result["value"].(map[string]interface{})["cid"].(string)
	fmt.Println("CID:", cid)
	return cid, nil
}
