package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"

	"github.com/klauspost/compress/s2"
	"go.uber.org/zap"
)

func getTraces(filename string) ([]int64, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	defer f.Close()
	var data map[string]map[string]int64
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		zap.S().Error(err)
		return nil, err
	}

	tracesMap := map[int64]bool{}
	for _, methods := range data {
		for _, height := range methods {
			tracesMap[height] = true
		}
	}
	traces := []int64{}
	for height := range tracesMap {
		traces = append(traces, height)
	}
	sort.Slice(traces, func(i, j int) bool {
		return traces[i] < traces[j]
	})
	return traces, nil
}

func writeToFile(path, filename string, data []byte) error {
	tmp, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("%s/%s", tmp, filename), data, fs.ModePerm)
}

func compress(format string, data []byte) ([]byte, error) {
	// Compress data using s2
	var b bytes.Buffer
	dataBuff := bytes.NewBuffer(data)

	var enc io.WriteCloser
	switch format {
	case "s2":
		enc = s2.NewWriter(&b)
	case "gz":
		enc = gzip.NewWriter(&b)
	default:
		return nil, fmt.Errorf("invalid format,expected s2 or gz")
	}

	_, err := io.Copy(enc, dataBuff)
	if err != nil {
		_ = enc.Close()
		return nil, err
	}
	// Blocks until compression is done.
	_ = enc.Close()

	return b.Bytes(), nil
}

func readGzFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("error creating gzip reader: %w", err)
	}
	defer gzipReader.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(gzipReader)
	if err != nil {
		return nil, fmt.Errorf("error reading from gzip reader: %w", err)
	}
	return buf.Bytes(), nil
}

func decompress(data []byte) ([]byte, error) {
	// Decompress data using s2
	b := bytes.NewBuffer(data)
	var out bytes.Buffer
	r := s2.NewReader(b)

	if _, err := io.Copy(&out, r); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
