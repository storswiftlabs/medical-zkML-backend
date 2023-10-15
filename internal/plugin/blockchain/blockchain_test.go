package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"testing"
)

func init() {
	NewEthClieng("https://rpc.ankr.com/filecoin_testnet")
}

var (
	SK   = "57748462d0245d088af13be5e3e97d643bd28b0a8a360ed059f8dbe5a5671018"
	ADDR = "0x4fda4174D5D07C906395bfB77806287cc65Fd129"
)

func TestBlockChain(t *testing.T) {

	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(chainid)

	//var (
	//	sk       = crypto.ToECDSAUnsafe(common.FromHex(SK))
	//	to       = common.HexToAddress("0x26AC9733FF86d28Bb6076E372482d89f7299bFe5")
	//	value    = new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether))
	//	sender   = common.HexToAddress(ADDR)
	//	gasLimit = uint64(800000)
	//)
	//
	//nonce, err := cl.PendingNonceAt(context.Background(), sender)
	//if err != nil {
	//	t.Fatalf(err.Error())
	//}
	//
	//// Get suggested gas price
	//tipCap, _ := cl.SuggestGasTipCap(context.Background())
	//feeCap, _ := cl.SuggestGasPrice(context.Background())
	//// Create a new transaction
	//tx := types.NewTx(
	//	&types.DynamicFeeTx{
	//		ChainID:   chainid,
	//		Nonce:     nonce,
	//		GasTipCap: tipCap,
	//		GasFeeCap: feeCap,
	//		Gas:       gasLimit,
	//		To:        &to,
	//		Value:     value,
	//		Data:      nil,
	//	})
	//// Sign the transaction using our keys
	//signedTx, _ := types.SignTx(tx, types.NewLondonSigner(chainid), sk)
	//// Send the transaction to our node
	//if err := cl.SendTransaction(context.Background(), signedTx); err != nil {
	//	t.Fatalf(err.Error())
	//}

	open, err := os.Open("../abi/test3.json")
	if err != nil {
		t.Fatalf(err.Error())
	}

	contractAbi, err := abi.JSON(open)
	if err != nil {
		t.Fatalf(err.Error())
	}

	callData, err := contractAbi.Pack("GetHelloWorld")
	if err != nil {
		t.Fatalf(err.Error())
	}

	contractAddr := common.HexToAddress("0x26AC9733FF86d28Bb6076E372482d89f7299bFe5")
	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: callData,
	}

	result, err := cl.CallContract(context.Background(), msg, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(result) == 0 {
		t.Logf("error: %s", err)

	}
	var temp string
	if err := contractAbi.UnpackIntoInterface(&temp, "GetHelloWorld", result); err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(temp)
}
