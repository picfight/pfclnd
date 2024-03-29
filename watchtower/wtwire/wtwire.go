package wtwire

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/picfight/pfcd/pfcec"
	"github.com/picfight/pfcd/wire"
	"github.com/picfight/pfclnd/lnwallet"
)

// WriteElement is a one-stop shop to write the big endian representation of
// any element which is to be serialized for the wire protocol. The passed
// io.Writer should be backed by an appropriately sized byte slice, or be able
// to dynamically expand to accommodate additional data.
func WriteElement(w io.Writer, element interface{}) error {
	switch e := element.(type) {
	case uint8:
		var b [1]byte
		b[0] = e
		if _, err := w.Write(b[:]); err != nil {
			return err
		}

	case uint16:
		var b [2]byte
		binary.BigEndian.PutUint16(b[:], e)
		if _, err := w.Write(b[:]); err != nil {
			return err
		}

	case uint32:
		var b [4]byte
		binary.BigEndian.PutUint32(b[:], e)
		if _, err := w.Write(b[:]); err != nil {
			return err
		}

	case uint64:
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], e)
		if _, err := w.Write(b[:]); err != nil {
			return err
		}

	case [16]byte:
		if _, err := w.Write(e[:]); err != nil {
			return err
		}

	case [32]byte:
		if _, err := w.Write(e[:]); err != nil {
			return err
		}

	case [33]byte:
		if _, err := w.Write(e[:]); err != nil {
			return err
		}

	case []byte:
		if err := wire.WriteVarBytes(w, 0, e); err != nil {
			return err
		}

	case lnwallet.SatPerKWeight:
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], uint64(e))
		if _, err := w.Write(b[:]); err != nil {
			return err
		}

	case ErrorCode:
		var b [2]byte
		binary.BigEndian.PutUint16(b[:], uint16(e))
		if _, err := w.Write(b[:]); err != nil {
			return err
		}

	case *pfcec.PublicKey:
		if e == nil {
			return fmt.Errorf("cannot write nil pubkey")
		}

		var b [33]byte
		serializedPubkey := e.SerializeCompressed()
		copy(b[:], serializedPubkey)
		if _, err := w.Write(b[:]); err != nil {
			return err
		}

	default:
		return fmt.Errorf("Unknown type in WriteElement: %T", e)
	}

	return nil
}

// WriteElements is writes each element in the elements slice to the passed
// io.Writer using WriteElement.
func WriteElements(w io.Writer, elements ...interface{}) error {
	for _, element := range elements {
		err := WriteElement(w, element)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadElement is a one-stop utility function to deserialize any datastructure
// encoded using the serialization format of lnwire.
func ReadElement(r io.Reader, element interface{}) error {
	switch e := element.(type) {
	case *uint8:
		var b [1]uint8
		if _, err := r.Read(b[:]); err != nil {
			return err
		}
		*e = b[0]

	case *uint16:
		var b [2]byte
		if _, err := io.ReadFull(r, b[:]); err != nil {
			return err
		}
		*e = binary.BigEndian.Uint16(b[:])

	case *uint32:
		var b [4]byte
		if _, err := io.ReadFull(r, b[:]); err != nil {
			return err
		}
		*e = binary.BigEndian.Uint32(b[:])

	case *uint64:
		var b [8]byte
		if _, err := io.ReadFull(r, b[:]); err != nil {
			return err
		}
		*e = binary.BigEndian.Uint64(b[:])

	case *[16]byte:
		if _, err := io.ReadFull(r, e[:]); err != nil {
			return err
		}

	case *[32]byte:
		if _, err := io.ReadFull(r, e[:]); err != nil {
			return err
		}

	case *[33]byte:
		if _, err := io.ReadFull(r, e[:]); err != nil {
			return err
		}

	case *[]byte:
		bytes, err := wire.ReadVarBytes(r, 0, 66000, "[]byte")
		if err != nil {
			return err
		}
		*e = bytes

	case *lnwallet.SatPerKWeight:
		var b [8]byte
		if _, err := io.ReadFull(r, b[:]); err != nil {
			return err
		}
		*e = lnwallet.SatPerKWeight(binary.BigEndian.Uint64(b[:]))

	case *ErrorCode:
		var b [2]byte
		if _, err := io.ReadFull(r, b[:]); err != nil {
			return err
		}
		*e = ErrorCode(binary.BigEndian.Uint16(b[:]))

	case **pfcec.PublicKey:
		var b [pfcec.PubKeyBytesLenCompressed]byte
		if _, err := io.ReadFull(r, b[:]); err != nil {
			return err
		}

		pubKey, err := pfcec.ParsePubKey(b[:], pfcec.S256())
		if err != nil {
			return err
		}
		*e = pubKey

	default:
		return fmt.Errorf("Unknown type in ReadElement: %T", e)
	}

	return nil
}

// ReadElements deserializes a variable number of elements into the passed
// io.Reader, with each element being deserialized according to the ReadElement
// function.
func ReadElements(r io.Reader, elements ...interface{}) error {
	for _, element := range elements {
		err := ReadElement(r, element)
		if err != nil {
			return err
		}
	}
	return nil
}
