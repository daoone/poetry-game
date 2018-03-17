package main

import (
	"fmt"
	"log"
	"math/big"
	//"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key) // 新账户
	fmt.Println("Base account: " + auth.From.Hex())

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	sim := backends.NewSimulatedBackend(alloc) // 虚拟节点

	// 部署此时transaction pending中
	addr, _, contract, err := DeployPoetry(auth, sim, big.NewInt(100000000000), uint8(2))
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
		panic(err)
	}

	// 测试环境下的提交挖矿
	fmt.Println("Mining...")
	sim.Commit()

	// 我们熟悉的地址
	fmt.Println("Poetry address: " + addr.Hex()) // Poetry合约部署地址
	xmbAddr, _ := contract.Xmb(nil)
	fmt.Println("XmbToken address: " + xmbAddr.Hex()) // 新建的XmbToken合约地址

	fmt.Println("Adding a poem...")
	// 参与写诗（发起一笔交易）
	contract.AddPoem(&bind.TransactOpts{
		From: 		auth.From,
		Signer: 	auth.Signer,
		GasLimit: 	uint64(3000000),
		Value: 		big.NewInt(0),
	}, "I love Xiaoming!")

	fmt.Println("Mining...")
	sim.Commit()

	// 看看交易结果
	poems, _ := contract.Poems(nil, big.NewInt(0))
	fmt.Println("Poem content: " + poems.Content,
			"Poem voted token: " + poems.Votes.String(),
			"Poem voter counts: " + poems.VoteCounts.String(),
			"Poet address: " + poems.PoetAddr.Hex())
}

