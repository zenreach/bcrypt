// Package main contains a command to encrypt passwords using bcrypt. There are
// three ways to provide a password: via the command line as the first
// argument, via a pipe, or via a prompt.
package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

	"github.com/spf13/pflag"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	flags := parseFlags()
	cost, _ := flags.GetInt("cost")
	sha, _ := flags.GetBool("sha256")

	var password []byte
	args := flags.Args()
	switch len(args) {
	case 0:
		var err error
		if isTTY() {
			fmt.Fprint(os.Stderr, "Password: ")
			password, err = terminal.ReadPassword(int(syscall.Stdin))
			fmt.Fprint(os.Stderr, "\n")
		} else {
			password, err = ioutil.ReadAll(os.Stdin)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading stdin: %s\n", err)
			os.Exit(1)
		}
	case 1:
		password = []byte(args[0])
	default:
		flags.Usage()
		os.Exit(1)
	}

	if sha {
		hash := sha256.Sum256(password)
		password = hash[:]
	}

	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error encrypting password: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", hash)
}

func parseFlags() *pflag.FlagSet {
	script := filepath.Base(os.Args[0])
	f := pflag.NewFlagSet(script, pflag.ExitOnError)
	f.SetOutput(os.Stderr)

	costHelp := fmt.Sprintf(
		"Hashing cost. Valid values are from %d to %d.",
		bcrypt.MinCost, bcrypt.MaxCost,
	)
	f.IntP("cost", "c", bcrypt.DefaultCost, costHelp)

	shaHelp := "SHA256 encode the password before encrypting."
	f.BoolP("sha256", "s", false, shaHelp)

	usage := func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [PASSWORD]\n", script)
		fmt.Fprintf(os.Stderr, "Encrypts a password with the bcrypt hashing algorithm.\n\n")
		fmt.Fprintln(os.Stderr, "Options:")
		f.PrintDefaults()
	}
	f.Usage = usage
	f.Parse(os.Args[1:])
	return f
}

func isTTY() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return false
	}
	return true
}
