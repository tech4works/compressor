package compressor

import (
	"testing"
)

func TestToDeflate(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("ToDeflate() did not panic")
					}
				}()
			}
			ToDeflate(tt.arg)
		})
	}
}

func TestToDeflateWithErr(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ToDeflateWithErr(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDeflateWithErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToDeflateBase64(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("ToDeflateBase64() did not panic")
					}
				}()
			}
			ToDeflateBase64(tt.arg)
		})
	}
}

func TestToDeflateBase64WithErr(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ToDeflateBase64WithErr(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDeflateBase64WithErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
