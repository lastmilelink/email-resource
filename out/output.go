package main

import (
	"strconv"
	"time"
)

type Output struct {
	Version  string           `json:"version"`
	Metadata []OutputMetadata `json:"metadata"`
}

type OutputMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func generateOutput(params Parameters) Output {
	var out Output
	out.Version = strconv.Itoa(time.Now().UTC().Nanosecond())

	out.Metadata = []OutputMetadata{
		OutputMetadata{Key: "TopicName", Value: params.TopicName},
	}

	return out
}
