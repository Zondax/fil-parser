package frc42

import (
	"encoding/binary"
	"fmt"
	"regexp"

	"golang.org/x/crypto/blake2b"
)

// https://github.com/filecoin-project/FIPs/blob/master/FRCs/frc-0042.md

// MethodNameError represents errors that can occur during method name validation or hashing
type MethodNameError struct {
	Message string
}

func (e MethodNameError) Error() string {
	return e.Message
}

// Predefined errors
var (
	ErrEmptyString       = MethodNameError{"empty method name provided"}
	ErrNotValidStart     = MethodNameError{"method name doesn't start with capital letter or _"}
	ErrIllegalCharacters = MethodNameError{"method name contains letters outside [a-zA-Z0-9_]"}
	ErrIndeterminableID  = MethodNameError{"unable to calculate method id, choose another method name"}
)

const (
	ConstructorMethodName   = "Constructor"
	ConstructorMethodNumber = uint64(1)
	FirstMethodNumber       = uint64(1 << 24) // 16777216
	DigestChunkLength       = 4
	DomainSeparationTag     = "1|"
)

// MethodHash generates a standard FRC-0042 compliant method number from a method name.
// This matches the behavior of the Rust `method_hash!` procedural macro.
//
// The method number is calculated as the first four bytes of `hash(method-name)`.
// The name "Constructor" is always hashed to 1 and other method names that hash to
// 0 or 1 are avoided via rejection sampling.
func MethodHash(methodName string) (uint64, error) {
	// Validate the method name
	if err := checkMethodName(methodName); err != nil {
		return 0, err
	}

	// Special case: Constructor always returns 1
	if methodName == ConstructorMethodName {
		return ConstructorMethodNumber, nil
	}

	// Prepend domain separator as per FRC-42 standard
	nameWithPrefix := DomainSeparationTag + methodName

	// Hash using Blake2b with 64-byte output (512 bits)
	hasher, err := blake2b.New512(nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create Blake2b hasher: %w", err)
	}

	hasher.Write([]byte(nameWithPrefix))
	digest := hasher.Sum(nil)

	// Process digest in 4-byte chunks
	for i := 0; i < len(digest); i += DigestChunkLength {
		if i+DigestChunkLength > len(digest) {
			// Last chunk may be smaller than 4 bytes
			break
		}

		chunk := digest[i : i+DigestChunkLength]
		methodID := binary.BigEndian.Uint32(chunk)

		// Method numbers below FirstMethodNumber are reserved
		if uint64(methodID) >= FirstMethodNumber {
			return uint64(methodID), nil
		}
	}

	return 0, ErrIndeterminableID
}

// checkMethodName validates that a method name is compliant with the FRC-0042 standard:
// - Only ASCII characters in [a-zA-Z0-9_] are allowed
// - Starts with a character in [A-Z_]
func checkMethodName(methodName string) error {
	if methodName == "" {
		return ErrEmptyString
	}

	// Check first character
	firstChar := rune(methodName[0])
	if !((firstChar >= 'A' && firstChar <= 'Z') || firstChar == '_') {
		return ErrNotValidStart
	}

	// Check all characters are valid
	validChars := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !validChars.MatchString(methodName) {
		return ErrIllegalCharacters
	}

	return nil
}
