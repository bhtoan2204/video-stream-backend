package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	repository_interface "github.com/bhtoan2204/user/internal/domain/repository/command"
)

func TestNewRefreshTokenService(t *testing.T) {
	type args struct {
		refreshTokenRepository repository_interface.RefreshTokenRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want *RefreshTokenService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRefreshTokenService(tt.args.refreshTokenRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRefreshTokenService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefreshTokenService_CreateRefreshToken(t *testing.T) {
	type fields struct {
		refreshTokenRepository repository_interface.RefreshTokenRepositoryInterface
	}
	type args struct {
		ctx        context.Context
		token      string
		userId     string
		expires_at time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RefreshTokenService{
				refreshTokenRepository: tt.fields.refreshTokenRepository,
			}
			if err := s.CreateRefreshToken(tt.args.ctx, tt.args.token, tt.args.userId, tt.args.expires_at); (err != nil) != tt.wantErr {
				t.Errorf("RefreshTokenService.CreateRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRefreshTokenService_HardDeleteByQuery(t *testing.T) {
	type fields struct {
		refreshTokenRepository repository_interface.RefreshTokenRepositoryInterface
	}
	type args struct {
		ctx   context.Context
		query map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RefreshTokenService{
				refreshTokenRepository: tt.fields.refreshTokenRepository,
			}
			if err := s.HardDeleteByQuery(tt.args.ctx, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("RefreshTokenService.HardDeleteByQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRefreshTokenService_RevokedByQuery(t *testing.T) {
	type fields struct {
		refreshTokenRepository repository_interface.RefreshTokenRepositoryInterface
	}
	type args struct {
		ctx   context.Context
		query map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RefreshTokenService{
				refreshTokenRepository: tt.fields.refreshTokenRepository,
			}
			if err := s.RevokedByQuery(tt.args.ctx, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("RefreshTokenService.RevokedByQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRefreshTokenService_DeleteByQuery(t *testing.T) {
	type fields struct {
		refreshTokenRepository repository_interface.RefreshTokenRepositoryInterface
	}
	type args struct {
		ctx   context.Context
		query map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RefreshTokenService{
				refreshTokenRepository: tt.fields.refreshTokenRepository,
			}
			if err := s.DeleteByQuery(tt.args.ctx, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("RefreshTokenService.DeleteByQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRefreshTokenService_CheckRefreshToken(t *testing.T) {
	type fields struct {
		refreshTokenRepository repository_interface.RefreshTokenRepositoryInterface
	}
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RefreshTokenService{
				refreshTokenRepository: tt.fields.refreshTokenRepository,
			}
			got, err := s.CheckRefreshToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshTokenService.CheckRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RefreshTokenService.CheckRefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
