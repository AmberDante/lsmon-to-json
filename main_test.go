package main

import (
	"reflect"
	"testing"
)

func Test_trimBorders(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test trim borders",
			args: args{`  Sentinel RMS Development Kit 9.1.0.0104 Application Monitor
  Copyright (C) 2016 SafeNet, Inc.

 [Contacting Sentinel RMS Development Kit server on host "999385-pc.samba.gazpromproject.ru"]


 |- Feature Information
   |- Feature name                   : "ARMS_ID"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 
   |- License Version                : 0x08600000
   |- Commuter license allowed       : NO
   |- Maximum concurrent user(s)     : 99999
     |- Sharing limit                  : 25
     |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
     |- Allowed on VM                  : YES
Press Enter to continue . . .
`},
			want: ` |- Feature Information
   |- Feature name                   : "ARMS_ID"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 
   |- License Version                : 0x08600000
   |- Commuter license allowed       : NO
   |- Maximum concurrent user(s)     : 99999
     |- Sharing limit                  : 25
     |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
     |- Allowed on VM                  : YES`,
		},
		{
			name: "test trim beginning",
			args: args{`  Sentinel RMS Development Kit 9.1.0.0104 Application Monitor
  Copyright (C) 2016 SafeNet, Inc.

 [Contacting Sentinel RMS Development Kit server on host "999385-pc.samba.gazpromproject.ru"]


 |- Feature Information
   |- Feature name                   : "ARMS_ID"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 
   |- License Version                : 0x08600000
   |- Commuter license allowed       : NO
   |- Maximum concurrent user(s)     : 99999
     |- Sharing limit                  : 25
     |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
     |- Allowed on VM                  : YES

`},
			want: ` |- Feature Information
   |- Feature name                   : "ARMS_ID"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 
   |- License Version                : 0x08600000
   |- Commuter license allowed       : NO
   |- Maximum concurrent user(s)     : 99999
     |- Sharing limit                  : 25
     |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
     |- Allowed on VM                  : YES`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBorders(tt.args.s); got != tt.want {
				t.Errorf("trimBorders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFeature(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want feature
	}{
		// TODO: Add test cases.
		{
			name: "parse feature 1",
			args: args{`   |- Feature name                   : "ARMS_ID"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 
   |- License Version                : 0x08600000
   |- Commuter license allowed       : NO
   |- Maximum concurrent user(s)     : 99999
   |- Unreserved tokens in use       : 0

`},
			want: feature{feature: map[string]string{
				"Feature name":               "ARMS_ID",
				"Feature version":            "1.0",
				"License type":               "Normal License",
				"License Version":            "0x08600000",
				"Commuter license allowed":   "NO",
				"Maximum concurrent user(s)": "99999",
				"Unreserved tokens in use":   "0",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFeature(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFeature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_textToMap(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name         string
		args         args
		wantBlockMap map[string]string
	}{
		// TODO: Add test cases.
		{
			name: "block 1",
			args: args{[]string{
				`   |- Feature name                   : "ARMS_ID"  	
        `,
				`   |- Feature version                : "1.0"
        `,
				"\n",
				`   |- License Version                : 0x08600000`,
			}},
			wantBlockMap: map[string]string{
				"Feature name":    "ARMS_ID",
				"Feature version": "1.0",
				"License Version": "0x08600000",
			},
		},
		{
			name: "block 2",
			args: args{[]string{
				"sjlinsjkjksfj",
				"83hkjsl8j8;98w",
			}},
			wantBlockMap: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBlockMap := textToMap(tt.args.slice); !reflect.DeepEqual(gotBlockMap, tt.wantBlockMap) {
				t.Errorf("textToMap() = %v, want %v", gotBlockMap, tt.wantBlockMap)
			}
		})
	}
}
