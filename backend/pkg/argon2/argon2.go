/*
sign up:
	INPUT: GET password
	FUNC_HASH:ENCODE password -> b64_hash:
		ENCODE password -> hash
		ENCODE hash -> b64_hash
		OUTPUT: b64_hash
	STORE b64_hash in DB

pseudo code:
	FUNC_HASH:ENCODE password -> b64_hash:
		arg = New(nil, nil)
		b64_hash = arg.Hash(password)
		OUTPUT: b64_hash

login:
	INPUT: GET password
	GET b64_hash from DB
	DECODE b64_hash -> hash
	ENCODE password -> hash
	COMPARE hash with hash in database
	OUTPUT: true/false
*/

package argon2

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	DefaultMemory      = 64 * 1024
	DefaultIterations  = 4
	DefaultParallelism = 4
	DefaultKeyLength   = 32
	DefaultSaltLength  = 16
)

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	KeyLength   uint32
}

type Argon2ID struct {
	Argon2Params
	Salt []byte
}

func New(params *Argon2Params, salt []byte) *Argon2ID {
	argon2ID := &Argon2ID{}
	if params == nil {
		argon2ID.fillDefaultParams()
	} else {
		argon2ID.Argon2Params = *params
		argon2ID.Salt = salt
	}
	return argon2ID
}

func (a *Argon2ID) fillDefaultParams() error {
	a.Memory = DefaultMemory
	a.Iterations = DefaultIterations
	a.Parallelism = DefaultParallelism
	a.KeyLength = DefaultKeyLength
	salt, err := getSalt(DefaultSaltLength)
	if err != nil {
		return err
	}
	a.Salt = salt
	return nil
}

func (a *Argon2ID) hash(password string) []byte {
	hash := argon2.IDKey([]byte(password), a.Salt, a.Iterations, a.Memory, a.Parallelism, a.KeyLength)
	return hash
}

func (a *Argon2ID) Hash(password string) ([]byte, error) {
	hash := a.hash(password)
	b64Hash := base64.StdEncoding.EncodeToString(hash)
	b64Salt := base64.StdEncoding.EncodeToString(a.Salt)
	hashString := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		a.Memory,
		a.Iterations,
		a.Parallelism,
		b64Salt,
		b64Hash,
	)
	return []byte(hashString), nil
}

func Verify(password string, encodedHash string) (bool, error) {
	params, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}
	a := New(params, salt)
	otherHash := a.hash(password)
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (*Argon2Params, []byte, []byte, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return nil, nil, nil, errors.New("invalid hash")
	}
	if parts[1] != "argon2id" {
		return nil, nil, nil, errors.New("unknown hash type")
	}

	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, errors.New("invalid hash")
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("invalid hash")
	}

	if parts[2] != fmt.Sprintf("v=%d", argon2.Version) {
		return nil, nil, nil, errors.New("invalid version")
	}
	params := &Argon2Params{}
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism)
	if err != nil {
		return nil, nil, nil, errors.New("invalid parameters")
	}

	salt, err := base64.StdEncoding.DecodeString(parts[4])
	if err != nil {
		return nil, nil, nil, errors.New("invalid salt")
	}
	hash, err := base64.StdEncoding.DecodeString(parts[5])
	if err != nil {
		return nil, nil, nil, errors.New("invalid hash")
	}

	params.KeyLength = uint32(len(hash))
	return params, salt, hash, nil
}

func getSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
