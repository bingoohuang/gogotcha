package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/goutils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func main() {
	if len(os.Args) < 2 { // nolint gomnd
		fmt.Fprintln(os.Stderr, "usage:", os.Args[0], "ip")
		fmt.Fprintln(os.Stderr, "unix http server usage:", os.Args[0], "/path.sock")
		fmt.Fprintln(os.Stderr, "unix http client usage:", os.Args[0], "/path.sock /uri [data]")

		return
	}

	arg1 := os.Args[1] // nolint gomnd
	p := arg1[0:1]     // nolint gomnd

	if p == "/" || p == "." || p == "~" { // tell if it is an unix sock file name
		if len(os.Args) == 2 { // nolint gomnd
			unitHTTPServer(arg1)
			return
		}

		uri := os.Args[2] // nolint gomnd
		postData := ""

		if len(os.Args) >= 4 { // nolint gomnd
			postData = os.Args[3] // nolint gomnd
		}

		unitHTTPClient(arg1, uri, postData)

		return
	}

	testDialOnSSHTunnel(arg1)
}

func testDialOnSSHTunnel(arg1 string) {
	sshKeyString, err := privateKeyPath("./rke_id_rsa")
	if err != nil {
		panic(err)
	}

	d := dialer{
		netConn:      "unix",
		username:     "rke",
		dockerSocket: "/var/run/docker.sock",
		sshAddress:   arg1 + ":22",
		sshKeyString: sshKeyString,
	}
	tr := &http.Transport{
		Dial: func(network, address string) (net.Conn, error) {
			return d.Dial("unix", "")
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://unix/containers/json")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	dump, _ := httputil.DumpResponse(resp, true)
	abbr, _ := goutils.Abbreviate(string(dump), 900)
	fmt.Println(abbr)
}

type dialer struct {
	sshKeyString    string
	sshCertString   string
	sshAddress      string
	username        string
	netConn         string
	dockerSocket    string
	useSSHAgentAuth bool
}

func (d *dialer) Dial(network, addr string) (net.Conn, error) {
	conn, err := d.getSSHTunnelConnection()
	if err != nil {
		return d.checkSSHTunelError(err)
	}

	// Docker Socket....
	if d.netConn == "unix" {
		addr = d.dockerSocket
		network = d.netConn
	}

	remote, err := conn.Dial(network, addr)
	if err != nil {
		errStr := err.Error()

		fmt.Printf("network:%s, addr:%s, error %+v\n", network, addr, err)

		if strings.Contains(errStr, "connect failed") {
			return nil, fmt.Errorf("unable to access the service on %s. "+
				"The service might be still starting up. Error: %v", addr, err)
		}

		if strings.Contains(errStr, "administratively prohibited") {
			return nil, fmt.Errorf("unable to access the Docker socket (%s). "+
				"Please check if the configured user can execute `docker ps` on the node, "+
				"and if the SSH server version is at least version 6.7 or higher. "+
				"If you are using RedHat/CentOS, you can't use the user `root`. "+
				"Please refer to the documentation for more instructions. Error: %v", addr, err)
		}

		return nil, fmt.Errorf("failed to dial to %s: %+v", addr, err)
	}

	return remote, err
}

func (d *dialer) checkSSHTunelError(err error) (net.Conn, error) {
	errStr := err.Error()

	if strings.Contains(errStr, "no key found") {
		return nil, fmt.Errorf("unable to access node with address [%s] using SSH. "+
			"Please check if the configured key or specified key file is a valid SSH Private Key."+
			" Error: %v", d.sshAddress, err)
	}

	if strings.Contains(errStr, "no supported methods remain") {
		return nil, fmt.Errorf("unable to access node with address [%s] using SSH. "+
			"Please check if you are able to SSH to the node using the specified SSH Private Key "+
			"and if you have configured the correct SSH username. Error: %v", d.sshAddress, err)
	}

	if strings.Contains(errStr, "cannot decode encrypted private keys") {
		return nil, fmt.Errorf("unable to access node with address [%s] using SSH. "+
			"Using encrypted private keys is only supported using ssh-agent. "+
			"Please configure the option `ssh_agent_auth: true` in the configuration file "+
			"or use --ssh-agent-auth as a parameter when running RKE. "+
			"This will use the `SSH_AUTH_SOCK` environment variable. Error: %v", d.sshAddress, err)
	}

	if strings.Contains(errStr, "operation timed out") {
		return nil, fmt.Errorf("unable to access node with address [%s] using SSH. "+
			"Please check if the node is up and is accepting SSH connections "+
			"or check network policies and firewall rules. Error: %v", d.sshAddress, err)
	}

	return nil, fmt.Errorf("failed to dial ssh using address [%s]: %v", d.sshAddress, err)
}

func (d *dialer) getSSHTunnelConnection() (*ssh.Client, error) {
	cfg, err := getSSHConfig(d.username, d.sshKeyString, d.sshCertString, d.useSSHAgentAuth)
	if err != nil {
		return nil, fmt.Errorf("error configuring SSH: %v", err)
	}

	// Establish connection with SSH server
	fmt.Printf("tcp %s with config:%+v\n", d.sshAddress, cfg)

	return ssh.Dial("tcp", d.sshAddress, cfg)
}

// nolint gosec
func getSSHConfig(username, sshPrivateKey, sshCertificate string, useAgentAuth bool) (*ssh.ClientConfig, error) {
	config := &ssh.ClientConfig{User: username, HostKeyCallback: ssh.InsecureIgnoreHostKey()}

	// Kind of a double check now.
	if useAgentAuth {
		if sshAgentSock := os.Getenv("SSH_AUTH_SOCK"); sshAgentSock != "" {
			sshAgent, err := net.Dial("unix", sshAgentSock)
			if err != nil {
				return config, fmt.Errorf("cannot connect to SSH Auth socket %q: %s", sshAgentSock, err)
			}

			config.Auth = append(config.Auth, ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers))

			logrus.Debugf("using %q SSH_AUTH_SOCK", sshAgentSock)

			return config, nil
		}
	}

	signer, err := parsePrivateKey(sshPrivateKey)
	if err != nil {
		return config, err
	}

	if len(sshCertificate) > 0 {
		key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(sshCertificate))
		if err != nil {
			return config, fmt.Errorf("unable to parse SSH certificate: %v", err)
		}

		if _, ok := key.(*ssh.Certificate); !ok {
			return config, fmt.Errorf("unable to cast public key to SSH Certificate")
		}

		signer, err = ssh.NewCertSigner(key.(*ssh.Certificate), signer)
		if err != nil {
			return config, err
		}
	}

	config.Auth = append(config.Auth, ssh.PublicKeys(signer))

	return config, nil
}

func parsePrivateKey(keyBuff string) (ssh.Signer, error) {
	return ssh.ParsePrivateKey([]byte(keyBuff))
}

func privateKeyPath(sshKeyPath string) (string, error) {
	if sshKeyPath[:2] == "~/" {
		sshKeyPath = filepath.Join(userHome(), sshKeyPath[2:])
	}

	buff, err := ioutil.ReadFile(sshKeyPath)
	if err != nil {
		return "", fmt.Errorf("error while reading SSH key file: %v", err)
	}

	return string(buff), nil
}

func userHome() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}

	homeDrive := os.Getenv("HOMEDRIVE")
	homePath := os.Getenv("HOMEPATH")

	if homeDrive != "" && homePath != "" {
		return homeDrive + homePath
	}

	return os.Getenv("USERPROFILE")
}
