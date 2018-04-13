# b64
> base64 command-line utlity

# Usage

After installing, you can run the command `b64` to see the following help menu:

```
NAME:
   b64 - base64 encoder and decoder

USAGE:
   b64 [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     encode   encode the given string(s) as base64
     decode   decode the given string(s) as base64
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

To encode a given string with base64 encoding, you can use the `encode` flag:

```shell
$ b64 encode "Example String"
RXhhbXBsZSBTdHJpbmc=
```

To decode a given string that been base64 encoded there's the `decode` flag:

```shell
$ b64 decode "RXhhbXBsZSBTdHJpbmc="
Example String
```
