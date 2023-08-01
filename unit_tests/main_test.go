package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result, err := Order("шаурма", 3)
		if result != "Замовлення = шаурма, кількість = 3" {
			t.Fail()
		}
		if err != nil {
			t.Fail()
		}
	})

	t.Run("does_not_exist", func(t *testing.T) {
		result, err := Order("круасон", 3)
		if result != "" {
			t.Fail()
		}
		if err == nil {
			t.Fail()
		}
	})
}

func TestOrder1(t *testing.T) {
	type args struct {
		dish   string
		amount uint
	}
	tests := map[string]struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		"success": {
			args: args{
				dish:   "шаурма",
				amount: 3,
			},
			want:    "Замовлення = шаурма, кількість = 3",
			wantErr: nil,
		},
		"does not exist": {
			args: args{
				dish:   "круасон",
				amount: 3,
			},
			want:    "",
			wantErr: fmt.Errorf("блюдо не знайдено"),
		},
	}
	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {
			got, err := Order(tt.args.dish, tt.args.amount)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
