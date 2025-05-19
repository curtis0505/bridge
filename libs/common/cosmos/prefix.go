package cosmos

const (
	chainATOM  = "ATOM"
	chainKAVA  = "KAVA"
	chainFNSA  = "FNSA"
	chainTFNSA = "TFNSA"
	chainOSMO  = "OSMO"
)

var addressPrefix = map[string]string{
	chainATOM:  "cosmos",
	chainKAVA:  "kava",
	chainOSMO:  "osmo",
	chainFNSA:  "link",
	chainTFNSA: "tlink",
}

func GetAddressPrefixByChain(chain string) string {
	return addressPrefix[chain]
}
