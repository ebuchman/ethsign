package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/howeyc/gopass"
	"github.com/urfave/cli"
)

func cliSign(c *cli.Context) error {

	fromHex := c.String("from")
	keyDir := c.String("keydir")
	toHex := c.String("to")
	nonce := c.Int("nonce")
	amountFinney := int64(c.Int("amount"))
	gas := int64(c.Int("gas"))
	gasPriceGWei := int64(c.Int("price"))
	dataHex := c.String("data")

	to, err := hex.DecodeString(toHex)
	if err != nil {
		Exit(err)
	}

	from, err := hex.DecodeString(fromHex)
	if err != nil {
		Exit(err)
	}

	data, err := hex.DecodeString(dataHex)
	if err != nil {
		Exit(err)
	}

	toAddress := common.BytesToAddress(to)
	fromAddress := common.BytesToAddress(from)

	accountManager, err := makeAccountManager(keyDir)
	if err != nil {
		Exit(err)
	}

	fmt.Println("Please enter password")
	password, err := gopass.GetPasswd()
	if err != nil {
		return err
	}

	a := accounts.Account{Address: common.BytesToAddress(from)}
	d := time.Duration(60) * time.Second
	if err := accountManager.TimedUnlock(a, string(password), d); err != nil {
		return err
	}

	finneyFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil) // 10**15 = number of wei in a finney
	gweiFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)    // 10**9 = number of wei in a gwei
	amountBig := new(big.Int).Mul(big.NewInt(amountFinney), finneyFactor)
	gasPriceBig := new(big.Int).Mul(big.NewInt(gasPriceGWei), gweiFactor)
	gasBig := big.NewInt(gas)

	var tx *types.Transaction
	if to == nil {
		tx = types.NewContractCreation(uint64(nonce), amountBig, gasBig, gasPriceBig, data)
	} else {
		tx = types.NewTransaction(uint64(nonce), toAddress, amountBig, gasBig, gasPriceBig, data)
	}

	fmt.Println("TX: ", tx)

	chainID := big.NewInt(1)
	signer := types.NewEIP155Signer(chainID)

	signature, err := accountManager.SignEthereum(fromAddress, signer.Hash(tx).Bytes())
	if err != nil {
		Exit(err)
	}
	signedTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		Exit(err)
	}

	fmt.Printf("Signed TX: %v\n", signedTx)

	signedTxBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		Exit(err)
	}

	fmt.Printf("Signed TX RLP: %X\n", signedTxBytes)
	return nil
}

func Exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func makeAccountManager(keyDir string) (am *accounts.Manager, err error) {
	scryptN := accounts.StandardScryptN
	scryptP := accounts.StandardScryptP
	/*if conf.UseLightweightKDF {
		scryptN = accounts.LightScryptN
		scryptP = accounts.LightScryptP
	}*/
	return accounts.NewManager(keyDir, scryptN, scryptP), nil
}
