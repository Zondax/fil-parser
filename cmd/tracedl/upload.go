package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/zondax/golem/pkg/cli"
	"go.uber.org/zap"
)

func upload(c *cli.CLI, cmd *cobra.Command, _ []string) {
	zap.S().Infof(c.GetVersionString())

	config, err := cli.LoadConfig[Config]()
	if err != nil {
		zap.S().Errorf("Error loading config: %s", err)
		return
	}

	traceFile, err := cmd.Flags().GetString("traceFile")
	if err != nil {
		zap.S().Errorf("Error loading outPath: %s", err)
		return
	}

	outPath, err := cmd.Flags().GetString("outPath")
	if err != nil {
		zap.S().Errorf("Error loading outPath: %s", err)
		return
	}

	if outPath == "" {
		zap.S().Errorf("outPath is required")
		return
	}

	dataStore, err := getDataStoreClient(config)
	if err != nil {
		zap.S().Error(err)
		return
	}

	file, err := os.Open(traceFile)
	if err != nil {
		zap.S().Errorf("Error opening traceFile: %s", err)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		zap.S().Errorf("Error getting file stats: %s", err)
		return
	}

	if err := dataStore.Client.UploadFromReader(file, fileStat.Size(), config.S3Bucket, outPath); err != nil {
		zap.S().Errorf("Error uploading traceFile: %s", err)
		return
	}

	zap.S().Infof("Trace file uploaded to %s: %s/%s", config.S3RawDataPath, config.S3Bucket, outPath)
}
