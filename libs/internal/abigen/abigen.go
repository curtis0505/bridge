package abigen

import (
	"fmt"
	"github.com/curtis0505/bridge/libs/internal/abigen/bind"
	"github.com/curtis0505/bridge/libs/logger"
	"github.com/spf13/cobra"
	"io"
	"os"
)

const (
	Generate = "generate"
)

func abigen() *cobra.Command {
	c := &cobra.Command{
		Use: Generate,
		Run: func(cmd *cobra.Command, args []string) {
			pkgFlag, _ := cmd.Flags().GetString("pkg")
			if pkgFlag == "" {
				logger.Fatalf("No destination package specified (--pkg)")
			}
			abiFlag, _ := cmd.Flags().GetString("abi")
			if abiFlag == "" {
				logger.Fatalf("No input ABI specified (--abi)")
			}
			typeFlag, _ := cmd.Flags().GetString("type")
			outFlag, _ := cmd.Flags().GetString("out")

			// Default output typing Language is Go
			lang := bind.LangGo

			// If the entire solidity code was specified, build and bind based on that
			var (
				abis    []string
				bins    []string
				types   []string
				sigs    []map[string]string
				libs    = make(map[string]string)
				aliases = make(map[string]string)
			)
			// Load up the ABI, optional bytecode and type name from the parameters
			var (
				abi []byte
				err error
			)

			input := abiFlag
			if input == "-" {
				abi, err = io.ReadAll(os.Stdin)
			} else {
				abi, err = os.ReadFile(input)
			}
			if err != nil {
				logger.Fatalf("Failed to read input ABI: %v", err)
			}
			abis = append(abis, string(abi))
			var bin []byte
			bins = append(bins, string(bin))
			kind := typeFlag
			if kind == "" {
				kind = pkgFlag
			}
			types = append(types, kind)

			code, err := bind.Bind(types, abis, bins, sigs, pkgFlag, lang, libs, aliases)
			if err != nil {
				logger.Fatalf("Failed to generate ABI binding: %v", err)
			}

			if outFlag == "" {
				fmt.Printf("%s\n", code)
				return
			} else {
				if err := os.WriteFile(outFlag, []byte(code), 0600); err != nil {
					logger.Fatalf("Failed to write ABI binding: %v", err)
				}
			}
		}}
	return c
}
