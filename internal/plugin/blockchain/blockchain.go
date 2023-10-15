package blockchain

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

var cl *ethclient.Client

func NewEthClieng(url string) {
	var err error
	cl, err = ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
}

//type GoEthClient struct {
//	Id       int64
//	Name     string
//	Client   *ethclient.Client
//	Abi      abi.abi
//	Amount   *big.Int
//	GasLimit uint64
//	ChainID  *big.Int
//}
//
//// ClientConfig config for get go-ethereum client
//type ClientConfig struct {
//	Id              int64
//	Name            string
//	Rpc             string
//	ContractAddress string
//	PrivateKey      string
//	WalletAddress   string
//	AbiPath         string
//	GasLimit        int64
//}

//var lock sync.Mutex
//func GetClient(id int64) (GoEthClient, error) {
//	client, ok :=
//}

// getGoEthClient get go-ethereum client
//func getGoEthClient(clientConfig ClientConfig) (GoEthClient, error) {
//	client, err := ethclient.Dial(clientConfig.Rpc)
//	if err != nil {
//		log.Println("ethclient.Dial error: ", err)
//		return GoEthClient{}, err
//	}
//
//	// open abi file and parse json
//	open, err := os.Open(clientConfig.AbiPath)
//	if err != nil {
//		log.Println("open abi file error: ", err)
//		return GoEthClient{}, err
//	}
//	contractAbi, err := abi.JSON(open)
//	if err != nil {
//		log.Println("abi.JSON error: ", err)
//		return GoEthClient{}, err
//	}
//
//	// transfer amount, if no set zero
//	amount := big.NewInt(0)
//	// gas limit
//	gasLimit := uint64(clientConfig.GasLimit)
//	// get chain id
//	chainID, err := client.NetworkID(context.Background())
//	if err != nil {
//		log.Println("client.NetworkID error: ", err)
//		return GoEthClient{}, err
//	}
//	// generate goEthClient
//	goEthClient := GoEthClient{
//		Id:       clientConfig.Id,
//		Name:     clientConfig.Name,
//		Client:   client,
//		Abi:      contractAbi,
//		Amount:   amount,
//		GasLimit: gasLimit,
//		ChainID:  chainID,
//	}
//	return goEthClient, nil
//}

//func VerifyContract(result, proof string) error {
//	callData, err := abi.ABI{}.Pack("")
//	msg := ethereum.CallMsg{
//		To:   "",
//		Data: callData,
//	}
//	cl.CallContract(context.Background(), msg, nil)
//
//}
