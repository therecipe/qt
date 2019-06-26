// +build sailfish

package deploy

import (
	"errors"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/crypto/ssh"

	"github.com/therecipe/qt/internal/utils"
)

func sailfish_ssh(port, login string, cmd ...string) error {

	typ := "SailfishOS_Emulator"
	if strings.HasPrefix(utils.QT_SAILFISH_VERSION(), "3.") {
		typ = "Sailfish_OS-Emulator-latest"
	}
	if port == "2222" {
		typ = "engine"
	}

	signer, err := ssh.ParsePrivateKey([]byte(utils.Load(filepath.Join(utils.SAILFISH_DIR(), "vmshare", "ssh", "private_keys", typ, login))))
	if err != nil {
		return err
	}

	client, err := ssh.Dial("tcp", "localhost:"+port, &ssh.ClientConfig{User: login, Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)}, HostKeyCallback: ssh.InsecureIgnoreHostKey()})
	if err != nil {
		return err
	}
	defer client.Close()

	sess, err := client.NewSession()
	if err != nil {
		return err
	}

	output, err := sess.CombinedOutput(strings.Join(cmd, " "))
	if err != nil {
		utils.Log.WithField("cmd", strings.Join(cmd, " ")).Debugf("failed to run ssh cmd for %v on %v", typ, runtime.GOOS)
		return errors.New(string(output))
	}

	return nil
}
