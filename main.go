package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config represents the structure of the configuration file
type Config map[string]SSHConfig

// SSHConfig represents the structure of SSH configuration for a particular environment
type SSHConfig struct {
	Username string       `yaml:"username"`
	IP       string       `yaml:"ip"`
	Services []SSHService `yaml:"services"`
}

// SSHService represents the structure of a service to be port forwarded
type SSHService struct {
	Name      string `yaml:"name"`
	Addr      string `yaml:"addr"`
	LocalPort int    `yaml:"local_port,omitempty"`
}

func main() {
	config, err := readConfig()
	if err != nil {
		fmt.Println("Error reading configuration:", err)
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		fmt.Println("Usage: sshpf <environment>")
		os.Exit(1)
	}

	envKey := "default"
	if len(os.Args) == 2 {
		envKey = os.Args[1]
	}

	targetConfig, ok := config[envKey]
	if !ok && envKey != "command" {
		fmt.Printf("Environment %s not found in configuration\n", envKey)
		os.Exit(1)
	}

	fmt.Print(generateSSHString(targetConfig))
}

func readConfig() (Config, error) {
	var config Config

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}

	configFilePath := filepath.Join(homeDir, "bin", "sshpf_config.yaml")
	file, err := os.Open(configFilePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func generateSSHString(config SSHConfig) string {
	var sshString strings.Builder

	sshString.WriteString(fmt.Sprintf("ssh %s@%s", config.Username, config.IP))

	for _, service := range config.Services {
		if service.LocalPort != 0 {
			sshString.WriteString(fmt.Sprintf(" -L %d:%s", service.LocalPort, service.Addr))
		} else {
			sshString.WriteString(fmt.Sprintf(" -L %s:%s", extractRemotePort(service.Addr), service.Addr))
		}
	}

	return strings.TrimSuffix(sshString.String(), ":")
}

func extractRemotePort(addr string) string {
	parts := strings.Split(addr, ":")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}
