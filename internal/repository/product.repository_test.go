package repository

import (
	"testing"

	"github.com/bagastri07/boilerplate-service/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewProductRepository(t *testing.T) {
	type args struct {
		DB *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want model.ProductRepository
	}{
		{
			name: "normal",
			args: args{
				DB: &gorm.DB{},
			},
			want: &productRepository{
				DB: &gorm.DB{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewProductRepository(tt.args.DB)
			assert.Equal(t, tt.want, got)
		})
	}
}
