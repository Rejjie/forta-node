package json_rpc_cache

import (
	"testing"
	"time"

	"github.com/forta-network/forta-core-go/protocol"
	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 500)

	cache.Append(blocks)

	blockNumber, ok := cache.Get(1, "eth_blockNumber", []byte("[]"))
	assert.True(t, ok)
	assert.Equal(t, "1", blockNumber)

	blockNumber, ok = cache.Get(2, "eth_blockNumber", []byte("[]"))
	assert.True(t, ok)
	assert.Equal(t, "101", blockNumber)

	block, ok := cache.Get(1, "eth_getBlockByNumber", []byte(`[ "1", true] `))
	assert.True(t, ok)
	assert.NotEmpty(t, block)

	logs, ok := cache.Get(1, "eth_getLogs", []byte(`[ { "toBlock":"1", "fromBlock":"1" } ]`))
	assert.True(t, ok)
	assert.NotEmpty(t, logs)

	trace, ok := cache.Get(1, "trace_block", []byte(`[ "1" ]`))
	assert.True(t, ok)
	assert.NotEmpty(t, trace)

	time.Sleep(time.Second)

	blockNumber, ok = cache.Get(1, "eth_blockNumber", []byte("[]"))
	assert.False(t, ok)
	assert.Empty(t, blockNumber)
}

var blocks = &protocol.BlocksData{
	Blocks: []*protocol.BlockData{
		{
			ChainID: 1,
			Block: &protocol.BlockWithTransactions{
				Hash:   "0xaaaa",
				Number: "1",
				Transactions: []*protocol.Transaction{
					{
						Hash: "0xbbbb",
						From: "0xcccc",
					},
				},
				Uncles: []string{"0xdddd"},
			},
			Logs: []*protocol.LogEntry{
				{
					Address: "0xcccc",
					Topics:  []string{"0xeeee"},
				},
			},
			Traces: []*protocol.Trace{
				{
					Action: &protocol.TraceAction{
						From: "0xcccc",
					},
					Result: &protocol.TraceResult{
						Address: "0xcccc",
					},
					TraceAddress: []int64{1},
				},
			},
		},
		{
			ChainID: 2,
			Block: &protocol.BlockWithTransactions{
				Hash:   "0xffff",
				Number: "100",
				Transactions: []*protocol.Transaction{
					{
						Hash: "0x1111",
						From: "0x2222",
					},
				},
				Uncles: []string{"0x3333"},
			},
			Logs: []*protocol.LogEntry{},
			Traces: []*protocol.Trace{
				{
					TraceAddress: []int64{2},
				},
			},
		},
		{
			ChainID: 2,
			Block: &protocol.BlockWithTransactions{
				Hash:   "0xfffd",
				Number: "101",
				Transactions: []*protocol.Transaction{
					{
						Hash: "0x1112",
						From: "0x2223",
					},
				},
				Uncles: []string{"0x3333"},
			},
			Logs: []*protocol.LogEntry{},
			Traces: []*protocol.Trace{
				{
					TraceAddress: []int64{1},
				},
			},
		},
	},
}
