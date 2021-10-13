package repository

import (
	"errors"
	"g-sig/pkg/domain/model"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestUserInfoRepository_Find(t *testing.T) {
	type args struct {
		userID string
	}
	var tests = []struct {
		name    string
		args    args
		want    *model.UserInfo
		wantErr error
	}{
		{
			name: "find_success",
			args: args{
				userID: "1234-1234-1234",
			},
			want: &model.UserInfo{
				UserID:      "1234-1234-1234",
				PublicIP:    "",
				PublicPort:  0,
				PrivateIP:   "",
				PrivatePort: 0,
				Latitude:    0,
				Longitude:   0,
			},
			wantErr: nil,
		},
		{
			name: "find_error",
			args: args{
				userID: "1111-1234-1234",
			},
			want:    nil,
			wantErr: model.ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserInfoRepository{
			}
			// リポジトリに登録
			err := u.Save(model.UserInfo{
				UserID:      "1234-1234-1234",
				PublicIP:    "",
				PublicPort:  0,
				PrivateIP:   "",
				PrivatePort: 0,
				Latitude:    0,
				Longitude:   0,
			})
			got, err := u.Find(tt.args.userID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
			u.Delete("1234-1234-1234")
		})
	}
}

func TestUserInfoRepository_FindAll(t *testing.T) {
	tests := []struct {
		name    string
		want    []*model.UserInfo
		wantErr error
	}{
		{
			name: "findAll_success",
			want: []*model.UserInfo{
				{
					UserID:      "1",
					PublicIP:    "",
					PublicPort:  0,
					PrivateIP:   "",
					PrivatePort: 0,
					Latitude:    0,
					Longitude:   0,
				},
				{
					UserID:      "2",
					PublicIP:    "",
					PublicPort:  0,
					PrivateIP:   "",
					PrivatePort: 0,
					Latitude:    0,
					Longitude:   0,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserInfoRepository{
			}
			err := u.Save(model.UserInfo{
				UserID:      "1",
				PublicIP:    "",
				PublicPort:  0,
				PrivateIP:   "",
				PrivatePort: 0,
				Latitude:    0,
				Longitude:   0,
			})
			err = u.Save(model.UserInfo{
				UserID:      "2",
				PublicIP:    "",
				PublicPort:  0,
				PrivateIP:   "",
				PrivatePort: 0,
				Latitude:    0,
				Longitude:   0,
			})
			got, err := u.FindAll()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
				t.Errorf("FindAll() diff = %v", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestUserInfoRepository_Save(t *testing.T) {
	type args struct {
		user model.UserInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "save_success",
			args: args{
				user: model.UserInfo{
					UserID:      "1234-1234",
					PublicIP:    "",
					PublicPort:  0,
					PrivateIP:   "",
					PrivatePort: 0,
					Latitude:    0,
					Longitude:   0,
				},
			},
			wantErr: nil,
		},
		{
			name: "save_error",
			args: args{
				user: model.UserInfo{
					UserID:      "1234-1234",
					PublicIP:    "",
					PublicPort:  0,
					PrivateIP:   "",
					PrivatePort: 0,
					Latitude:    0,
					Longitude:   0,
				},
			},
			wantErr: model.ErrUserAlreadyExisted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserInfoRepository{
			}
			err := u.Save(tt.args.user)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserInfoRepository_Update(t *testing.T) {
	type args struct {
		user model.UserInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "update_success",
			args: args{
				user: model.UserInfo{
					UserID:      "1234-1234",
					PublicIP:    "",
					PublicPort:  0,
					PrivateIP:   "",
					PrivatePort: 0,
					Latitude:    1,
					Longitude:   1,
				},
			},
			wantErr: nil,
		},
		{
			name: "update_error",
			args: args{
				user: model.UserInfo{
					UserID:      "00000",
					PublicIP:    "",
					PublicPort:  0,
					PrivateIP:   "",
					PrivatePort: 0,
					Latitude:    1,
					Longitude:   1,
				},
			},
			wantErr: model.ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserInfoRepository{
			}
			err := u.Save(model.UserInfo{
				UserID:      "1234-1234",
				PublicIP:    "",
				PublicPort:  0,
				PrivateIP:   "",
				PrivatePort: 0,
				Latitude:    0,
				Longitude:   0,
			})
			err = u.Update(tt.args.user)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserInfoRepository_Delete(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "delete_success",
			args: args{
				userID: "1111-1111",
			},
			wantErr: nil,
		},
		{
			name: "delete_error",
			args: args{
				userID: "",
			},
			wantErr: model.ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserInfoRepository{
			}
			err := u.Save(model.UserInfo{
				UserID:      "1111-1111",
				PublicIP:    "",
				PublicPort:  0,
				PrivateIP:   "",
				PrivatePort: 0,
				Latitude:    0,
				Longitude:   0,
			})
			err = u.Delete(tt.args.userID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
