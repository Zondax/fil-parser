package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/spf13/cobra"
	"github.com/zondax/golem/pkg/cli"
	"golang.org/x/xerrors"
)

func GetGasCommand(c *cli.CLI) *cobra.Command {
	fmt.Println("Loading config...")
	cmd := &cobra.Command{
		Use:   "gas",
		Short: "Gas",
		Run: func(cmd *cobra.Command, args []string) {
			runCheck(c, cmd, args)
		},
	}
	cmd.Flags().Int64("targetHeight", 255233, "--targetHeight 255233")
	cmd.Flags().String("myPremium", "1000000000000000000", "--myPremium 1000000000000000000")
	cmd.Flags().Int64("prevHeights", 10, "--prevHeights 10")
	return cmd
}

func runCheck(c *cli.CLI, cmd *cobra.Command, _ []string) error {
	ctx := context.Background()
	fmt.Println("Loading config...")
	config, err := cli.LoadConfig[Config]()
	if err != nil {
		fmt.Printf("Error loading config: %s", err)
		return err
	}
	rpcClient, err := newFilecoinRPCClient(config.NodeURL, config.NodeToken)
	if err != nil {
		fmt.Printf("Error loading config: %s", err)
		return err
	}

	targetHeight, err := cmd.Flags().GetInt64("targetHeight")
	if err != nil {
		fmt.Printf("Error getting targetHeight: %s", err)
		return err
	}

	myPremium, err := cmd.Flags().GetString("myPremium")
	if err != nil {
		fmt.Printf("Error getting myPremium: %s", err)
		return err
	}

	prevHeights, err := cmd.Flags().GetInt64("prevHeights")
	if err != nil {
		fmt.Printf("Error getting prevHeights: %s", err)
		return err
	}

	myPremiumInt, ok := new(big.Int).SetString(myPremium, 10)
	if !ok {
		return xerrors.Errorf("invalid format for my-premium: %s", myPremium)
	}

	fmt.Printf("Comparing GasPremiums around height %d (checking %d to %d)\n", targetHeight, targetHeight-prevHeights, targetHeight)
	fmt.Printf("Your message premium: %s\n", myPremiumInt.String())
	fmt.Println("-----------------------------------------------------")

	startHeight := targetHeight - prevHeights
	if startHeight < 0 {
		startHeight = 0
	}

	for h := startHeight; h <= targetHeight; h++ {
		fmt.Printf("\n[Height %d]\n", h)

		ts, err := rpcClient.client.ChainGetTipSetByHeight(ctx, abi.ChainEpoch(h), types.EmptyTSK)
		if err != nil {
			fmt.Printf("  Error getting tipset: %v\n", err)
			continue // Skip this height if we can't get the tipset
		}

		if ts == nil || len(ts.Blocks()) == 0 {
			fmt.Println("  Tipset is nil or empty.")
			continue
		}

		fmt.Printf("  Tipset CIDs: %s\n", ts.Key().String())

		// for _, blockHeader := range ts.Blocks() {
		// 	fmt.Printf("  Block: %s (Miner: %s)\n", blockHeader.Cid(), blockHeader.Miner)

		blockMsgs, err := rpcClient.client.ChainGetMessagesInTipset(ctx, ts.Key())
		if err != nil {
			fmt.Printf("    Error getting block messages: %v\n", err)
			continue // Skip this block if messages can't be retrieved
		}

		var premiums []*big.Int

		for _, msg := range blockMsgs {
			premiums = append(premiums, msg.Message.GasPremium.Int)
		}

		if len(premiums) == 0 {
			fmt.Println("    No Messages with GasPremium found in this block.")
			tmp, _ := json.Marshal(blockMsgs)
			fmt.Println(string(tmp))
			continue
		}

		sort.Slice(premiums, func(i, j int) bool {
			return premiums[i].Cmp(premiums[j]) < 0 // Sort ascending
		})

		highestSeen := premiums[len(premiums)-1] // Last element after sort
		sumPremium := new(big.Int)
		for _, p := range premiums {
			sumPremium.Add(sumPremium, p)
		}
		avgPremium := new(big.Int).Div(sumPremium, big.NewInt(int64(len(premiums))))

		fmt.Printf("    Messages: %d | Highest Premium: %s | Avg Premium: %s\n",
			len(premiums), highestSeen.String(), avgPremium.String())

		if highestSeen.Cmp(myPremiumInt) > 0 {
			fmt.Printf("    * This block contained messages with premium HIGHER than yours (%s).\n", myPremiumInt.String())
		} else {
			fmt.Printf("    * All included message premiums were <= yours (%s).\n", myPremiumInt.String())
		}
	}
	// }

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Comparison complete.")
	return nil
}
