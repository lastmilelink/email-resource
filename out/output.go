package main

type Output struct {
	Version  string           `json:"version"`
	Metadata []OutputMetadata `json:"metadata"`
}

type OutputMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
