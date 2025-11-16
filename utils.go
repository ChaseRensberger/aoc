package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func loadInputFile(year int, day int) (string, error) {
	inputPath := fmt.Sprintf("inputs/%d/day%d/input.txt", year, day)

	if _, err := os.Stat(inputPath); err != nil {
		input, err := retrieveInputFromRemote(year, day)
		if err != nil {
			return "", err
		}

		dir := fmt.Sprintf("inputs/%d/day%d", year, day)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}

		if err := os.WriteFile(inputPath, []byte(input), 0644); err != nil {
			return "", err
		}
	}

	data, err := os.ReadFile(inputPath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func retrieveInputFromRemote(year int, day int) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	sessionCookie := os.Getenv("AOC_SESSION")
	if sessionCookie == "" {
		return "", fmt.Errorf("AOC_SESSION environment variable not set")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch input: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return string(body), nil
}
