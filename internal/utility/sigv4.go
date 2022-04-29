package utility

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	config "github.com/aws/aws-sdk-go-v2/config"
	"io/ioutil"
	"net/http"
	"time"
)

func SigV4SignRequest(r *http.Request, awsProfile, region, sigV4Service string) error {
	signer := v4.Signer{}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("unable to read body of connection %w", err)
	}
	r.Body = ioutil.NopCloser(bytes.NewReader(payload))

	payloadBytes := sha256.Sum256([]byte(payload))
	hash := hex.EncodeToString(payloadBytes[:])

	var cfg aws.Config
	switch awsProfile {
	case "":
		cfg, err = config.LoadDefaultConfig(context.Background())
	default:
		cfg, err = config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(awsProfile))
	}
	if err != nil {
		return fmt.Errorf("failed to generate aws sigv4 signed connection %w", err)
	}

	credentials, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		return fmt.Errorf("failed to retreive aws credentials while generating aws sigv4 signed connection %w", err)
	}

	err = signer.SignHTTP(context.Background(), credentials, r, hash, sigV4Service, region, time.Now())
	if err != nil {
		return fmt.Errorf("failed to sign aws sigv4 connection %w", err)
	}

	return nil
}
