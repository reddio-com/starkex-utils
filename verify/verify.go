package verify

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func VerifyTransaction(stark_key, amount, nonce, asset_id, sender_vault_id, receiver_public_key, receiver_vault_id, expiration_timestamp, r, s string) (bool, error) {
	cmdStr := fmt.Sprintf("python3 verify.py --stark_key %s --amount %s --nonce %s --asset_id %s --sender_vault_id %s --receiver_public_key %s --receiver_vault_id %s --expiration_timestamp %s --r %s --s %s", stark_key, amount, nonce, asset_id, sender_vault_id, receiver_public_key, receiver_vault_id, expiration_timestamp, r, s)
	cmd := exec.Command("sh", "-c", cmdStr)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()

	if err != nil {
		return false, fmt.Errorf("%s %s", err, errb.String())
	}

	if strings.Trim(outb.String(), "\n") == "True" {
		return true, nil
	}

	if strings.Trim(outb.String(), "\n") == "False" {
		return false, nil
	}

	return false, fmt.Errorf("%s", outb.String())
}
