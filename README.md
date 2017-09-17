[![Build Status](https://travis-ci.org/zenreach/bcrypt.svg?branch=master)](https://travis-ci.org/zenreach/bcrypt) [![Go Report Card](https://goreportcard.com/badge/github.com/zenreach/bcrypt)](https://goreportcard.com/report/github.com/zenreach/bcrypt)

BCrypt Password Encryptor
=========================
This package contains a tool to encrypt passwords using the bcrypt password
algorithm.

Install
-------
You may download a binary for your plaform from the [releases page][1].

Alternatively you can install the latest master commit with `go get`:

    $ go get github.com/zenreach/bcrypt

This will build and install the `bcrypt` binary into your `$GOBIN`.

Usage
-----
There are three ways to get a password into `bcrypt`.

As the first positional argument on the command line:

    $ bcrypt secret
    $2a$10$p9XJ4Np9HzSSZsTJqkun/eDAXpCbYl8xncr2srCfHsLnTYHIEsx/m

Pipe the password from from another process:

    $ echo secret | bcrypt
    $2a$10$hDlfF5bNs1Mcwx4sJdvV9e9.FmCaphoUms6q0qFRq2YgVjG2iOY4e

Let `bcrypt` prompt you for a password:

    $ bcrypt
    password:
    $2a$10$BQ6oxJbFwGYJqPwu9Hp.V.Y3qarGXMPnnp9NrmIpY34AqLAv3pz4S

The hashing cost can be adjusted with the `-c` flag. Higher cost values
increase the time it takes to encrypt a password and, by extension, compare
passwords. See the output of `-h` for the default and valid options. This
example sets the cost to 15:

    $ bcrypt -c 15 secret
    $2a$15$zsZTRJ9wFqnLbjrM4vgmNO46lL13yolww5yKSRNuiVY.9Afaq07Ia

The bcrypt algorithm has a length limit of up to 55 bytes. As a result longer
passwords are effectively truncated to this length. This can be worked around
by using the `-s` flag to SHA256 encode the password before encrypting:

    $ bcrypt -s secret
    $2a$10$Rk2ZHFXHR/7sYcli6BudfuX/8xVUMCIedxTnZ0ZDAHVPDvIzXi82O

The `bcrypt` command is careful to output prompts and error messages to stdin.
Only the hash is printed to stdout. This alows its output to be redirected for
use in scripts. This simple example prompts the user for a password and writes
the encrypted hash to a file:

    $ bcrypt > password.hash
    Password:
    $ cat password.hash
    $2a$10$aNpcOMIZjYTmG1EzAg.lleDmnmHKvsg13xhG55J.3xEme4GU.HqHu

Within a script you may wish to capture the password to a variable:

    #!/bin/bash
    HASH=$(bcrypt)
    echo 'user:$HASH' >> /etc/myapp/users.pw


[1]: https://github.com/zenreach/bcrypt/releases "Releases"
