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
			want: feature{Feature: map[string]string{
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
		{
			name: "block 3",
			args: args{[]string{
				`     |- License storage name           : C:\AVEVA\AVEVA Licensing System\RMS\lservrc_AVEVA`,
			}},
			wantBlockMap: map[string]string{
				"License storage name": "C:\\AVEVA\\AVEVA Licensing System\\RMS\\lservrc_AVEVA",
			},
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

func Test_getFeturesInfo(t *testing.T) {
	type args struct {
		lsmonOut string
	}
	tests := []struct {
		name string
		args args
		want fetures
	}{
		// TODO: Add test cases.
		{
			name: "Test_getFeturesInfo test border trim",
			args: args{`Sentinel RMS Development Kit 9.1.0.0104 Application Monitor
  Copyright (C) 2016 SafeNet, Inc.

 [Contacting Sentinel RMS Development Kit server on host "999385-pc.samba.gazpromproject.ru"]


 |- Feature Information
   |- Feature name                   : "ARMS_ID"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 
   |- License Version                : 0x08600000
   |- Commuter license allowed       : NO
   |- Maximum concurrent user(s)     : 99999
 Press Enter to continue . . .`},
			want: fetures{Features: []feature{
				{
					Feature: map[string]string{
						"Feature name":               "ARMS_ID",
						"Feature version":            "1.0",
						"License Version":            "0x08600000",
						"License type":               "Normal License",
						"Commuter license allowed":   "NO",
						"Maximum concurrent user(s)": "99999",
					},
					LicenseInformation: nil,
					ClientInformation:  nil,
				},
			},
			},
		},
		{
			name: "Test_getFeturesInfo test full type of data",
			args: args{`  Sentinel RMS Development Kit 9.1.0.0104 Application Monitor
  Copyright (C) 2016 SafeNet, Inc.

 [Contacting Sentinel RMS Development Kit server on host "999385-pc.samba.gazpromproject.ru"]


 |- Feature Information
   |- Feature name                   : "ARMS_ID"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 
   |- License Version                : 0x08600000
   |- Commuter license allowed       : NO
   |- Unreserved tokens in use       : 0
   |- Public vendor information      : C=GAZPROM_PROJECTIROVANIE;LFID=de5bf;M=nZKXn1G9ndbfrJiWmdy1nJa0nZGZodG2qtq1odK4ourfnZnbnKi7nZy=;
   |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 2D298D42085926FC
     |- Log encryption level           : 2
     |- Check time tamper              : Yes
     |- Server #1 locking code         : Primary   = 2014-*16THL3AVVU6W4MC 
     |- Sharing limit                  : 254
     |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
     |- Allowed on VM                  : YES

 |- Feature Information
   |- Feature name                   : "AVEVA201"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License"
   |- Sharing limit                  : 25
   |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
   |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 28BCBCF9FA5CA671
     |- Sharing limit                  : 25
     |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
     |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 00238CDCDFD2C987
     |- License type                   : "Normal License" 
     |- Allowed on VM                  : YES

   |- Client Information
     |- User name                      : enssap
     |- Host name                      : OENS-1-8222
     |- X display name                 : local
     |- Group name                     : DefaultGrp
     |- Status                         : Running since Tue Nov 03 11:34:40 2020 
     |- Is commuter token              : NO

   |- Client Information
     |- User name                      : akotovskii
     |- Host name                      : CIM-1-8303
     |- X display name                 : local
     |- Group name                     : DefaultGrp
     |- Status                         : Running since Tue Nov 03 11:19:09 2020 
     |- Is commuter token              : NO

   |- Client Information
     |- User name                      : NikishinAV
     |- Host name                      : aso-34
     |- X display name                 : local
     |- Group name                     : DefaultGrp
     |- Status                         : Running since Tue Nov 03 10:06:18 2020 
     |- Is commuter token              : NO
 Press Enter to continue . . .`},
			want: fetures{Features: []feature{
				{
					Feature: map[string]string{
						"Feature name":              "ARMS_ID",
						"Feature version":           "1.0",
						"License type":              "Normal License",
						"License Version":           "0x08600000",
						"Commuter license allowed":  "NO",
						"Unreserved tokens in use":  "0",
						"Public vendor information": "C=GAZPROM_PROJECTIROVANIE;LFID=de5bf;M=nZKXn1G9ndbfrJiWmdy1nJa0nZGZodG2qtq1odK4ourfnZnbnKi7nZy=;",
						"Allowed on VM":             "YES",
					},
					LicenseInformation: []licenseInformation{
						licenseInformation{
							LicenseInformation: map[string]string{
								"License Hash":               "2D298D42085926FC",
								"Log encryption level":       "2",
								"Check time tamper":          "Yes",
								"Server #1 locking code":     "Primary = 2014-*16THL3AVVU6W4MC",
								"Sharing limit":              "254",
								"Token lifetime (heartbeat)": "300 secs (5 min(s))",
								"Allowed on VM":              "YES",
							},
						},
					},
					ClientInformation: nil,
				},
				{
					Feature: map[string]string{
						"Feature name":               "AVEVA201",
						"Feature version":            "1.0",
						"License type":               "Normal License",
						"Sharing limit":              "25",
						"Token lifetime (heartbeat)": "300 secs (5 min(s))",
						"Allowed on VM":              "YES",
					},
					LicenseInformation: []licenseInformation{
						licenseInformation{
							LicenseInformation: map[string]string{
								"License Hash":               "28BCBCF9FA5CA671",
								"Sharing limit":              "25",
								"Token lifetime (heartbeat)": "300 secs (5 min(s))",
								"Allowed on VM":              "YES",
							},
						},
						licenseInformation{
							LicenseInformation: map[string]string{
								"License Hash":  "00238CDCDFD2C987",
								"License type":  "Normal License",
								"Allowed on VM": "YES",
							},
						},
					},
					ClientInformation: []clientInformation{
						clientInformation{
							ClientInformation: map[string]string{
								"User name":         "enssap",
								"Host name":         "OENS-1-8222",
								"X display name":    "local",
								"Group name":        "DefaultGrp",
								"Status":            "Running since Tue Nov 03 11:34:40 2020",
								"Is commuter token": "NO",
							},
						},
						clientInformation{
							ClientInformation: map[string]string{
								"User name":         "akotovskii",
								"Host name":         "CIM-1-8303",
								"X display name":    "local",
								"Group name":        "DefaultGrp",
								"Status":            "Running since Tue Nov 03 11:19:09 2020",
								"Is commuter token": "NO",
							},
						},
						clientInformation{
							ClientInformation: map[string]string{
								"User name":         "NikishinAV",
								"Host name":         "aso-34",
								"X display name":    "local",
								"Group name":        "DefaultGrp",
								"Status":            "Running since Tue Nov 03 10:06:18 2020",
								"Is commuter token": "NO",
							},
						},
					},
				},
			},
			},
		},
		{
			name: "Test_getFeturesInfo test full type of data without cuts",
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
   |- Unreserved tokens in use       : 0
   |- Reserved tokens in use         : 0
   |- Available reserved             : 0
   |- Soft limit on users            : Unlimited
   |- License start date             : Fri Nov 29 00:00:00 2019
   |- Expiration date                : Sat Nov 19 23:59:59 2022
   |- Log encryption level           : 2
   |- Check time tamper              : Yes
   |- Application-server locking     : Server-locked license.
   |- Server #1 locking code         : Primary   = 2014-*16THL3AVVU6W4MC 
   |- Additive/exclusive/aggregate   : Aggregate license(Additive without restrictions).
   |- Sharing criterion              : Vendor defined criteria based sharing.
   |- Sharing limit                  : 254
   |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
   |- Public vendor information      : C=GAZPROM_PROJECTIROVANIE;LFID=de5bf;M=nZKXn1G9ndbfrJiWmdy1nJa0nZGZodG2qtq1odK4ourfnZnbnKi7nZy=;
   |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 2D298D42085926FC
     |- License type                   : "Normal License" 
     |- License Version                : 0x08600000
     |- License storage name           : C:\AVEVA\AVEVA Licensing System\RMS\lservrc_AVEVA
     |- License status                 : Active
     |- Commuter license allowed       : NO
     |- Maximum concurrent user(s)     : 99999
     |- Soft limit on users            : Unlimited
     |- License start date             : Fri Nov 29 00:00:00 2019
     |- Expiration date                : Sat Nov 19 23:59:59 2022
     |- Log encryption level           : 2
     |- Check time tamper              : Yes
     |- Application-server locking     : Server-locked license.
     |- Server #1 locking code         : Primary   = 2014-*16THL3AVVU6W4MC 
     |- Additive/exclusive/aggregate   : Aggregate license(Additive without restrictions).
     |- Sharing criterion              : Vendor defined criteria based sharing.
     |- Sharing limit                  : 254
     |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
     |- Allowed on VM                  : YES

 |- Feature Information
   |- Feature name                   : "AVEVA201"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 
   |- License Version                : 0x08600000
   |- Commuter license allowed       : NO
   |- Maximum concurrent user(s)     : 130
   |- Unreserved tokens in use       : 13
   |- Reserved tokens in use         : 0
   |- Available reserved             : 0
   |- Soft limit on users            : Unlimited
   |- License start date             : Fri Jun 19 00:00:00 2020
   |- Expiration date                : Sat Nov 19 23:59:59 2022
   |- Log encryption level           : 2
   |- Check time tamper              : Yes
   |- Application-server locking     : Server-locked license.
   |- Server #1 locking code         : Primary   = 2014-*16THL3AVVU6W4MC 
   |- Additive/exclusive/aggregate   : Aggregate license(Additive without restrictions).
   |- Sharing criterion              : Vendor defined criteria based sharing.
   |- Sharing limit                  : 25
   |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
   |- Public vendor information      : C=GAZPROM_PROJECTIROVANIE;LFID=de5bf;M=nZKXn1G9reeZm0u1odq4qteZnum1mZzeoeeWmKq5qZq2qKu1mKu7nZy=;PC=V00FN201;PN=QVZFVkEgRXZlcnl0aGluZzNE4oSi;
   |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 28BCBCF9FA5CA671
     |- License type                   : "Normal License" 
     |- License Version                : 0x08600000
     |- License storage name           : C:\AVEVA\AVEVA Licensing System\RMS\lservrc_AVEVA
     |- License status                 : Active
     |- Commuter license allowed       : NO
     |- Maximum concurrent user(s)     : 130
     |- Soft limit on users            : Unlimited
     |- License start date             : Fri Jun 19 00:00:00 2020
     |- Expiration date                : Sat Nov 19 23:59:59 2022
     |- Log encryption level           : 2
     |- Check time tamper              : Yes
     |- Application-server locking     : Server-locked license.
     |- Server #1 locking code         : Primary   = 2014-*16THL3AVVU6W4MC 
     |- Additive/exclusive/aggregate   : Aggregate license(Additive without restrictions).
     |- Sharing criterion              : Vendor defined criteria based sharing.
     |- Sharing limit                  : 25
     |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
     |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 00238CDCDFD2C987
     |- License type                   : "Normal License" 
     |- License Version                : 0x08600000
     |- License storage name           : C:\AVEVA\AVEVA Licensing System\RMS\lservrc_AVEVA
     |- License status                 : Inactive
     |- Commuter license allowed       : NO
     |- Maximum concurrent user(s)     : 5
     |- Soft limit on users            : Unlimited
     |- License start date             : Mon Jun 22 00:00:00 2020
     |- Expiration date                : Tue Jun 23 23:59:59 2020
     |- Log encryption level           : 2
     |- Check time tamper              : Yes
     |- Application-server locking     : Server-locked license.
     |- Server #1 locking code         : Primary   = 2014-*16THL3AVVU6W4MC 
     |- Additive/exclusive/aggregate   : Aggregate license(Additive without restrictions).
     |- Sharing criterion              : Vendor defined criteria based sharing.
     |- Sharing limit                  : 25
     |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
     |- Allowed on VM                  : YES

   |- Client Information
     |- User name                      : enssap
     |- Host name                      : OENS-1-8222
     |- X display name                 : local
     |- Group name                     : DefaultGrp
     |- Status                         : Running since Tue Nov 03 11:34:40 2020 
     |- Is commuter token              : NO

   |- Client Information
     |- User name                      : akotovskii
     |- Host name                      : CIM-1-8303
     |- X display name                 : local
     |- Group name                     : DefaultGrp
     |- Status                         : Running since Tue Nov 03 11:19:09 2020 
     |- Is commuter token              : NO

   |- Client Information
     |- User name                      : NikishinAV
     |- Host name                      : aso-34
     |- X display name                 : local
     |- Group name                     : DefaultGrp
     |- Status                         : Running since Tue Nov 03 10:06:18 2020 
     |- Is commuter token              : NO
 Press Enter to continue . . .`},
			want: fetures{Features: []feature{
				{
					Feature: map[string]string{
						"Feature name":                 "ARMS_ID",
						"Feature version":              "1.0",
						"License type":                 "Normal License",
						"License Version":              "0x08600000",
						"Commuter license allowed":     "NO",
						"Maximum concurrent user(s)":   "99999",
						"Unreserved tokens in use":     "0",
						"Reserved tokens in use":       "0",
						"Available reserved":           "0",
						"Soft limit on users":          "Unlimited",
						"License start date":           "Fri Nov 29 00:00:00 2019",
						"Expiration date":              "Sat Nov 19 23:59:59 2022",
						"Log encryption level":         "2",
						"Check time tamper":            "Yes",
						"Application-server locking":   "Server-locked license.",
						"Server #1 locking code":       "Primary = 2014-*16THL3AVVU6W4MC",
						"Additive/exclusive/aggregate": "Aggregate license(Additive without restrictions).",
						"Sharing criterion":            "Vendor defined criteria based sharing.",
						"Sharing limit":                "254",
						"Token lifetime (heartbeat)":   "300 secs (5 min(s))",
						"Public vendor information":    "C=GAZPROM_PROJECTIROVANIE;LFID=de5bf;M=nZKXn1G9ndbfrJiWmdy1nJa0nZGZodG2qtq1odK4ourfnZnbnKi7nZy=;",
						"Allowed on VM":                "YES",
					},
					LicenseInformation: []licenseInformation{
						licenseInformation{
							LicenseInformation: map[string]string{
								"License Hash":                 "2D298D42085926FC",
								"License type":                 "Normal License",
								"License Version":              "0x08600000",
								"License storage name":         "C:\\AVEVA\\AVEVA Licensing System\\RMS\\lservrc_AVEVA",
								"License status":               "Active",
								"Commuter license allowed":     "NO",
								"Maximum concurrent user(s)":   "99999",
								"Soft limit on users":          "Unlimited",
								"License start date":           "Fri Nov 29 00:00:00 2019",
								"Expiration date":              "Sat Nov 19 23:59:59 2022",
								"Log encryption level":         "2",
								"Check time tamper":            "Yes",
								"Application-server locking":   "Server-locked license.",
								"Server #1 locking code":       "Primary = 2014-*16THL3AVVU6W4MC",
								"Additive/exclusive/aggregate": "Aggregate license(Additive without restrictions).",
								"Sharing criterion":            "Vendor defined criteria based sharing.",
								"Sharing limit":                "254",
								"Token lifetime (heartbeat)":   "300 secs (5 min(s))",
								"Allowed on VM":                "YES",
							},
						},
					},
					ClientInformation: nil,
				},
				{
					Feature: map[string]string{
						"Feature name":                 "AVEVA201",
						"Feature version":              "1.0",
						"License type":                 "Normal License",
						"License Version":              "0x08600000",
						"Commuter license allowed":     "NO",
						"Maximum concurrent user(s)":   "130",
						"Unreserved tokens in use":     "13",
						"Reserved tokens in use":       "0",
						"Available reserved":           "0",
						"Soft limit on users":          "Unlimited",
						"License start date":           "Fri Jun 19 00:00:00 2020",
						"Expiration date":              "Sat Nov 19 23:59:59 2022",
						"Log encryption level":         "2",
						"Check time tamper":            "Yes",
						"Application-server locking":   "Server-locked license.",
						"Server #1 locking code":       "Primary = 2014-*16THL3AVVU6W4MC",
						"Additive/exclusive/aggregate": "Aggregate license(Additive without restrictions).",
						"Sharing criterion":            "Vendor defined criteria based sharing.",
						"Sharing limit":                "25",
						"Token lifetime (heartbeat)":   "300 secs (5 min(s))",
						"Public vendor information":    "C=GAZPROM_PROJECTIROVANIE;LFID=de5bf;M=nZKXn1G9reeZm0u1odq4qteZnum1mZzeoeeWmKq5qZq2qKu1mKu7nZy=;PC=V00FN201;PN=QVZFVkEgRXZlcnl0aGluZzNE4oSi;",
						"Allowed on VM":                "YES",
					},
					LicenseInformation: []licenseInformation{
						licenseInformation{
							LicenseInformation: map[string]string{
								"License Hash":                 "28BCBCF9FA5CA671",
								"License type":                 "Normal License",
								"License Version":              "0x08600000",
								"License storage name":         "C:\\AVEVA\\AVEVA Licensing System\\RMS\\lservrc_AVEVA",
								"License status":               "Active",
								"Commuter license allowed":     "NO",
								"Maximum concurrent user(s)":   "130",
								"Soft limit on users":          "Unlimited",
								"License start date":           "Fri Jun 19 00:00:00 2020",
								"Expiration date":              "Sat Nov 19 23:59:59 2022",
								"Log encryption level":         "2",
								"Check time tamper":            "Yes",
								"Application-server locking":   "Server-locked license.",
								"Server #1 locking code":       "Primary = 2014-*16THL3AVVU6W4MC",
								"Additive/exclusive/aggregate": "Aggregate license(Additive without restrictions).",
								"Sharing criterion":            "Vendor defined criteria based sharing.",
								"Sharing limit":                "25",
								"Token lifetime (heartbeat)":   "300 secs (5 min(s))",
								"Allowed on VM":                "YES",
							},
						},
						licenseInformation{
							LicenseInformation: map[string]string{
								"License Hash":                 "00238CDCDFD2C987",
								"License type":                 "Normal License",
								"License Version":              "0x08600000",
								"License storage name":         "C:\\AVEVA\\AVEVA Licensing System\\RMS\\lservrc_AVEVA",
								"License status":               "Inactive",
								"Commuter license allowed":     "NO",
								"Maximum concurrent user(s)":   "5",
								"Soft limit on users":          "Unlimited",
								"License start date":           "Mon Jun 22 00:00:00 2020",
								"Expiration date":              "Tue Jun 23 23:59:59 2020",
								"Log encryption level":         "2",
								"Check time tamper":            "Yes",
								"Application-server locking":   "Server-locked license.",
								"Server #1 locking code":       "Primary = 2014-*16THL3AVVU6W4MC",
								"Additive/exclusive/aggregate": "Aggregate license(Additive without restrictions).",
								"Sharing criterion":            "Vendor defined criteria based sharing.",
								"Sharing limit":                "25",
								"Token lifetime (heartbeat)":   "300 secs (5 min(s))",
								"Allowed on VM":                "YES",
							},
						},
					},
					ClientInformation: []clientInformation{
						clientInformation{
							ClientInformation: map[string]string{
								"User name":         "enssap",
								"Host name":         "OENS-1-8222",
								"X display name":    "local",
								"Group name":        "DefaultGrp",
								"Status":            "Running since Tue Nov 03 11:34:40 2020",
								"Is commuter token": "NO",
							},
						},
						clientInformation{
							ClientInformation: map[string]string{
								"User name":         "akotovskii",
								"Host name":         "CIM-1-8303",
								"X display name":    "local",
								"Group name":        "DefaultGrp",
								"Status":            "Running since Tue Nov 03 11:19:09 2020",
								"Is commuter token": "NO",
							},
						},
						clientInformation{
							ClientInformation: map[string]string{
								"User name":         "NikishinAV",
								"Host name":         "aso-34",
								"X display name":    "local",
								"Group name":        "DefaultGrp",
								"Status":            "Running since Tue Nov 03 10:06:18 2020",
								"Is commuter token": "NO",
							},
						},
					},
				},
			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFeturesInfo(tt.args.lsmonOut); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFeturesInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitFeature(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name               string
		args               args
		wantFeatreInfo     string
		wantAdditionalInfo string
	}{
		// TODO: Add test cases.
		{
			name: "test one feature",
			args: args{` |- Feature Information
   |- Feature name                   : "ARMS_ID"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 
   |- License Version                : 0x08600000
   |- Commuter license allowed       : NO
   |- Maximum concurrent user(s)     : 99999`},
			wantFeatreInfo: ` |- Feature Information
   |- Feature name                   : "ARMS_ID"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 
   |- License Version                : 0x08600000
   |- Commuter license allowed       : NO
   |- Maximum concurrent user(s)     : 99999`,
			wantAdditionalInfo: "",
		},
		{
			name: "test feature and additional stuff",
			args: args{` |- Feature Information
   |- Feature name                   : "PID-MANAGER"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 

   |- License Information
     |- License Hash                   : 91F4477587CCB838
     |- License type                   : "Normal License" 
     |- License Version                : 0x08600000
     |- License storage name           : C:\AVEVA\AVEVA Licensing System\RMS\lservrc_AVEVA
     |- License status                 : Active
     |- Commuter license allowed       : NO
`},
			wantFeatreInfo: ` |- Feature Information
   |- Feature name                   : "PID-MANAGER"  	
   |- Feature version                : "1.0"

   |- License type                   : "Normal License" 

`,
			wantAdditionalInfo: `   |- License Information
     |- License Hash                   : 91F4477587CCB838
     |- License type                   : "Normal License" 
     |- License Version                : 0x08600000
     |- License storage name           : C:\AVEVA\AVEVA Licensing System\RMS\lservrc_AVEVA
     |- License status                 : Active
     |- Commuter license allowed       : NO
`,
		},
		{
			name: "test feature and additional stuff",
			args: args{` |- Feature Information
   |- Feature name                   : "AVEVA201"  	
   |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 28BCBCF9FA5CA671
     |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 00238CDCDFD2C987
     |- Allowed on VM                  : YES

   |- Client Information
     |- User name                      : enssap
     |- Is commuter token              : NO

   |- Client Information
     |- User name                      : akotovskii
     |- Is commuter token              : NO
`},
			wantFeatreInfo: ` |- Feature Information
   |- Feature name                   : "AVEVA201"  	
   |- Allowed on VM                  : YES

`,
			wantAdditionalInfo: `   |- License Information
     |- License Hash                   : 28BCBCF9FA5CA671
     |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 00238CDCDFD2C987
     |- Allowed on VM                  : YES

   |- Client Information
     |- User name                      : enssap
     |- Is commuter token              : NO

   |- Client Information
     |- User name                      : akotovskii
     |- Is commuter token              : NO
`,
		},
		{
			name: "test feature and additional stuff mixed",
			args: args{` |- Feature Information
   |- Feature name                   : "AVEVA201"  	
   |- Allowed on VM                  : YES

   |- Client Information
     |- User name                      : enssap
     |- Is commuter token              : NO

   |- License Information
     |- License Hash                   : 28BCBCF9FA5CA671
     |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 00238CDCDFD2C987
     |- Allowed on VM                  : YES

   |- Client Information
     |- User name                      : akotovskii
     |- Is commuter token              : NO
`},
			wantFeatreInfo: ` |- Feature Information
   |- Feature name                   : "AVEVA201"  	
   |- Allowed on VM                  : YES

`,
			wantAdditionalInfo: `   |- Client Information
     |- User name                      : enssap
     |- Is commuter token              : NO

   |- License Information
     |- License Hash                   : 28BCBCF9FA5CA671
     |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 00238CDCDFD2C987
     |- Allowed on VM                  : YES

   |- Client Information
     |- User name                      : akotovskii
     |- Is commuter token              : NO
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFeatreInfo, gotAdditionalInfo := splitFeature(tt.args.s)
			if gotFeatreInfo != tt.wantFeatreInfo {
				t.Errorf("splitFeature() gotFeatreInfo = %v, want %v", gotFeatreInfo, tt.wantFeatreInfo)
			}
			if gotAdditionalInfo != tt.wantAdditionalInfo {
				t.Errorf("splitFeature() gotAdditionalInfo = %v, want %v", gotAdditionalInfo, tt.wantAdditionalInfo)
			}
		})
	}
}

func Test_getLicenseInformation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []licenseInformation
	}{
		// TODO: Add test cases.
		{
			name: "get licenses 1",
			args: args{`
   |- License Information
     |- License Hash                   : 28BCBCF9FA5CA671
     |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 00238CDCDFD2C987
     |- Allowed on VM                  : YES

   |- Client Information
     |- User name                      : enssap
     |- Is commuter token              : NO

   |- Client Information
     |- User name                      : akotovskii
     |- Is commuter token              : NO
`},
			want: []licenseInformation{
				licenseInformation{map[string]string{
					"License Hash":  "28BCBCF9FA5CA671",
					"Allowed on VM": "YES",
				}},
				licenseInformation{map[string]string{
					"License Hash":  "00238CDCDFD2C987",
					"Allowed on VM": "YES",
				}},
			},
		},
		{
			name: "get licenses 2",
			args: args{`   |- Client Information
     |- User name                      : enssap
     |- Is commuter token              : NO

   |- License Information
     |- License Hash                   : 28BCBCF9FA5CA671
     |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 00238CDCDFD2C987
     |- Allowed on VM                  : YES

   |- Client Information
     |- User name                      : akotovskii
     |- Is commuter token              : NO
`},
			want: []licenseInformation{
				licenseInformation{map[string]string{
					"License Hash":  "28BCBCF9FA5CA671",
					"Allowed on VM": "YES",
				}},
				licenseInformation{map[string]string{
					"License Hash":  "00238CDCDFD2C987",
					"Allowed on VM": "YES",
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLicenseInformation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLicenseInformation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getClientInformation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []clientInformation
	}{
		// TODO: Add test cases.
		{
			name: "test clients info 1",
			args: args{`
   |- Client Information
     |- User name                      : enssap
     |- Host name                      : OENS-1-8222
     |- Is commuter token              : NO

   |- License Information
     |- License Hash                   : 28BCBCF9FA5CA671
     |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 00238CDCDFD2C987
     |- Allowed on VM                  : YES

   |- Client Information
     |- User name                      : akotovskii
     |- Is commuter token              : NO`},
			want: []clientInformation{
				{map[string]string{
					"User name":         "enssap",
					"Host name":         "OENS-1-8222",
					"Is commuter token": "NO",
				}},
				{map[string]string{
					"User name":         "akotovskii",
					"Is commuter token": "NO",
				}},
			},
		},
		{
			name: "test clients info 1",
			args: args{`
   |- License Information
     |- License Hash                   : 28BCBCF9FA5CA671
     |- License type                   : "Normal License" 
     |- License Version                : 0x08600000
     |- License storage name           : C:\AVEVA\AVEVA Licensing System\RMS\lservrc_AVEVA
     |- License status                 : Active
     |- Commuter license allowed       : NO
     |- Maximum concurrent user(s)     : 130
     |- Soft limit on users            : Unlimited
     |- License start date             : Fri Jun 19 00:00:00 2020
     |- Expiration date                : Sat Nov 19 23:59:59 2022
     |- Log encryption level           : 2
     |- Check time tamper              : Yes
     |- Application-server locking     : Server-locked license.
     |- Server #1 locking code         : Primary   = 2014-*16THL3AVVU6W4MC 
     |- Additive/exclusive/aggregate   : Aggregate license(Additive without restrictions).
     |- Sharing criterion              : Vendor defined criteria based sharing.
     |- Sharing limit                  : 25
     |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
     |- Allowed on VM                  : YES

   |- License Information
     |- License Hash                   : 00238CDCDFD2C987
     |- License type                   : "Normal License" 
     |- License Version                : 0x08600000
     |- License storage name           : C:\AVEVA\AVEVA Licensing System\RMS\lservrc_AVEVA
     |- License status                 : Inactive
     |- Commuter license allowed       : NO
     |- Maximum concurrent user(s)     : 5
     |- Soft limit on users            : Unlimited
     |- License start date             : Mon Jun 22 00:00:00 2020
     |- Expiration date                : Tue Jun 23 23:59:59 2020
     |- Log encryption level           : 2
     |- Check time tamper              : Yes
     |- Application-server locking     : Server-locked license.
     |- Server #1 locking code         : Primary   = 2014-*16THL3AVVU6W4MC 
     |- Additive/exclusive/aggregate   : Aggregate license(Additive without restrictions).
     |- Sharing criterion              : Vendor defined criteria based sharing.
     |- Sharing limit                  : 25
     |- Token lifetime (heartbeat)     : 300 secs (5 min(s))
     |- Allowed on VM                  : YES

   |- Client Information
     |- User name                      : enssap
     |- Host name                      : OENS-1-8222
     |- X display name                 : local
     |- Group name                     : DefaultGrp
     |- Status                         : Running since Tue Nov 03 11:34:40 2020 
     |- Is commuter token              : NO

   |- Client Information
     |- User name                      : akotovskii
     |- Host name                      : CIM-1-8303
     |- X display name                 : local
     |- Group name                     : DefaultGrp
     |- Status                         : Running since Tue Nov 03 11:19:09 2020 
     |- Is commuter token              : NO
`},
			want: []clientInformation{
				{map[string]string{
					"User name":         "enssap",
					"Host name":         "OENS-1-8222",
					"X display name":    "local",
					"Group name":        "DefaultGrp",
					"Status":            "Running since Tue Nov 03 11:34:40 2020",
					"Is commuter token": "NO",
				}},
				{map[string]string{
					"User name":         "akotovskii",
					"Host name":         "CIM-1-8303",
					"X display name":    "local",
					"Group name":        "DefaultGrp",
					"Status":            "Running since Tue Nov 03 11:19:09 2020",
					"Is commuter token": "NO",
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getClientInformation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClientInformation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createJSON(t *testing.T) {
	type args struct {
		s fetures
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test border trim",
			args: args{
				fetures{Features: []feature{
					{
						Feature: map[string]string{
							"Feature name":               "ARMS_ID",
							"Feature version":            "1.0",
							"License Version":            "0x08600000",
							"License type":               "Normal License",
							"Commuter license allowed":   "NO",
							"Maximum concurrent user(s)": "99999",
						},
						LicenseInformation: nil,
						ClientInformation:  nil,
					},
				},
				},
			},
			want:    "{\"Features\":[{\"Feature\":{\"Commuter license allowed\":\"NO\",\"Feature name\":\"ARMS_ID\",\"Feature version\":\"1.0\",\"License Version\":\"0x08600000\",\"License type\":\"Normal License\",\"Maximum concurrent user(s)\":\"99999\"},\"LicenseInformation\":null,\"ClientInformation\":null}]}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createJSON(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("createJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
