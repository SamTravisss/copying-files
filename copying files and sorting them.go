import (

	// #nosec

	"testing"

	"github.com/stretchr/testify/require"
)

for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := ValidateApplePay(tt.msg)
			require.Equal(t, tt.wantErr, err != nil)
		})
