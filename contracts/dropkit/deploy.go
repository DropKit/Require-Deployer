package dropkitContract

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func Deploy() (string, string) {
	quorumEndpoint := viper.GetString(`QUORUM.ENDPOINT`)
	privatekeyHex := viper.GetString(`ACCOUNT.PRIVATEKEY`)

	var address common.Address
	var transaction *types.Transaction

	quorumClient, err := ethclient.Dial(quorumEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privatekeyHex)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal(err)
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := quorumClient.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := quorumClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(5000000)
	auth.GasPrice = gasPrice

	address, transaction, _, err = DeployDropkitContract(auth, quorumClient, accountAddress, accountAddress)
	if err != nil {
		log.Fatal(err)
	}
	return address.Hex(), transaction.Hash().Hex()
}
