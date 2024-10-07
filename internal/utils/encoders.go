package utils

import (
	b64 "encoding/base64"
	"encoding/binary"
	"errors"
	"github.com/google/uuid"
)

func EncodeCookieValue(sessionID uuid.UUID, userID int64) string {
	cookieValueBytes := make([]byte, 0, 24)
	uuidBytes, _ := sessionID.MarshalBinary()
	cookieValueBytes = append(cookieValueBytes, uuidBytes...)
	cookieValueBytes = binary.LittleEndian.AppendUint64(cookieValueBytes, uint64(userID))
	return b64.StdEncoding.EncodeToString(cookieValueBytes)
}

func DecodeCookieValue(cookieValue string) (sessionID uuid.UUID, userID int64, err error) {
	var decodedValue []byte
	decodedValue, err = b64.StdEncoding.DecodeString(cookieValue)
	if err != nil {
		return
	}

	if len(decodedValue) != 24 {
		err = errors.New("invalid cookie value")
		return
	}

	if err = sessionID.UnmarshalBinary(decodedValue[:16]); err != nil {
		return
	}
	userID = int64(binary.LittleEndian.Uint64(decodedValue[16:]))

	return
}
