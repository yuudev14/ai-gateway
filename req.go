package main

// func main() {
// 	url := "http://172.18.32.1:11434/api/chat"
// 	jsonData := `{
//         "model":"qwen3:8b",
//         "messages":[{"role":"user","content":"What is the capital of japan?"}],
//         "stream":true
//     }`

// 	req, err := http.NewRequest("POST", url, strings.NewReader(jsonData))
// 	if err != nil {
// 		panic(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	// Read streaming response line by line
// 	scanner := bufio.NewScanner(resp.Body)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		fmt.Println(line)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		panic(err)
// 	}
// }
