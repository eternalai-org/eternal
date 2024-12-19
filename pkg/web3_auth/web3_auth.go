package web3auth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"eternal/pkg/eth"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/go-jose/go-jose/v3"
	jose_jwt "github.com/golang-jwt/jwt/v5"
)

// IDToken represents the structure of the ID token
type IDToken struct {
	Iat               int          `json:"iat"`
	Aud               string       `json:"aud"`
	Iss               string       `json:"iss"`
	Email             string       `json:"email"`
	Name              string       `json:"name"`
	ProfileImage      string       `json:"profileImage"`
	Verifier          string       `json:"verifier"`
	VerifierID        string       `json:"verifierId"`
	AggregateVerifier string       `json:"aggregateVerifier"`
	Exp               int          `json:"exp"`
	Wallets           []UserWallet `json:"wallets"`
	Address           string       `json:"address"`
	SessionID         string       `json:"session_id"`
	IsViaWeb3Auth     bool         `json:"isViaWeb3Auth"`
	IsBVMUser         bool         `json:"isBVMUser"`
	TwitterUsername   string       `json:"twitterUsername"`
}

// UserWallet represents a user's wallet information
type UserWallet struct {
	PublicKey string `json:"public_key"`
	Type      string `json:"type"`
	Curve     string `json:"curve"`
}

// JWKSet represents a set of JSON Web Keys
type JWKSet struct {
	Keys []jose.JSONWebKey `json:"keys"`
}

