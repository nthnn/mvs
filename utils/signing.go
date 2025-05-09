package utils

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nthnn/mvs/core"
)

func EnsureKeys() error {
	if _, err := os.Stat(
		core.PrivateKeyPem,
	); err == nil {
		return nil
	}

	if err := os.MkdirAll(
		core.KeyDirectory,
		0700,
	); err != nil {
		return fmt.Errorf("could not create key dir: %w", err)
	}

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("key generation failed: %w", err)
	}

	privBlock := &pem.Block{
		Type:  core.GlobalConfiguration.PrivateKey,
		Bytes: priv.Seed(),
	}

	if err := AtomicWriteFile(
		core.PrivateKeyPem,
		pem.EncodeToMemory(privBlock),
		0600,
	); err != nil {
		return fmt.Errorf("writing private key: %w", err)
	}

	pubBlock := &pem.Block{
		Type:  core.GlobalConfiguration.PublicKey,
		Bytes: pub,
	}

	if err := AtomicWriteFile(
		core.PublicKeyPem,
		pem.EncodeToMemory(pubBlock),
		0644,
	); err != nil {
		return fmt.Errorf("writing public key: %w", err)
	}

	return nil
}

func Sign(path string) error {
	if err := EnsureKeys(); err != nil {
		return err
	}

	data, err := os.ReadFile(core.PrivateKeyPem)
	if err != nil {
		return fmt.Errorf("reading private key: %w", err)
	}

	block, _ := pem.Decode(data)
	if block.Type != core.GlobalConfiguration.PrivateKey ||
		block == nil {
		return fmt.Errorf("invalid private key PEM")
	}

	seed := block.Bytes
	priv := ed25519.NewKeyFromSeed(seed)

	plain, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}

	sig := ed25519.Sign(priv, plain)
	return AtomicWriteFile(path+".sig", sig, 0644)
}

func Verify(path string) error {
	data, err := os.ReadFile(core.PublicKeyPem)
	if err != nil {
		return fmt.Errorf("reading public key: %w", err)
	}

	block, _ := pem.Decode(data)
	if block.Type != core.GlobalConfiguration.PublicKey ||
		block == nil {
		return fmt.Errorf("invalid public key PEM")
	}

	pub := ed25519.PublicKey(block.Bytes)
	plain, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}

	sig, err := os.ReadFile(path + ".sig")
	if err != nil {
		return fmt.Errorf("reading signature: %w", err)
	}

	if !ed25519.Verify(pub, plain, sig) {
		return fmt.Errorf(
			"signature verification failed for %s",
			filepath.Base(path),
		)
	}

	return nil
}
