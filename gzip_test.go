package compressor

import (
	"testing"
)

func TestToGzip(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("ToGzip() did not panic")
					}
				}()
			}
			ToGzip(tt.arg)
		})
	}
}

func TestToGzipWithErr(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ToGzipWithErr(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToGzipWithErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToGzipBase64(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("ToGzipBase64() did not panic")
					}
				}()
			}
			ToGzipBase64(tt.arg)
		})
	}
}

func TestToGzipBase64WithErr(t *testing.T) {
	for _, tt := range initTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ToGzipBase64WithErr(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToGzipBase64WithErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
