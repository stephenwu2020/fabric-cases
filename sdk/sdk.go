package sdk

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/pkg/errors"
)

var (
	wallet    *gateway.Wallet
	user      = "appuser"
	network   *gateway.Network
	contract  *gateway.Contract
	err       error
	channel   = "c1"
	chaincode = "mycc"
)

func Init() error {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")

	logging.SetLevel("fabsdk/core", logging.ERROR)

	//wallet, err = gateway.NewFileSystemWallet("wallet")
	wallet = gateway.NewInMemoryWallet()

	if err = populateWallet(wallet); err != nil {
		return errors.WithMessage(err, "Populate wallet faild")
	}

	ccpPath := filepath.Join(
		"..",
		"devnet",
		"organizations",
		"peerOrganizations",
		"org1.develop.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, user),
	)
	if err != nil {
		return errors.WithMessage(err, "Connect gateway failed")
	}

	network, err = gw.GetNetwork(channel)
	if err != nil {
		return errors.WithMessage(err, "Get network failed")
	}

	contract = network.GetContract(chaincode)
	return nil
}

func ChannelQuery(fcn string, args ...[]byte) ([]byte, error) {
	result, err := contract.EvaluateTransactionWithBytes(fcn, args...)
	return result, err
}

func ChannelExecute(fcn string, args ...[]byte) ([]byte, error) {
	result, err := contract.SubmitTransactionWithBytes(fcn, args...)
	return result, err
}

func populateWallet(wallet *gateway.Wallet) error {
	if wallet.Exists(user) {
		return nil
	}
	credPath := filepath.Join(
		"..",
		"devnet",
		"organizations",
		"peerOrganizations",
		"org1.develop.com",
		"users",
		"User1@org1.develop.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	err = wallet.Put(user, identity)
	if err != nil {
		return err
	}
	return nil
}
