package compressor

type baseCase struct {
	name    string
	arg     any
	wantErr bool
}

func initTestCases() []baseCase {
	p := 1
	var pn *int

	return []baseCase{
		{
			name:    "String",
			arg:     "test",
			wantErr: false,
		},
		{
			name:    "Int",
			arg:     123,
			wantErr: false,
		},
		{
			name:    "Uint",
			arg:     uint(123),
			wantErr: false,
		},
		{
			name:    "Float",
			arg:     3.1,
			wantErr: false,
		},
		{
			name:    "Bool",
			arg:     true,
			wantErr: false,
		},
		{
			name:    "Complex",
			arg:     3.1 + 1i,
			wantErr: false,
		},
		{
			name:    "Slice",
			arg:     []any{123, 12.23, "test"},
			wantErr: false,
		},
		{
			name:    "Map",
			arg:     map[string]any{"test": 123},
			wantErr: false,
		},
		{
			name:    "Pointer",
			arg:     &p,
			wantErr: false,
		},
		{
			name:    "Bytes",
			arg:     []byte("Hello world"),
			wantErr: false,
		},
		{
			name:    "Nil Pointer",
			arg:     pn,
			wantErr: true,
		},
		{
			name:    "Chan",
			arg:     make(chan int),
			wantErr: true,
		},
	}
}
