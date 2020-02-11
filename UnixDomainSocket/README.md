# unix domain http server/client

测试unix domain socket方式的http server, http client，以及基于ssh tunnel的unix socket连接方式。

## 检查selinux状态

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

## Enabling/Disabling SELinux

1. Open the file  /etc/selinux/config  and change the option SELINUX to disabled or vice versa.
1. Restart the machine to take effect.

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

## Checking log

```bash
# ls -lh /var/run
lrwxrwxrwx. 1 root root 6 Dec  7  2018 /var/run -> ../run
# journalctl -f
Jan 20 17:36:34 fs01 setroubleshoot[6109]: SELinux is preventing /usr/sbin/sshd from connectto access on the unix_stream_socket /run/docker.sock. For complete SELinux messages run: sealert -l 9ca80d93-76bf-47fb-9116-0df43240da8b
Jan 20 17:36:34 fs01 python[6109]: SELinux is preventing /usr/sbin/sshd from connectto access on the unix_stream_socket /run/docker.sock.

                                                   *****  Plugin catchall_boolean (89.3 confidence) suggests   ******************

                                                   If you want to allow daemons to enable cluster mode
                                                   Then you must tell SELinux about this by enabling the 'daemons_enable_cluster_mode' boolean.

                                                   Do
                                                   setsebool -P daemons_enable_cluster_mode 1

                                                   *****  Plugin catchall (11.6 confidence) suggests   **************************

                                                   If you believe that sshd should be allowed connectto access on the docker.sock unix_stream_socket by default.
                                                   Then you should report this as a bug.
                                                   You can generate a local policy module to allow this access.
                                                   Do
                                                   allow this access for now by executing:
                                                   # ausearch -c 'sshd' --raw | audit2allow -M my-sshd
                                                   # semodule -i my-sshd.pp
```

## Build

### Build local

`go fmt ./...&&goimports -w .&&golint ./...&&golangci-lint run --enable-all&& go install  -ldflags="-s -w" ./..`

### Build for linux

`env GOOS=linux GOARCH=amd64 go install -ldflags="-s -w" ./...`

## Unix HTTP server

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

## Unix HTTP client

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
