package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockLogger struct {}

func (m *MockLogger) Printf(format string, v ...interface{}) {}
func (m *MockLogger) Println(v ...interface{}) {}

func TestPackageService_GetPackageSize(t *testing.T) {
	packageSizes := []int{250,500,1000,2000,5000}

	tests := []struct {
		name string
		target int
		packSizes []int
		want map[int]int
	}{
		{ 
			name: "test 1",
			target: 1,
			packSizes: packageSizes,
			want: map[int]int{250: 1},
		},
		{
			name: "test 250",
			target: 250,
			packSizes: packageSizes,
			want: map[int]int{250: 1},
		},
		{
			name: "test 251",
			target: 251,
			packSizes: packageSizes,
			want: map[int]int{500: 1},
		},
		{
			name: "test 500",
			target: 500,
			packSizes: packageSizes,
			want: map[int]int{500: 1},
		},
		{
			name: "test 501",
			target: 501,
			packSizes: packageSizes,
			want: map[int]int{500: 1, 250: 1},
		},
		{
			name: "test 12001",
			target: 12001,
			packSizes: packageSizes,
			want: map[int]int{250: 1, 2000: 1, 5000: 2},
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &MockLogger{}
			packageService := NewPackageService(tt.packSizes, logger)

			got := packageService.GetPackageSize(context.Background(), tt.target)
			assert.Equal(t, got, tt.want)
		})
	}
}