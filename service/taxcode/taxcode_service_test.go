//go:build unit

package taxcode

import (
	"cloud.google.com/go/civil"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"taxcode-converter/service"
	"testing"
)

// TODO to fix all the tests
func TestService_CalculatePersonData(t *testing.T) {
	type fields struct {
		validator validator.Validate
	}
	type args struct {
		req service.CalculatePersonDataRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *service.CalculatePersonDataResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "",
			fields:  fields{validator: CreateValidator()},
			args:    args{req: service.CalculatePersonDataRequest{TaxCode: ""}},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTaxCodeService(tt.fields.validator)
			got, err := s.CalculatePersonData(tt.args.req)
			if !tt.wantErr(t, err, fmt.Sprintf("CalculatePersonData(%v)", tt.args.req)) {
				return
			}
			assert.Equalf(t, tt.want, got, "CalculatePersonData(%v)", tt.args.req)
		})
	}
}

func TestService_CalculateTaxCode(t *testing.T) {
	type fields struct {
		validator validator.Validate
	}
	type args struct {
		req service.CalculateTaxCodeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *service.CalculateTaxCodeResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "",
			fields: fields{validator: CreateValidator()},
			args: args{req: service.CalculateTaxCodeRequest{
				Gender:      "",
				Name:        "",
				Surname:     "",
				DateOfBirth: civil.Date{},
				BirthPlace:  "",
				Province:    "",
			}},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTaxCodeService(tt.fields.validator)
			got, err := s.CalculateTaxCode(tt.args.req)
			if !tt.wantErr(t, err, fmt.Sprintf("CalculateTaxCode(%v)", tt.args.req)) {
				return
			}
			assert.Equalf(t, tt.want, got, "CalculateTaxCode(%v)", tt.args.req)
		})
	}
}
