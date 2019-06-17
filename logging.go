package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"

	"github.com/docker/docker/api/types/plugins/logdriver"
	pb "github.com/gogo/protobuf/io"
	"github.com/pkg/errors"
	"github.com/tonistiigi/fifo"
)

//HelloLogger maintains the state for the log line handler
type HelloLogger struct {
	logFile io.ReadWriteCloser
}

func newLogger(file string) (HelloLogger, error) {
	inputFile, err := fifo.OpenFifo(context.Background(), file, syscall.O_RDONLY, 0700)
	if err != nil {
		return HelloLogger{}, errors.Wrapf(err, "error opening logger fifo: %q", file)
	}
	fmt.Fprintf(os.Stderr, "Created new logger for %s\n", file)
	return HelloLogger{logFile: inputFile}, nil
}

func (logger *HelloLogger) consumeLogs() {
	reader := pb.NewUint32DelimitedReader(logger.logFile, binary.BigEndian, 2e6)
	defer reader.Close()
	var log logdriver.LogEntry
	for {
		err := reader.ReadMsg(&log)
		if err != nil {
			if err == io.EOF || err == os.ErrClosed || strings.Contains(err.Error(), "file already closed") {
				logger.logFile.Close()
				return
			}
			reader = pb.NewUint32DelimitedReader(logger.logFile, binary.BigEndian, 2e6)
		}
		fmt.Fprintf(os.Stderr, "Got log line: %s\n", log.Line)
		log.Reset()
	}

}
