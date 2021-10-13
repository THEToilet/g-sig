package application

import (
	"g-sig/pkg/domain/model"
	"g-sig/pkg/domain/repository"
	"github.com/rs/zerolog"
	"reflect"
	"testing"
)

/*
func getSignalingUseCase() *SignalingUseCase {
	file, err := os.Open("../../../config.conf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	config := config.NewConfig(buffer)
	fmt.Println(config)

	logger, err := logger2.NewLogger(config)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info().Str("Title", config.Title).Msg("Config")
	logger.Info().Str("LogLevel", config.LogInfo.Level).Msg("Config")
	// Repository
	userRepository := repository.NewUserRepository(logger)
	userInfoRepository := repository.NewUserInfoRepository(logger)

	// UseCase
	signalingUseCase := NewSignalingUseCase(userRepository, userInfoRepository, logger)

	return signalingUseCase
}
 */

func TestSignalingUseCase_Register(t *testing.T) {
	type fields struct {
		userRepository     repository.UserRepository
		userInfoRepository repository.UserInfoRepository
		logger             *zerolog.Logger
	}
	type args struct {
		userInfo model.UserInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SignalingUseCase{
				userRepository:     tt.fields.userRepository,
				userInfoRepository: tt.fields.userInfoRepository,
				logger:             tt.fields.logger,
			}
			got, err := s.Register(tt.args.userInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignalingUseCase_Update(t *testing.T) {
	type fields struct {
		userRepository     repository.UserRepository
		userInfoRepository repository.UserInfoRepository
		logger             *zerolog.Logger
	}
	type args struct {
		userInfo model.UserInfo
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
			s := &SignalingUseCase{
				userRepository:     tt.fields.userRepository,
				userInfoRepository: tt.fields.userInfoRepository,
				logger:             tt.fields.logger,
			}
			if err := s.Update(tt.args.userInfo); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignalingUseCase_Delete(t *testing.T) {
	type fields struct {
		userRepository     repository.UserRepository
		userInfoRepository repository.UserInfoRepository
		logger             *zerolog.Logger
	}
	type args struct {
		userInfo model.UserInfo
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
			s := &SignalingUseCase{
				userRepository:     tt.fields.userRepository,
				userInfoRepository: tt.fields.userInfoRepository,
				logger:             tt.fields.logger,
			}
			if err := s.Delete(tt.args.userInfo); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignalingUseCase_StaticSearch(t *testing.T) {
	type fields struct {
		userRepository     repository.UserRepository
		userInfoRepository repository.UserInfoRepository
		logger             *zerolog.Logger
	}
	type args struct {
		userInfo       model.UserInfo
		searchDistance float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.UserInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SignalingUseCase{
				userRepository:     tt.fields.userRepository,
				userInfoRepository: tt.fields.userInfoRepository,
				logger:             tt.fields.logger,
			}
			got, err := s.StaticSearch(tt.args.userInfo, tt.args.searchDistance)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaticSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
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
		logger             *zerolog.Logger
	}
	type args struct {
		userInfo       model.UserInfo
		searchDistance float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.UserInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SignalingUseCase{
				userRepository:     tt.fields.userRepository,
				userInfoRepository: tt.fields.userInfoRepository,
				logger:             tt.fields.logger,
			}
			got, err := s.DynamicSearch(tt.args.userInfo, tt.args.searchDistance)
			if (err != nil) != tt.wantErr {
				t.Errorf("DynamicSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DynamicSearch() got = %v, want %v", got, tt.want)
			}
		})
	}
}