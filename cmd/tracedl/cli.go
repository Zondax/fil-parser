package main

import (
	"github.com/spf13/cobra"
	"github.com/zondax/golem/pkg/cli"
)

func GetStartCommand(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get",
		Run: func(cmd *cobra.Command, args []string) {
			download(c, cmd, args)
		},
	}
	cmd.Flags().String("type", "traces", "--type traces")
	cmd.Flags().String("outPath", ".", "--outPath ../")
	cmd.Flags().String("compress", "gz", "--compress s2")
	cmd.Flags().UintSlice("heights", []uint{387926}, "--heights 387926")
	cmd.Flags().Bool("useDataStore", false, "--useDataStore true")
	return cmd
}

func GetActorParseCommand(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "actor-parse",
		Short: "Actor Parse",
		Run: func(cmd *cobra.Command, args []string) {
			parse(c, cmd, args)
		},
	}
	cmd.Flags().String("tracesPath", ".", "--tracesPath .")
	cmd.Flags().Uint64("height", 387926, "--height 387926")
	cmd.Flags().Bool("useDataStore", false, "--useDataStore true")
	cmd.Flags().String("actorAddress", "", "--actorAddress f01")
	cmd.Flags().String("actorName", "", "--actorName account")
	cmd.Flags().String("actorMethod", "", "--actorMethod Constructor")
	cmd.Flags().Bool("parseSubTxs", false, "--parseSubTxs true")
	return cmd
}
