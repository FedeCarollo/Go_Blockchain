package blockchain

import (
	"fmt"
	"os"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func generateKeyPair(save bool, path string) (*secp256k1.PrivateKey, *secp256k1.PublicKey) {
	privateKey, err := secp256k1.GeneratePrivateKey()

	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PubKey()

	if save {
		saveKeyToFile(path, privateKey.Serialize())
		// saveKeyToFile("public.pem", publicKey.SerializeUncompressed())
	}
	return privateKey, publicKey

}

func printKeys(privateKey secp256k1.PrivateKey, publicKey secp256k1.PublicKey) {
	fmt.Printf("Private key: %x\n", privateKey.Serialize())
	fmt.Printf("Public key (compressed): %x\n", publicKey.SerializeCompressed())
	fmt.Printf("Public key: %x\n", publicKey.SerializeUncompressed())
}

func saveKeyToFile(path string, key []byte) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(key)

	if err != nil {
		panic(err)
	}
}

func readPrivateKeyFromFile(path string) *secp256k1.PrivateKey {
	file, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	key := make([]byte, 32)
	_, err = file.Read(key)

	if err != nil {
		panic(err)
	}
	privateKey := secp256k1.PrivKeyFromBytes(key)
	return privateKey
}

func readPublicKeyFromFile(path string) *secp256k1.PublicKey {
	file, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	key := make([]byte, 33)
	_, err = file.Read(key)

	if err != nil {
		panic(err)
	}

	publicKey, err := secp256k1.ParsePubKey(key)
	if err != nil {
		panic(err)
	}
	return publicKey
}
