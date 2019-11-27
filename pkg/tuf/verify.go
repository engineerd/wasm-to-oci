package tuf

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/theupdateframework/notary/client"
)

// VerifyFileTrust ensures the trust metadata for a given GUN matches the computed metadata of the local file
func VerifyFileTrust(ref, localFile, trustServer, tlscacert, trustDir, timeout string) error {
	target, trustedSHA, err := GetTargetAndSHA(ref, trustServer, tlscacert, trustDir, timeout)
	if err != nil {
		return err
	}
	log.Infof("Pulled trust data for %v, with role %v - SHA256: %v", ref, target.Role, trustedSHA)

	buf, err := ioutil.ReadFile(localFile)
	if err != nil {
		return err
	}

	err = verifyTargetSHAFromBytes(target, buf)
	if err == nil {
		log.Infof("The SHA sums are equal: %v\n", trustedSHA)
	}

	return err
}

func verifyTargetSHAFromBytes(target *client.TargetWithRole, buf []byte) error {
	trustedSHA := hex.EncodeToString(target.Hashes["sha256"])
	hasher := sha256.New()
	hasher.Write(buf)
	computedSHA := hex.EncodeToString(hasher.Sum(nil))

	log.Infof("Computed SHA: %v\n", computedSHA)
	if trustedSHA != computedSHA {
		return fmt.Errorf("the digest sum of the artifact from the trusted collection %v is not equal to the computed digest %v",
			trustedSHA, computedSHA)
	}
	return nil
}

// GetTargetAndSHA returns the target with roles and the SHA256 of the target file
func GetTargetAndSHA(ref, trustServer, tlscacert, trustDir, timeout string) (*client.TargetWithRole, string, error) {
	gun, name := SplitTargetRef(ref)
	target, err := GetTargetWithRole(gun, name, trustServer, tlscacert, trustDir, timeout)
	if err != nil {
		return nil, "", err
	}

	return target, hex.EncodeToString(target.Hashes["sha256"]), nil
}
