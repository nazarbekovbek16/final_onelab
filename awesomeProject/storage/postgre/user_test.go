package postgre

import (
	"awesomeProject/config"
	"awesomeProject/model"
	"context"
	"reflect"
	"testing"
)

func TestUserRepositry_CreateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{context.Background(), model.User{ID: 6}},
			want:    6,
			wantErr: false,
		},
	}
	repo := CreateUserRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.CreateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func CreateUserRepo(t *testing.T) *UserRepositry {
	cfg := config.NewConfig()
	db, _ := OpenDB(cfg)
	return NewUserRepositry(db)
}

func TestUserRepositry_GetUser(t *testing.T) {
	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		args    args
		want    model.User
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				ID:  3,
			},
			want: model.User{
				ID:       3,
				Name:     "bek1",
				Surname:  "zz",
				Email:    "bek1@mail.ru",
				Password: "$2a$12$4xOfXywc4mKYM.Dka4AdquhJxcDzunRvhbURQxI5XMlebSeWausdK",
			},
			wantErr: false,
		},
	}
	repo := CreateUserRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetUser(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
