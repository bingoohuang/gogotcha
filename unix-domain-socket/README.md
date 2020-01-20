测试unix domain socket方式的http server, http client，以及基于ssh tunnel的unix socket连接方式。

检查selinux状态

```bash
$ bssh -H root12616,root12618 sestatus
Select Server :root12616,root12618
Run Command   :sestatus
root12618 ::  SELinux status:                 enabled
root12618 ::  SELinuxfs mount:                /sys/fs/selinux
root12618 ::  SELinux root directory:         /etc/selinux
root12618 ::  Loaded policy name:             targeted
root12618 ::  Current mode:                   permissive
root12618 ::  Mode from config file:          disabled
root12618 ::  Policy MLS status:              enabled
root12618 ::  Policy deny_unknown status:     allowed
root12618 ::  Max kernel policy version:      31
root12616 ::  SELinux status:                 enabled
root12616 ::  SELinuxfs mount:                /sys/fs/selinux
root12616 ::  SELinux root directory:         /etc/selinux
root12616 ::  Loaded policy name:             targeted
root12616 ::  Current mode:                   permissive
root12616 ::  Mode from config file:          disabled
root12616 ::  Policy MLS status:              enabled
root12616 ::  Policy deny_unknown status:     allowed
root12616 ::  Max kernel policy version:      31

$ bssh -H root12616,root12618 getenforce
Select Server :root12616,root12618
Run Command   :getenforce
root12618 ::  Permissive
root12616 ::  Permissive
```

Enabling/Disabling SELinux

    Open the file  /etc/selinux/config  and change the option SELINUX to disabled or vice versa.
    Restart the machine to take effect.


To change the mode of SELinux which is running

```bash
$ setenforce
usage:  setenforce [ Enforcing | Permissive | 1 | 0 ]
$ #To Set mode to Permissive
$ setenforce Permissive
```

1. `ssh -i /path/to/id_rsa user@server.nixcraft.com`
1. `curl -XGET --unix-socket /var/run/docker.sock http://unix/containers/json | cut -b 1-400`
1. `ssh -i ./rke_id_rsa rke@192.168.126.16 "curl -XGET --unix-socket /var/run/docker.sock http://unix/containers/json | cut -b 1-400"`

Build local
    
    go fmt ./...&&goimports -w .&&golint ./...&&golangci-lint run --enable-all&& go install  -ldflags="-s -w" ./..

Build for linux

    env GOOS=linux GOARCH=amd64 go install -ldflags="-s -w" ./...

Unix HTTP server

```bash
$ unix-domain-socket ./a.sock                 
Unix HTTP server
listened on unix ./a.sock
GET /hello HTTP/1.1
Host: unix
Accept-Encoding: gzip
User-Agent: Go-http-client/1.1


POST /hello HTTP/1.1
Host: unix
Accept-Encoding: gzip
Content-Length: 4
Content-Type: application/octet-stream
User-Agent: Go-http-client/1.1

baby
^CGot signal: interrupt
```

Unix HTTP client

```bash
$ unix-domain-socket ./a.sock /hello
Unix HTTP client
HTTP/1.1 200 OK
Content-Length: 6
Content-Type: text/plain; charset=utf-8
Date: Mon, 20 Jan 2020 08:59:19 GMT

Hello

$ unix-domain-socket ./a.sock /hello baby
Unix HTTP client
HTTP/1.1 200 OK
Content-Length: 6
Content-Type: text/plain; charset=utf-8
Date: Mon, 20 Jan 2020 08:59:26 GMT

Hello

$ 
```
