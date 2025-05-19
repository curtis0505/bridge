package pos

import "github.com/curtis0505/bridge/libs/util"

type ResponseBlockIncluded struct {
	HeaderBlockNumber string `json:"headerBlockNumber"`
	BlockNumber       string `json:"blockNumber"`
	Start             string `json:"start"`
	End               string `json:"end"`
	Proposer          string `json:"proposer"`
	Root              string `json:"root"`
	CreatedAt         string `json:"createdAt"`
	Message           string `json:"message"`
	Error             bool   `json:"error"`
}

func (resp ResponseBlockIncluded) RootBlockInfo() OutputRootBlockInfo {
	return OutputRootBlockInfo{
		Start: util.ToBigInt(resp.Start),
		End:   util.ToBigInt(resp.End),
	}
}

type ResponseFastMerkleProof struct {
	Proof string `json:"proof"`
}
