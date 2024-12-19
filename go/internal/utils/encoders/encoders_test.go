package encoders

import (
	"github.com/google/uuid"
	"testing"
)

func TestDecodeCookieValue(t *testing.T) {
	validSessionID := uuid.New()
	validUserID := int64(1234)

	testCases := []struct {
		name              string
		value             string
		expectedSessionID uuid.UUID
		expectedUserID    int64
		expectedErr       bool
	}{
		{
			name:              "valid value",
			value:             EncodeCookieValue(validSessionID, validUserID),
			expectedSessionID: validSessionID,
			expectedUserID:    validUserID,
		},
		{
			name:        "invalid value",
			value:       "invalid",
			expectedErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			sessionID, userID, err := DecodeCookieValue(test.value)
			if test.expectedErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if sessionID != test.expectedSessionID {
					t.Errorf("expected sessionID %s, got %s", test.expectedSessionID, sessionID)
				}
				if userID != test.expectedUserID {
					t.Errorf("expected userID %d, got %d", test.expectedUserID, userID)
				}
			}
		})
	}
}
