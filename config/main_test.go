package config_test

import (
	"github.com/kayslay/todaytvseries_organizer/config"
	"regexp"
	"testing"
)

var defReg = (regexp.MustCompile(`(.+)?\.S(\d{2})`))

func TestConfig_GetDir(t *testing.T) {
	type fields struct {
		DeleteAfter bool
		MoveDir     string
		Path        string
		Ext         string
		MatchExt    string
		FolderName  config.ConfigReg
		WorkerCount int8
	}
	type args struct {
		name string
	}
	defaultField := fields{
		DeleteAfter: false,
		MoveDir:     "./",
		Path:        "./",
		Ext:         ".zip",
		MatchExt:    ".go",
		FolderName:  config.ConfigReg{defReg},
		WorkerCount: 1,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"default_config", defaultField, args{"GOT.S02.zip"}, "GOT/02"},
		{"default_config", defaultField, args{"GOT.S02E01.zip"}, "GOT/02"},
		{"default_config", defaultField, args{"24hours.S02.zip"}, "24hours/02"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := config.Config{
				DeleteAfter: tt.fields.DeleteAfter,
				MoveDir:     tt.fields.MoveDir,
				Path:        tt.fields.Path,
				Ext:         tt.fields.Ext,
				MatchExt:    tt.fields.MatchExt,
				FolderName:  tt.fields.FolderName,
				WorkerCount: tt.fields.WorkerCount,
			}
			if got := config.GetDir(tt.args.name); got != tt.want {
				t.Errorf("Config.GetDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
