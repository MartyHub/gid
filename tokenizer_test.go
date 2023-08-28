package gid

import "testing"

func TestTokenizer_ToCamel(t *testing.T) {
	t.Parallel()

	tokenizer := Default()

	tests := []struct {
		name         string
		wantExport   string
		wantUnexport string
	}{
		{
			name:         "",
			wantExport:   "A",
			wantUnexport: "a",
		},
		{
			name:         "123",
			wantExport:   "A123",
			wantUnexport: "_123",
		},
		{
			name:         "word",
			wantExport:   "Word",
			wantUnexport: "word",
		},
		{
			name:         "Word",
			wantExport:   "Word",
			wantUnexport: "word",
		},
		{
			name:         "WORD",
			wantExport:   "Word",
			wantUnexport: "word",
		},
		{
			name:         "word1 word2_word3-word4#word5@word6&word7",
			wantExport:   "Word1Word2Word3Word4Word5Word6Word7",
			wantUnexport: "word1Word2Word3Word4Word5Word6Word7",
		},
		{
			name:         "id",
			wantExport:   "ID",
			wantUnexport: "id",
		},
		{
			name:         "myId",
			wantExport:   "MyID",
			wantUnexport: "myID",
		},
		{
			name:         "jsonId",
			wantExport:   "JSONId",
			wantUnexport: "jsonID",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tokenizer.ExportID(tt.name); got != tt.wantExport {
				t.Errorf("ExportID( %v ) = %v, want %v", tt.name, got, tt.wantExport)
			}

			if got := tokenizer.UnexportID(tt.name); got != tt.wantUnexport {
				t.Errorf("UnexportID( %v ) = %v, want %v", tt.name, got, tt.wantUnexport)
			}
		})
	}
}
