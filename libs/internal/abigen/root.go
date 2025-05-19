package abigen

import "github.com/spf13/cobra"

var (
	// abigen --abi abi.json --pkg derivative --type EthBoost --out derivative.go
	abiGenCmd = &cobra.Command{
		Use:   "abigen",
		Short: "generate go typings from abi (EVM only)",
	}
)

func GetCommand() *cobra.Command {
	abiGenCmd.PersistentFlags().StringP("abi", "a", "", "Path to the Ethereum contract ABI json to bind, - for STDIN\"")
	abiGenCmd.PersistentFlags().StringP("type", "t", "", "Struct name for the binding (default = package name)")
	abiGenCmd.PersistentFlags().StringP("pkg", "p", "", "Package name to generate the binding into")
	abiGenCmd.PersistentFlags().StringP("out", "o", "", "Output file for the generated binding (default = stdout)")

	err := abiGenCmd.MarkPersistentFlagRequired("abi")
	if err != nil {
		panic(err)
	}
	err = abiGenCmd.MarkPersistentFlagRequired("type")
	if err != nil {
		panic(err)
	}
	err = abiGenCmd.MarkPersistentFlagRequired("pkg")
	if err != nil {
		panic(err)
	}

	abiGenCmd.AddCommand(abigen())

	return abiGenCmd
}
