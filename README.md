[![Go Report](https://goreportcard.com/badge/github.com/zyfdegh/remote-docker-exec)](https://goreportcard.com/report/github.com/zyfdegh/remote-docker-exec)

# remote-docker-exec
Connect to docker daemon or swarm with TLS. Run command 'sh' in container.

Screenshot:

![Mdviewer logo](https://raw.githubusercontent.com/zyfdegh/remote-docker-exec/master/raw/screenshot-01.png)

# Precondition
To run this programme, a TLS enabled docker daemon is required with port binded.

TCP port is not enable in docker daemon by default, click [here][1] to see how to bind a port.

To enable TLS(SSL), you need generate cert files by yourself with **openssl**, follow this [article][2] for detail. After that, restart docker daemon with --tls flags to enable TLS, follow docker [docs][3].

# Params

**PublicIP:** public IP of docker daemon or swarm

**DockerDaemonPort:** listening port of docker daemon or swarm

**ContainerId:** name or ID of container. For docker daemon, support containers created by daemon. For swarm, support containers created by all daemons in cluster.

# Version
Support docker 1.11.x with API version 1.23

Note that docker 1.12 has merged swarm into daemon, not test so far.

# LICENSE
MIT

[1]:https://docs.docker.com/engine/reference/commandline/dockerd/#bind-docker-to-another-host-port-or-a-unix-socket
[2]:https://jamielinux.com/docs/openssl-certificate-authority/
[3]:https://docs.docker.com/engine/security/https/
