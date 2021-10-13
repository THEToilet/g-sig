package application

import (
	"errors"
	"g-sig/pkg/domain/model"
	"g-sig/pkg/gateway/repository"
	"github.com/rs/zerolog"
	"os"
	"reflect"
	"testing"
)

func TestSignalingUseCase_Register(t *testing.T) {
	type fields struct {
		userRepository     repository.UserRepository
		userInfoRepository repository.UserInfoRepository
		logger             zerolog.Logger
	}
	type args struct {
		userID      string
		geoLocation model.GeoLocation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "register_success",
			fields: fields{
				userRepository:     *repository.NewUserRepository(),
				userInfoRepository: *repository.NewUserInfoRepository(),
				logger:             zerolog.New(os.Stdout).With().Timestamp().Logger(),
			},
			args: args{
				userID: "1234",
				geoLocation: model.GeoLocation{
					Latitude:  0,
					Longitude: 0,
				},
			},
			wantErr: nil,
		},
		{
			name: "register_error",
			fields: fields{
				userRepository:     *repository.NewUserRepository(),
				userInfoRepository: *repository.NewUserInfoRepository(),
				logger:             zerolog.New(os.Stdout).With().Timestamp().Logger(),
			},
			args: args{
				userID: "1234",
				geoLocation: model.GeoLocation{
					Latitude:  0,
					Longitude: 0,
				},
			},
			wantErr: model.ErrUserAlreadyExisted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SignalingUseCase{
				userRepository:     tt.fields.userRepository,
				userInfoRepository: tt.fields.userInfoRepository,
				logger:             &tt.fields.logger,
			}
			tt.fields.userInfoRepository.Save(model.UserInfo{
				UserID:    "1234",
				Latitude:  0,
				Longitude: 0,
			})
			err := s.Register(tt.args.userID, tt.args.geoLocation)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSignalingUseCase_Update(t *testing.T) {
	type fields struct {
		userRepository     repository.UserRepository
		userInfoRepository repository.UserInfoRepository
		logger             zerolog.Logger
	}
	type args struct {
		userInfo model.UserInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "update_success",
			fields: fields{
				userRepository:     *repository.NewUserRepository(),
				userInfoRepository: *repository.NewUserInfoRepository(),
				logger:             zerolog.New(os.Stdout).With().Timestamp().Logger(),
			},
			args: args{
				userInfo: model.UserInfo{
					UserID:    "",
					Latitude:  0,
					Longitude: 0,
				},
			},
			wantErr: nil,
		},
		{
			name: "update_error",
			fields: fields{
				userRepository:     *repository.NewUserRepository(),
				userInfoRepository: *repository.NewUserInfoRepository(),
				logger:             zerolog.New(os.Stdout).With().Timestamp().Logger(),
			},
			args: args{
				userInfo: model.UserInfo{
					UserID:    "",
					Latitude:  0,
					Longitude: 0,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SignalingUseCase{
				userRepository:     tt.fields.userRepository,
				userInfoRepository: tt.fields.userInfoRepository,
				logger:             &tt.fields.logger,
			}
			err := s.Update(tt.args.userInfo)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignalingUseCase_Delete(t *testing.T) {
	type fields struct {
		userRepository     repository.UserRepository
		userInfoRepository repository.UserInfoRepository
		logger             zerolog.Logger
	}
	type args struct {
		userInfo model.UserInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "delete_success",
			fields: fields{
				userRepository:     *repository.NewUserRepository(),
				userInfoRepository: *repository.NewUserInfoRepository(),
				logger:             zerolog.New(os.Stdout).With().Timestamp().Logger(),
			},
			args: args{
				userInfo: model.UserInfo{
					UserID:    "",
					Latitude:  0,
					Longitude: 0,
				},
			},
			wantErr: nil,
		},
		{
			name: "delete_error",
			fields: fields{
				userRepository:     *repository.NewUserRepository(),
				userInfoRepository: *repository.NewUserInfoRepository(),
				logger:             zerolog.New(os.Stdout).With().Timestamp().Logger(),
			},
			args: args{
				userInfo: model.UserInfo{
					UserID:    "",
					Latitude:  0,
					Longitude: 0,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SignalingUseCase{
				userRepository:     tt.fields.userRepository,
				userInfoRepository: tt.fields.userInfoRepository,
				logger:             &tt.fields.logger,
			}
			err := s.Delete(tt.args.userInfo)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignalingUseCase_StaticSearch(t *testing.T) {
	type fields struct {
		userRepository     repository.UserRepository
		userInfoRepository repository.UserInfoRepository
		logger             zerolog.Logger
	}
	type args struct {
		userInfo       model.UserInfo
		searchDistance float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*model.UserInfo
	}{
		{
			name: "staticSearch_success",
			fields: fields{
				userRepository:     *repository.NewUserRepository(),
				userInfoRepository: *repository.NewUserInfoRepository(),
				logger:             zerolog.New(os.Stdout).With().Timestamp().Logger(),
			},
			args: args{
				userInfo: model.UserInfo{
					UserID:    "",
					Latitude:  0,
					Longitude: 0,
				},
				searchDistance: 0,
			},
			want: nil,
		},
		{
			name: "staticSearch_error",
			fields: fields{
				userRepository:     *repository.NewUserRepository(),
				userInfoRepository: *repository.NewUserInfoRepository(),
				logger:             zerolog.New(os.Stdout).With().Timestamp().Logger(),
			},
			args: args{
				userInfo: model.UserInfo{
					UserID:    "",
					Latitude:  0,
					Longitude: 0,
				},
				searchDistance: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SignalingUseCase{
				userRepository:     tt.fields.userRepository,
				userInfoRepository: tt.fields.userInfoRepository,
				logger:             &tt.fields.logger,
			}
			got := s.StaticSearch(tt.args.userInfo, tt.args.searchDistance)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StaticSearch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignalingUseCase_DynamicSearch(t *testing.T) {
	type fields struct {
		userRepository     repository.UserRepository
		userInfoRepository repository.UserInfoRepository
		logger             zerolog.Logger
	}
	type args struct {
		userInfo       model.UserInfo
		searchDistance float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*model.UserInfo
	}{
		{
			name: "dynamicSearch_success",
			fields: fields{
				userRepository:     *repository.NewUserRepository(),
				userInfoRepository: *repository.NewUserInfoRepository(),
				logger:             zerolog.New(os.Stdout).With().Timestamp().Logger(),
			},
			args: args{
				userInfo: model.UserInfo{
					UserID:    "",
					Latitude:  0,
					Longitude: 0,
				},
				searchDistance: 0,
			},
			want: nil,
		},
		{
			name: "dynamicSearch_error",
			fields: fields{
				userRepository:     *repository.NewUserRepository(),
				userInfoRepository: *repository.NewUserInfoRepository(),
				logger:             zerolog.New(os.Stdout).With().Timestamp().Logger(),
			},
			args: args{
				userInfo: model.UserInfo{
					UserID:    "",
					Latitude:  0,
					Longitude: 0,
				},
				searchDistance: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SignalingUseCase{
				userRepository:     tt.fields.userRepository,
				userInfoRepository: tt.fields.userInfoRepository,
				logger:             &tt.fields.logger,
			}
			got := s.DynamicSearch(tt.args.userInfo, tt.args.searchDistance)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DynamicSearch() got = %v, want %v", got, tt.want)
			}
		})
	}
}
