```sh
go vet cmd/nocopy/nocopy.go 
# command-line-arguments
cmd/nocopy/nocopy.go:16:13: Copy passes lock by value: command-line-arguments.Demo contains command-line-arguments.noCopy
cmd/nocopy/nocopy.go:17:12: call of CopyTwice copies lock value: command-line-arguments.Demo contains command-line-arguments.noCopy
cmd/nocopy/nocopy.go:19:18: CopyTwice passes lock by value: command-line-arguments.Demo contains command-line-arguments.noCopy
cmd/nocopy/nocopy.go:23:20: call of fmt.Printf copies lock value: command-line-arguments.Demo contains command-line-arguments.noCopy
cmd/nocopy/nocopy.go:25:7: call of Copy copies lock value: command-line-arguments.Demo contains command-line-arguments.noCopy
cmd/nocopy/nocopy.go:27:20: call of fmt.Printf copies lock value: command-line-arguments.Demo contains command-line-arguments.noCopy

```