package main

import "time"

type Output struct {
	Version  string           `json:"version"`
	Metadata []OutputMetadata `json:"metadata"`
}

type OutputMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func generateOutput(params EmailParams) Output {
	var out Output
	out.Version = time.Now().String()

	out.Metadata = []OutputMetadata{
		OutputMetadata{Key: "TopicName", Value: params.TopicName},
	}

	return out
}
