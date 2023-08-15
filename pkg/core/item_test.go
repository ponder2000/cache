package core

import (
	"reflect"
	"testing"
	"time"
)

func TestItem_GetData(t *testing.T) {
	type fields struct {
		data               any
		createdOn          time.Time
		expiredOn          time.Time
		expirationDuration time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   any
	}{
		{
			name:   "item1",
			fields: fields{data: "some data 1"},
			want:   "some data 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Item{
				data:               tt.fields.data,
				createdOn:          tt.fields.createdOn,
				expiredOn:          tt.fields.expiredOn,
				expirationDuration: tt.fields.expirationDuration,
			}
			if got := i.GetData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItem_IsExpired(t *testing.T) {
	now := time.Now()

	type fields struct {
		data               any
		createdOn          time.Time
		expiredOn          time.Time
		expirationDuration time.Duration
	}
	type args struct {
		expirationTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "1", fields: fields{expiredOn: now.Add(4 * time.Second)}, args: struct{ expirationTime time.Time }{expirationTime: now.Add(5 * time.Second)}, want: true},
		{name: "2", fields: fields{expiredOn: now.Add(10 * time.Second)}, args: struct{ expirationTime time.Time }{expirationTime: now.Add(5 * time.Second)}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Item{
				data:               tt.fields.data,
				createdOn:          tt.fields.createdOn,
				expiredOn:          tt.fields.expiredOn,
				expirationDuration: tt.fields.expirationDuration,
			}
			if got := i.IsExpired(tt.args.expirationTime); got != tt.want {
				t.Errorf("IsExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}
