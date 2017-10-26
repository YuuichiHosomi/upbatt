package main

import (
	"bufio"
	"os"
	"time"
)

const dataLogPath = "/var/lib/upbatt/data.log"

// DataLog test struct
type DataLog struct {
	File   *os.File
	Writer *bufio.Writer
	Chan   chan string
}

// NewDataLog test
func NewDataLog() (*DataLog, error) {
	var dl DataLog
	f, err := os.OpenFile(dataLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	dl.Writer = bufio.NewWriter(f)
	dl.File = f
	dl.Chan = make(chan string) // unbuffered, to keep sync
	go dl.pump()

	return &dl, nil
}

// internal channel pump
func (dl *DataLog) pump() {
	for str := range dl.Chan {
		dl.Writer.WriteString(str)
		dl.Writer.Flush()
	}
}

// AppendRaw test
func (dl *DataLog) AppendRaw(str string) {
	dl.Chan <- str
}

// Append test
func (dl *DataLog) Append(str string) {
	dl.AppendRaw(time.Now().Format(time.RFC3339) + ";" + str + "\n")
}