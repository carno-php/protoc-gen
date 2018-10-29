package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Error(messages ...string) {
	Fatal(errors.New(strings.Join(messages, " ")))
}

func Fatal(err error) {
	fmt.Fprintln(os.Stderr, "fatal error: "+err.Error())
	os.Exit(1)
}
