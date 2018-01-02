package main

import (
	"time"
)

type Version struct {
	Time time.Time
}

type Output struct {
	Version  Version          `json:"version"`
	Metadata []OutputMetadata `json:"metadata"`
}

type OutputMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func generateOutput(params Parameters) Output {
	var out Output

	out.Version = Version{Time: time.Now().UTC()}

	out.Metadata = []OutputMetadata{
		OutputMetadata{Key: "TopicName", Value: params.TopicName},
	}

	return out
}
