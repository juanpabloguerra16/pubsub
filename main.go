package main

import (
	"bytes"
	"fmt"
	"os"
	"jppubsub.com/pubsubpublisher"
)



func main()  {
	var b bytes.Buffer
	projectID, ok := os.LookupEnv("projectID")
	if !ok {
		fmt.Println(ok)
	}
	topicID := "my-topic"
	msg := "Hello World"

	err := pubsubpublisher.Publish(&b, projectID, topicID, msg)
	if err != nil {
		fmt.Println(err)
	}
	
}