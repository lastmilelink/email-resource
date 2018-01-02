package main

import "time"

type Output struct {
	Version  int64            `json:"version"`
	Metadata []OutputMetadata `json:"metadata"`
}

type OutputMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func generateOutput(params Parameters) Output {
	var out Output
	out.Version = time.Now().Unix()

	out.Metadata = []OutputMetadata{
		OutputMetadata{Key: "TopicName", Value: params.TopicName},
	}

	return out
}
