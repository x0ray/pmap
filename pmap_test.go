// pmap package unit tests
package pmap

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name string
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// Test cases
		{
			name:    "good arg",
			args:    args{name: "any", path: ""},
			wantErr: false,
		},
		{
			name:    "bad arg",
			args:    args{name: "", path: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.name, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPmap_Copy(t *testing.T) {
	type fields struct {
		Name string
		Path string
		Pm   map[string]interface{}
	}
	type args struct {
		newname string
	}

	f := fields{"test", "/tmp/path", map[string]interface{}{"a": 123, "b": "abc"}}
	pm, _ := New("new", f.Path)
	pm.Add("a", 123)
	pm.Add("b", "abc")

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Pmap
		wantErr bool
	}{
		// Test cases
		{
			name:    "good arg",
			fields:  f,
			args:    args{newname: "new"},
			want:    pm,
			wantErr: false,
		},
		{
			name:    "bad name arg",
			fields:  f,
			args:    args{newname: ""},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pmap{
				Name: tt.fields.Name,
				Path: tt.fields.Path,
				Pm:   tt.fields.Pm,
			}
			got, err := p.Copy(tt.args.newname)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pmap.Copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pmap.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
