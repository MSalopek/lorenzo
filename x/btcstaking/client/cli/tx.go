package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	lrz "github.com/Lorenzo-Protocol/lorenzo/types"

	"github.com/spf13/cobra"

	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewCreateBTCStakingWithBTCProofCmd(),
	)

	return cmd
}

func NewCreateBTCStakingWithBTCProofCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-btcstaking-with-btc-proof [btc_tx_bytes] [proof]",
		Short: "Create a new btc staking request with proof from bitcoin-cli getrawtransaction&gettxoutproof output",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txBytes, err := hex.DecodeString(args[0])
			if err != nil {
				return fmt.Errorf("failed to decode tx bytes: %s", err)
			}
			proofRaw, err := hex.DecodeString(args[1])
			if err != nil {
				return fmt.Errorf("failed to decode proof bytes: %s", err)
			}
			merkleBlk, err := keeper.ParseMerkleBlock(proofRaw)
			if err != nil {
				return err
			}
			txIndex, proofBytes, err := keeper.ParseBTCProof(merkleBlk)

			blkHdr := &merkleBlk.Header

			var blkHdrHashBytes lrz.BTCHeaderHashBytes
			tmp := blkHdr.BlockHash()
			blkHdrHashBytes.FromChainhash(&tmp)

			msg := types.MsgCreateBTCStaking{
				Signer: clientCtx.GetFromAddress().String(),
				StakingTx: &types.TransactionInfo{
					Key: &types.TransactionKey{
						Index: txIndex,
						Hash:  &blkHdrHashBytes,
					},
					Transaction: txBytes,
					Proof:       proofBytes,
				},
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
