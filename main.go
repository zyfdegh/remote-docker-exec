// Author: Zhang Yifa
// Email: yzhang3@linkernetworks.com

package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/fsouza/go-dockerclient"
	"github.com/spf13/cobra"
)

func main() {
	//call console
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("execute command error: %v", err)
		os.Exit(-1)
	}
}

// RootCmd is root command of remote-docker-exec
var rootCmd = &cobra.Command{
	Use:     "remote-docker-exec [PublicIP] [DockerDaemonPort] [ContainerId]",
	Short:   "remote-docker-exec is a WebSocket with SSL shell, it can connect to docker container, and act as a remote docker exec.",
	Long:    "remote-docker-exec is a WebSocket with SSL shell, it can connect to docker container, and act as a remote docker exec. For linker internal use only.",
	Example: "./remote-docker-exec 52.78.72.139 2376 ffcede5a47cb",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "linux" {
			log.Fatalln("Only linux is supported now.")
			return
		}

		// example
		// remote-docker-exec 52.77.230.38 2376 c3598346f0d7
		if cmd.Flags().NArg() == 3 {
			publicIP := cmd.Flags().Args()[0]
			dockerDaemonPort := cmd.Flags().Args()[1]
			containerId := cmd.Flags().Args()[2]
			remoteDockerExec(publicIP, dockerDaemonPort, containerId)
		}

		// others
		cmd.SetArgs([]string{"--help"})
		if err := cmd.Execute(); err != nil {
			log.Fatalf("command arguments error: %v", err)
			os.Exit(-1)
		}
		return
	},
}

func remoteDockerExec(ip, port, containerId string) {
	fmt.Println("Welcome to linker web console!")

	endpoint := fmt.Sprintf("tcp://%s:%s", ip, port)
	fmt.Printf("Connecting to %s, please wait...\n", endpoint)

	ca, cert, key := "./ca.pem", "./cert.pem", "./key.pem"
	client, err := docker.NewTLSClient(endpoint, cert, key, ca)
	if err != nil {
		log.Fatalf("new tls client error: %v\n", err)
	}

	fmt.Printf("Connecting to container %s, please wait...\n", containerId)

	// select shell order
	shells := "bash || sh || csh || zsh"

	// create exec
	createOpts := docker.CreateExecOptions{}
	createOpts.AttachStdin = true
	createOpts.AttachStdout = true
	createOpts.AttachStderr = true
	createOpts.Tty = true
	createOpts.Cmd = []string{shells}
	createOpts.Container = containerId

	exec, err := client.CreateExec(createOpts)
	if err != nil {
		log.Fatalf("create exec error: %v\n", err)
	}

	// start exec
	startOpts := docker.StartExecOptions{}
	startOpts.Tty = true
	startOpts.RawTerminal = true
	startOpts.Detach = false
	startOpts.ErrorStream = os.Stderr
	startOpts.InputStream = os.Stdin
	startOpts.OutputStream = os.Stdout

	err = client.StartExec(exec.ID, startOpts)
	if err != nil {
		log.Fatalf("start exec error: %v\n", err)
	}
}
