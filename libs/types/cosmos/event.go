package cosmos

import (
	"encoding/json"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	"strconv"
)

type EventWithdrawReward struct {
	Validator string `json:"validator"`
	Amount    string `json:"amount"`
}

func (event EventWithdrawReward) EventName() string {
	return distributiontypes.EventTypeWithdrawRewards
}

func (event EventWithdrawReward) Coins() (cosmossdk.Coins, error) {
	return cosmossdk.ParseCoinsNormalized(event.Amount)
}

type EventIBCTransferSendPacket struct {
	ChannelOrdering    string `json:"packet_channel_ordering"`
	Connection         string `json:"packet_connection"`
	Sequence           string `json:"packet_sequence,omitempty"`
	SourcePort         string `json:"packet_src_port,omitempty"`
	SourceChannel      string `json:"packet_src_channel,omitempty"`
	DestinationPort    string `json:"packet_dst_port,omitempty"`
	DestinationChannel string `json:"packet_dst_channel,omitempty"`
	Data               string `json:"packet_data,omitempty"`
	DataHex            string `json:"packet_data_hex"`
	TimeoutHeight      string `json:"packet_timeout_height"`
	TimeoutTimestamp   string `json:"packet_timeout_timestamp,omitempty"`
}

func (event EventIBCTransferSendPacket) GetSequence() uint64 {
	s, err := strconv.Atoi(event.Sequence)
	if err != nil {
		return 0
	}
	return uint64(s)
}

func (event *EventIBCTransferSendPacket) Packet() ibcchanneltypes.Packet {
	return ibcchanneltypes.Packet{
		Sequence:           event.GetSequence(),
		SourcePort:         event.SourcePort,
		SourceChannel:      event.SourceChannel,
		DestinationPort:    event.DestinationPort,
		DestinationChannel: event.DestinationChannel,
		Data:               []byte(event.Data),
	}
}

func (event EventIBCTransferSendPacket) FungibleTokenPacketData() ibctransfertypes.FungibleTokenPacketData {
	var packet ibctransfertypes.FungibleTokenPacketData
	json.Unmarshal([]byte(event.Data), &packet)
	return packet
}