// VerifyIDToken verifies the ID token and returns the IDToken object
func VerifyIDToken(idTokenStr, web3authClientID string) (*IDToken, error) {
	webKey, err := getJWKSet()
	if err != nil {
		return nil, err
	}

	token, err := jose_jwt.Parse(idTokenStr, func(token *jose_jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jose_jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return webKey.Key, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	var idToken IDToken
	if claims, ok := token.Claims.(jose_jwt.MapClaims); ok {
		if err = mapClaimsToIDToken(claims, &idToken); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid token claims")
	}

	if !strings.EqualFold(idToken.Aud, web3authClientID) {
		return nil, fmt.Errorf("invalid token audience %v", idToken.Aud)
	}

	userAddress, err := extractUserAddress(idToken.Wallets)
	if err != nil {
		return nil, err
	}

	idToken.Address = strings.ToLower(userAddress)
	idToken.IsViaWeb3Auth = true

	return &idToken, nil
}

// getJWKSet fetches and parses the JSON Web Key Set
func getJWKSet() (*jose.JSONWebKey, error) {
	url := "https://api-auth.web3auth.io/jwks"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var jwkSet JWKSet
	if err = json.Unmarshal(body, &jwkSet); err != nil {
		return nil, err
	}

	return &jwkSet.Keys[0], nil
}

// mapClaimsToIDToken maps JWT claims to IDToken
func mapClaimsToIDToken(claims jose_jwt.MapClaims, idToken *IDToken) error {
	b, err := json.MarshalIndent(claims, "", "  ")
	if err != nil {
		return err
	}

	return json.Unmarshal(b, idToken)
}

// extractUserAddress derives Ethereum address from user wallet data
func extractUserAddress(wallets []UserWallet) (string, error) {
	if len(wallets) == 0 {
		return "", errors.New("no wallets found")
	}
	userPubkey := wallets[0].PublicKey
	return eth.WalletAddressFromCompressedPublicKey(userPubkey)
}

// AuthToken represents the authentication token structure
type AuthToken struct {
	Address   string `json:"address"`
	Exp       int64  `json:"exp"`
	SessionID string `json:"session_id"`
}

// NewAuthToken creates a new AuthToken
func NewAuthToken(address, sessionID string, exp int64) *AuthToken {
	return &AuthToken{
		Address:   address,
		Exp:       exp,
		SessionID: sessionID,
	}
}

// EncryptAndSignAuthToken encrypts and signs the authentication token
func EncryptAndSignAuthToken(authToken AuthToken, serverKey string) (string, error) {
	encrypted, err := EncryptAuthToken(authToken, serverKey)
	if err != nil {
		return "", err
	}

	signed, err := SignAuthToken(encrypted, serverKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s", encrypted, signed), nil
}

// DecryptAndVerifyAuthToken decrypts and verifies the authentication token
func DecryptAndVerifyAuthToken(authToken, serverKey string) (*AuthToken, error) {
	encrypted, signed := SplitAuthToken(authToken)
	if signed == "" {
		return nil, errors.New("invalid token format")
	}

	masterWallet, err := crypto.HexToECDSA(serverKey)
	if err != nil {
		return nil, err
	}

	publicKey := masterWallet.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	serverAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	if err = VerifyAuthToken(encrypted, signed, serverAddress); err != nil {
		return nil, err
	}

	decrypted, err := DecryptAuthToken(encrypted, serverKey)
	if err != nil {
		return nil, err
	}

	if decrypted.Exp < time.Now().Unix() && decrypted.Exp != 0 {
		return nil, errors.New("token expired")
	}

	return decrypted, nil
}

// SplitAuthToken splits the auth token into encrypted and signed parts
func SplitAuthToken(authToken string) (string, string) {
	parts := strings.Split(authToken, ".")
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}

// EncryptAuthToken encrypts the authentication token
func EncryptAuthToken(authToken AuthToken, serverKey string) (string, error) {
	return Encrypt(authToken, serverKey)
}

// SignAuthToken signs the authentication token string
func SignAuthToken(authTokenStr, privateKey string) (string, error) {
	return Sign(authTokenStr, privateKey)
}

// DecryptAuthToken decrypts the authentication token string
func DecryptAuthToken(authTokenStr, serverKey string) (*AuthToken, error) {
	var authToken AuthToken
	if err := Decrypt(authTokenStr, serverKey, &authToken); err != nil {
		return nil, err
	}
	return &authToken, nil
}

// VerifyAuthToken verifies the authentication token signature
func VerifyAuthToken(authTokenStr, signature, address string) error {
	dataHash := crypto.Keccak256Hash([]byte(authTokenStr))
	return VerifySig(dataHash.Bytes(), signature, address)
}

// VerifySignature verifies the Ethereum signature against the address
func VerifySignature(fromAddress, signatureHex, message string) error {
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return err
	}

	signature[crypto.RecoveryIDOffset] -= 27 // Adjust recovery ID

	messageHash := accounts.TextHash([]byte(message))
	pubKey, err := crypto.SigToPub(messageHash, signature)
	if err != nil {
		return err
	}

	if common.HexToAddress(fromAddress) != crypto.PubkeyToAddress(*pubKey) {
		return errors.New("failed to verify signature")
	}

	return nil
}

// Encrypt encrypts data with the provided private key
func Encrypt(data interface{}, privKey string) (string, error) {
	masterWallet, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	publicKey := masterWallet.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	pubkey := ecies.ImportECDSAPublic(publicKeyECDSA)

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	ct, err := ecies.Encrypt(rand.Reader, pubkey, dataBytes, nil, nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ct), nil
}

// Decrypt decrypts the data with the provided private key
func Decrypt(dataStr, privKey string, result interface{}) error {
	masterWallet, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return err
	}

	encryptedBytes, err := base64.StdEncoding.DecodeString(dataStr)
	if err != nil {
		return err
	}

	decryptedBytes, err := ecies.ImportECDSA(masterWallet).Decrypt(encryptedBytes, nil, nil)
	if err != nil {
		return err
	}

	return json.Unmarshal(decryptedBytes, result)
}

// Sign signs the data with the provided private key
func Sign(dataStr, privKey string) (string, error) {
	masterWallet, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	dataHash := crypto.Keccak256Hash([]byte(dataStr))
	signature, err := crypto.Sign(dataHash.Bytes(), masterWallet)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifySig verifies the signature for the provided data hash and address
func VerifySig(dataHash []byte, sig64, address string) error {
	signature, err := base64.StdEncoding.DecodeString(sig64)
	if err != nil {
		return err
	}

	if signature[crypto.RecoveryIDOffset] == 27 || signature[crypto.RecoveryIDOffset] == 28 {
		signature[crypto.RecoveryIDOffset] -= 27
	}

	sigPub, err := crypto.SigToPub(dataHash, signature)
	if err != nil {
		return errors.New("invalid signature: " + err.Error())
	}

	sigAddress := crypto.PubkeyToAddress(*sigPub).Hex()
	if !strings.EqualFold(sigAddress, address) {
		return errors.New("invalid signature")
	}

	return nil
}
