package service

import (
	"awesomeProject/config"
	"awesomeProject/logger"
	"awesomeProject/model"
	"awesomeProject/storage"
	"context"
	"go.uber.org/zap"
	"log"
	"reflect"
	"testing"
)

func TestUserService_Create(t *testing.T) {
	conf := config.NewConfig()

	ctx, cancel := context.WithCancel(context.Background())

	l, err := logger.Init(conf)

	defer func(l *zap.Logger) {
		err = l.Sync()
		if err != nil {
			log.Fatalln(err)
		}
	}(l)
	repo, _ := storage.NewStorage(l, ctx, conf)

	defer cancel()

	type fields struct {
		repo *storage.Storage
	}
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Fail",
			fields: fields{repo: repo},
			args: args{
				ctx: ctx,
				user: model.User{
					ID:       1,
					Name:     "Nurkhat",
					Surname:  "Amanbayev",
					Email:    "nurkhat@mail.ru",
					Password: "1234",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserService{
				repo: tt.fields.repo,
			}
			if err := s.Create(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_Get(t *testing.T) {
	conf := config.NewConfig()

	ctx, cancel := context.WithCancel(context.Background())

	l, err := logger.Init(conf)
	defer func(l *zap.Logger) {
		err = l.Sync()
		if err != nil {
			log.Fatalln(err)
		}
	}(l)
	repo, _ := storage.NewStorage(l, ctx, conf)

	defer cancel()

	type fields struct {
		repo *storage.Storage
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.User
		wantErr bool
	}{
		{
			name:   "Success",
			fields: fields{repo: repo},
			args: args{
				ctx:      ctx,
				username: "bek@mail.ru",
			},
			want: model.User{
				ID:       2,
				Name:     "bek",
				Surname:  "zz",
				Email:    "bek@mail.ru",
				Password: "$2a$12$cGBISo.i8ukOtRanss/Hi.GQilJxgem5mZWK0W6.S6ZrFWCPp3RHu",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserService{
				repo: tt.fields.repo,
			}
			got, err := s.Get(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
