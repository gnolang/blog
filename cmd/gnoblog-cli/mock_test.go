package main

import (
	"io"
)

type (
	mockGetPassword func() (string, error)
)

type mockIO struct {
	getPassword mockGetPassword
}

func (m mockIO) In() io.Reader {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) Out() io.WriteCloser {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) Err() io.WriteCloser {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) SetIn(in io.Reader) {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) SetOut(out io.WriteCloser) {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) SetErr(err io.WriteCloser) {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) Println(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) Printf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) Printfln(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) ErrPrintln(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) ErrPrintfln(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) GetCheckPassword(prompts [2]string, insecure bool) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) GetConfirmation(prompt string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockIO) GetPassword(prompt string, insecure bool) (string, error) {
	if m.getPassword != nil {
		return m.getPassword()
	}
	return "", nil
}

func (m mockIO) GetString(prompt string) (string, error) {
	//TODO implement me
	panic("implement me")
}
