package bathtext

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		expected string
		wantErr  bool
	}{
		{
			name:     "positive amount",
			amount:   100.50,
			expected: "100.50 บาท",
			wantErr:  false,
		},
		{
			name:     "zero amount",
			amount:   0,
			expected: "0.00 บาท",
			wantErr:  false,
		},
		{
			name:     "negative amount",
			amount:   -10,
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Convert(tt.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if result != tt.expected {
				t.Errorf("Convert() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestConvertBaht(t *testing.T) {
	amount := 50.25
	expected := "50.25 บาท"
	
	result, err := ConvertBaht(amount)
	if err != nil {
		t.Errorf("ConvertBaht() unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("ConvertBaht() = %v, want %v", result, expected)
	}
}

func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}
}