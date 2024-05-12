package main

import (
	"log"
	"net"
	"net/textproto"

	"github.com/ryuichi1208/go-milter"
)

type MyMilter struct {
	milter.Milter
}

// Connect logs SMTP connection data.
func (m *MyMilter) Connect(host string, family string, port uint16, addr net.IP, mod *milter.Modifier) (milter.Response, error) {
	log.Printf("Connect from host: %s, family: %s, port: %d, IP: %s", host, family, port, addr)
	return milter.RespContinue, nil
}

// Helo logs HELO/EHLO command.
func (m *MyMilter) Helo(name string, mod *milter.Modifier) (milter.Response, error) {
	log.Printf("HELO/EHLO: %s", name)
	return milter.RespContinue, nil
}

// MailFrom logs the MAIL FROM command.
func (m *MyMilter) MailFrom(from string, mod *milter.Modifier) (milter.Response, error) {
	log.Printf("Mail from: %s", from)
	return milter.RespContinue, nil
}

// RcptTo logs the RCPT TO command.
func (m *MyMilter) RcptTo(rcptTo string, mod *milter.Modifier) (milter.Response, error) {
	log.Printf("Rcpt to: %s", rcptTo)
	return milter.RespContinue, nil
}

// Header logs each message header.
func (m *MyMilter) Header(name string, value string, mod *milter.Modifier) (milter.Response, error) {
	log.Printf("Header: %s: %s", name, value)
	return milter.RespContinue, nil
}

// Headers is called when all headers have been processed.
func (m *MyMilter) Headers(h textproto.MIMEHeader, mod *milter.Modifier) (milter.Response, error) {
	log.Printf("All headers processed.")
	return milter.RespContinue, nil
}

// BodyChunk processes body chunks.
func (m *MyMilter) BodyChunk(chunk []byte, mod *milter.Modifier) (milter.Response, error) {
	log.Printf("Body chunk: %d bytes", len(chunk))
	return milter.RespContinue, nil
}

// Body is called at the end of the message.
func (m *MyMilter) Body(mod *milter.Modifier) (milter.Response, error) {
	log.Println("End of message body.")
	return milter.RespContinue, nil
}

// Abort handles message aborts.
func (m *MyMilter) Abort(mod *milter.Modifier) error {
	log.Println("Message processing aborted.")
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("Failed to listen on :12345: %v", err)
	}

	newMilter := func() milter.Milter {
		return &MyMilter{}
	}

	server := milter.Server{
		NewMilter: newMilter,
	}

	log.Println("Starting milter server on :12345")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
