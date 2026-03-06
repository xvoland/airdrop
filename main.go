package main

/*
#cgo LDFLAGS: -L. -lairdrop -framework Cocoa
#include "airdrop.h"
#include <stdlib.h>
*/
import "C"

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var (
	verbose     bool
	showVersion bool
	version     = "0.3.6"
)

var tempFiles []string

func colorize(text, colorCode string) string {
	reset := "\033[0m"
	return colorCode + text + reset
}

func logf(format string, a ...interface{}) {
	if verbose {
		fmt.Fprintf(os.Stderr, format+"\n", a...)
	}
}

func cleanup() {
	for _, f := range tempFiles {
		_ = os.Remove(f)
	}
}

func detectExtFromBytes(sample []byte) string {
	ct := http.DetectContentType(sample)
	switch {
	case strings.HasPrefix(ct, "image/png"):
		return ".png"
	case strings.HasPrefix(ct, "image/jpeg"):
		return ".jpg"
	case strings.HasPrefix(ct, "image/gif"):
		return ".gif"
	case strings.HasPrefix(ct, "application/pdf"):
		return ".pdf"
	case strings.HasPrefix(ct, "text/"):
		return ".txt"
	default:
		return ""
	}
}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.BoolVar(&verbose, "v", false, "verbose logging")
	flag.BoolVar(&showVersion, "version", false, "show version information")

	flag.Usage = func() {
		cyan := "\033[36m"
		yellow := "\033[33m"
		bold := "\033[1m"

		year := time.Now().Year()
		fmt.Fprintf(os.Stderr, "%s\n\n", colorize(fmt.Sprintf("Copyright © %d, Vitalii Tereshchuk | All rights reserved | https://dotoca.net/airdrop", year), cyan))
		fmt.Fprintf(os.Stderr, "%s%s%s\n\n", bold, colorize("CLI Utility for Apple AirDrop — version "+version, yellow), "\033[0m")
		fmt.Fprintf(os.Stderr, "%s\n", colorize("Usage:", bold))
		fmt.Fprintf(os.Stderr, "  %s file1 file2 ...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  cat file.pdf | %s \n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s\n", colorize("Options:", bold))
		flag.PrintDefaults()
	}

	flag.Parse()

	if showVersion {
		fmt.Println("airdrop version " + version)
		os.Exit(0)
	}

	info, _ := os.Stdin.Stat()
	noArgs := flag.NArg() == 0
	stdinIsPipe := (info.Mode() & os.ModeCharDevice) == 0

	if noArgs && !stdinIsPipe {
		flag.Usage()
		os.Exit(1)
	}

	defer cleanup()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-c
		logf("Caught signal %v — cleaning up", s)
		cleanup()
		os.Exit(130)
	}()

	var files []string

	if flag.NArg() > 0 {
		files = append(files, flag.Args()...)
	} else {
		const sniffLen = 512
		buf := make([]byte, sniffLen)
		n, err := io.ReadAtLeast(os.Stdin, buf, 1)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			fmt.Fprintln(os.Stderr, "Failed to read stdin for sniffing:", err)
			os.Exit(2)
		}
		sample := buf[:n]

		ext := detectExtFromBytes(sample)
		tmp, err := os.CreateTemp("", "airdrop_stdin_*"+ext)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create temp file:", err)
			os.Exit(3)
		}
		tmpName := tmp.Name()
		tempFiles = append(tempFiles, tmpName)

		if _, err = tmp.Write(sample); err != nil {
			_ = tmp.Close()
			fmt.Fprintln(os.Stderr, "Failed to write to temp file:", err)
			os.Exit(4)
		}
		if _, err = io.Copy(tmp, os.Stdin); err != nil {
			_ = tmp.Close()
			fmt.Fprintln(os.Stderr, "Failed to write rest of stdin:", err)
			os.Exit(5)
		}
		if err := tmp.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to close temp file:", err)
			os.Exit(6)
		}
		files = append(files, tmpName)
		logf("stdin saved to %s (detected ext %s)", tmpName, ext)
	}

	for _, f := range files {
		if _, err := os.Stat(f); err != nil {
			fmt.Fprintln(os.Stderr, "File not accessible:", f, ":", err)
			os.Exit(10)
		}
	}

	cFiles := make([]*C.char, len(files))
	for i, f := range files {
		abs := f
		if !filepath.IsAbs(f) {
			if a, err := filepath.Abs(f); err == nil {
				abs = a
			}
		}
		cFiles[i] = C.CString(abs)
		defer C.free(unsafe.Pointer(cFiles[i]))
		logf("prepared file %d -> %s", i, abs)
	}

	if len(cFiles) == 0 {
		fmt.Fprintln(os.Stderr, "No files to share")
		os.Exit(11)
	}

	ptr := (**C.char)(unsafe.Pointer(&cFiles[0]))
	res := C.ShareViaAirDrop(ptr, C.int(len(cFiles)))
	if res != 0 {
		fmt.Fprintln(os.Stderr, "AirDrop failed with code:", int(res))
		os.Exit(20 + int(res))
	}
	fmt.Println("AirDrop succeeded")
}
