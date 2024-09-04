package user

import "github.com/decred/dcrd/dcrec/secp256k1/v4"

type User struct {
	PrivateKey *secp256k1.PrivateKey
	PublicKey  *secp256k1.PublicKey
}

func CreateUser() User {
	privateKey, publicKey := generateKeyPair(true)
	return User{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

func GetUser(privateKey *secp256k1.PrivateKey) User {
	return User{
		PrivateKey: privateKey,
		PublicKey:  privateKey.PubKey(),
	}
}

func GetUserFromFile(path string) User {
	privateKey := readPrivateKeyFromFile(path)
	return User{
		PrivateKey: privateKey,
		PublicKey:  privateKey.PubKey(),
	}
}

func (u *User) GetPrivateKey() *secp256k1.PrivateKey {
	return u.PrivateKey
}

func (u *User) GetPublicKey() *secp256k1.PublicKey {
	return u.PublicKey
}

//TODO: used for testing, remove later
func (u *User) PrintKeys() {
	printKeys(*u.PrivateKey, *u.PublicKey)
}

func (u *User) GetUserId() []byte {
	return u.PublicKey.SerializeCompressed()
}
