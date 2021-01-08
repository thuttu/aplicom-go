package dprotocol

import (
	"bufio"
	"io"
)

// Scanner scans D protocol packets.
type Scanner struct {
	sc     *bufio.Scanner
	packet Packet
	err    error
}

// NewScanner creates a new scanner to scan D protocol packets from the provided io.Reader.
func NewScanner(r io.Reader) *Scanner {
	sc := bufio.NewScanner(r)
	sc.Split(ScanPackets)
	return &Scanner{sc: sc}
}

// ScanPacket advances the Scanner to the next packet, which will then be
// available through the Packet or Bytes method. It returns false when the
// scan stops, either by reaching the end of the input or an error.
// After ScanPacket returns false, the Err method will return any error that
// occurred during scanning, except that if it was io.EOF, Err
// will return nil.
func (s *Scanner) ScanPacket() bool {
	if !s.sc.Scan() {
		s.err = s.sc.Err()
		return false
	}
	if err := s.packet.UnmarshalBinary(s.sc.Bytes()); err != nil {
		s.err = err
		return false
	}
	return true
}

// Packet returns a reference to the last scanned packet.
func (s *Scanner) Packet() *Packet {
	return &s.packet
}

// Bytes returns the raw bytes of the last scanned packet.
func (s *Scanner) Bytes() []byte {
	return s.sc.Bytes()
}

// Err returns the first non-EOF error that was encountered by the Scanner.
func (s *Scanner) Err() error {
	return s.err
}
