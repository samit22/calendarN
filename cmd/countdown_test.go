package cmd

import (
	"os"
	"strings"
	"testing"
	"time"
)

func Test_runCountdown(t *testing.T) {
	type args struct {
		args []string
		run  int64
		save bool
	}
	newT := time.Now().Add(50 * time.Hour)

	tests := []struct {
		name        string
		args        args
		wantCdTimes int
		wantErr     bool
	}{
		{
			name: "empty arguments",
			args: args{
				args: nil,
			},
			wantErr: true,
		},
		{
			name: "empty date",
			args: args{
				args: []string{""},
			},
			wantErr: true,
		},
		{
			name: "failed date",
			args: args{
				args: []string{"20"},
			},
			wantErr: true,
		},
		{
			name: "empty time",
			args: args{
				args: []string{newT.Format("2006-01-02")},
				run:  1,
			},
			wantErr: false,
		},
		{
			name: "Date and time given",
			args: args{
				args: []string{newT.Format("2006-01-02"), "10:00:03"},
				run:  1,
			},
			wantErr: false,
		}, {
			name: "Date and time and run default time",
			args: args{
				args: []string{newT.Format("2006-01-02"), "10:00:03"},
				run:  5,
			},
			wantErr: false,
		},
		{
			name: "Date and time and run given",
			args: args{
				args: []string{newT.Format("2006-01-02"), "10:00:03"},
				run:  3,
			},
			wantErr: false,
		},
		{
			name: "Date and time and run and save given",
			args: args{
				args: []string{newT.Format("2006-01-02"), "10:00:03"},
				run:  3,
				save: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			run = tt.args.run
			save = tt.args.save
			if save {
				dir := t.TempDir()
				filePath = dir + "/" + fileName
			}

			gotCdTimes, err := runCountdown(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("runCountdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotCdTimes != int(tt.args.run) {
				t.Errorf("expected to have printed countdown %d times got %d", tt.args.run, gotCdTimes)
			}

		})
	}
}

// Test for loadDataToFile overwrite fix
func Test_loadDataToFile_OverwriteBehavior(t *testing.T) {
	futureDate := time.Now().Add(100 * time.Hour).Format("2006-01-02")

	tests := []struct {
		name           string
		existingData   string
		newName        string
		newDate        string
		overwriteFlag  bool
		wantErr        bool
		errContains    string
		wantFileData   string
	}{
		{
			name:          "save new countdown - no existing data",
			existingData:  "",
			newName:       "test1",
			newDate:       futureDate,
			overwriteFlag: false,
			wantErr:       false,
			wantFileData:  "test1 :: " + futureDate,
		},
		{
			name:          "save new countdown - with existing different name",
			existingData:  "existing :: 2099-01-01\n",
			newName:       "test2",
			newDate:       futureDate,
			overwriteFlag: false,
			wantErr:       false,
			wantFileData:  "test2 :: " + futureDate,
		},
		{
			name:          "duplicate name without overwrite - should error",
			existingData:  "duplicate :: 2099-01-01\n",
			newName:       "duplicate",
			newDate:       futureDate,
			overwriteFlag: false,
			wantErr:       true,
			errContains:   "already exists",
		},
		{
			name:          "duplicate name with overwrite - should update",
			existingData:  "overwriteme :: 2099-01-01\n",
			newName:       "overwriteme",
			newDate:       futureDate,
			overwriteFlag: true,
			wantErr:       false,
			wantFileData:  "overwriteme :: " + futureDate,
		},
		{
			name:          "overwrite preserves other entries",
			existingData:  "keep :: 2099-02-02\noverwriteme :: 2099-01-01\n",
			newName:       "overwriteme",
			newDate:       futureDate,
			overwriteFlag: true,
			wantErr:       false,
			wantFileData:  futureDate, // Just check the new date is there
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup temp directory and file
			dir := t.TempDir()
			testFilePath := dir + "/" + fileName

			// Set global variables for the test
			originalFilePath := filePath
			originalOverwrite := overwrite
			filePath = testFilePath
			overwrite = tt.overwriteFlag
			defer func() {
				filePath = originalFilePath
				overwrite = originalOverwrite
			}()

			// Write existing data if any
			if tt.existingData != "" {
				err := os.WriteFile(testFilePath, []byte(tt.existingData), 0644)
				if err != nil {
					t.Fatalf("Failed to setup test file: %v", err)
				}
			}

			// Run the function
			err := loadDataToFile(tt.newName, tt.newDate)

			// Check error expectation
			if tt.wantErr {
				if err == nil {
					t.Errorf("loadDataToFile() expected error but got none")
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("loadDataToFile() error = %q, want error containing %q", err.Error(), tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("loadDataToFile() unexpected error: %v", err)
				return
			}

			// Verify file contents
			data, err := os.ReadFile(testFilePath)
			if err != nil {
				t.Errorf("Failed to read test file: %v", err)
				return
			}

			if tt.wantFileData != "" && !strings.Contains(string(data), tt.wantFileData) {
				t.Errorf("File contents = %q, want to contain %q", string(data), tt.wantFileData)
			}
		})
	}
}

// Test that overwrite actually replaces and doesn't duplicate
func Test_loadDataToFile_OverwriteNoDuplication(t *testing.T) {
	dir := t.TempDir()
	testFilePath := dir + "/" + fileName

	// Set global variables
	originalFilePath := filePath
	originalOverwrite := overwrite
	filePath = testFilePath
	defer func() {
		filePath = originalFilePath
		overwrite = originalOverwrite
	}()

	futureDate1 := time.Now().Add(100 * time.Hour).Format("2006-01-02")
	futureDate2 := time.Now().Add(200 * time.Hour).Format("2006-01-02")

	// First save
	overwrite = false
	err := loadDataToFile("testname", futureDate1)
	if err != nil {
		t.Fatalf("First save failed: %v", err)
	}

	// Overwrite with new date
	overwrite = true
	err = loadDataToFile("testname", futureDate2)
	if err != nil {
		t.Fatalf("Overwrite failed: %v", err)
	}

	// Read and verify
	data, err := os.ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	content := string(data)
	
	// Should contain new date
	if !strings.Contains(content, futureDate2) {
		t.Errorf("File should contain new date %s, got: %s", futureDate2, content)
	}

	// Should NOT contain old date (it was overwritten)
	if strings.Contains(content, futureDate1) {
		t.Errorf("File should NOT contain old date %s after overwrite, got: %s", futureDate1, content)
	}

	// Should have only one entry for "testname"
	count := strings.Count(content, "testname")
	if count != 1 {
		t.Errorf("File should have exactly 1 entry for 'testname', found %d occurrences: %s", count, content)
	}
}
