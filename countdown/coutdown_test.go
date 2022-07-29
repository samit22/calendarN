package countdown

import (
	"testing"
	"time"
)

type mockT struct {
	t   string
	loc *time.Location
}

func (m *mockT) GetCurrentTime() time.Time {
	t, _ := time.ParseInLocation(acceptFormatDateTime, m.t, m.loc)
	return t
}
func (m *mockT) SetLocation(loc *time.Location) {
	m.loc = loc
}
func Test_new_GetEnglishCountdown(t *testing.T) {
	type fields struct {
		currentTime CurrentTimeIface
	}
	type args struct {
		d  string
		t  string
		tz string
	}
	mockCT := &mockT{
		t:   "2022-01-02 00:00:00",
		loc: time.Local,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Invalid date",
			fields: fields{
				currentTime: mockCT,
			},
			args: args{
				d: "2021",
			},
			want:    &Response{},
			wantErr: true,
		},
		{
			name: "Count down past",
			fields: fields{
				currentTime: mockCT,
			},
			args: args{
				d: "2020-01-00",
				t: "00:00:00",
			},
			want:    &Response{},
			wantErr: true,
		},
		{
			name: "Future",
			fields: fields{
				currentTime: mockCT,
			},
			args: args{
				d: "2022-01-05",
				t: "01:02:03",
			},
			want: &Response{
				Days:    3,
				Hours:   1,
				Minutes: 2,
				Seconds: 3,
			},
			wantErr: false,
		},
		{
			name: "Bad location is provided",
			fields: fields{
				currentTime: mockCT,
			},
			args: args{
				d:  "2022-01-05",
				t:  "01:02:03",
				tz: "Asia",
			},
			want: &Response{
				Days:    3,
				Hours:   1,
				Minutes: 2,
				Seconds: 3,
			},
			wantErr: true,
		},
		{
			name: "Location is provided",
			fields: fields{
				currentTime: mockCT,
			},
			args: args{
				d:  "2022-01-05",
				t:  "01:02:03",
				tz: "UTC",
			},
			want: &Response{
				Days:    3,
				Hours:   1,
				Minutes: 2,
				Seconds: 3,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &new{
				currentTime: tt.fields.currentTime,
			}
			got, err := n.GetEnglishCountdown(tt.args.d, tt.args.t, tt.args.tz)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error %v", err)
				return
			}

			if got.Days != tt.want.Days {
				t.Errorf("test case: %s, expected days %d got %d", tt.name, tt.want.Days, got.Days)
			}
			if got.Hours != tt.want.Hours {
				t.Errorf("test case: %s, expected hours %d got %d", tt.name, tt.want.Hours, got.Hours)
			}
			if got.Minutes != tt.want.Minutes {
				t.Errorf("test case: %s, expected minutes %d got %d", tt.name, tt.want.Minutes, got.Minutes)
			}
			if got.Seconds != tt.want.Seconds {
				t.Errorf("test case: %s, expected minutes %d got %d", tt.name, tt.want.Seconds, got.Seconds)
			}
		})

	}
}

func TestNewCountdown(t *testing.T) {
	tests := []struct {
		name string
		want *new
	}{
		{
			want: &new{
				currentTime: &tm{
					location: time.Local,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCountdown()
			if got.currentTime.GetCurrentTime().Format("2006-01-02 16:04:05") != tt.want.currentTime.GetCurrentTime().Format("2006-01-02 16:04:05") {
				t.Errorf("NewCountdown() = %v, want %v", got, tt.want)
				return
			}

			got.currentTime.GetCurrentTime()
		})
	}
}
