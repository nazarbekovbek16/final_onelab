package service

import (
	"awesomeProject/model"
	"awesomeProject/storage"
	mock_storage "awesomeProject/storage/mocks"
	"context"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestBookService_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		book model.Book
	}
	tests := []struct {
		name    string
		prepare func(f *mock_storage.MockIBookRepository)
		args    args
		want    uint
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				book: model.Book{
					ID:     1,
					Title:  "Big Dream",
					Author: "Haruki",
				},
			},
			want:    1,
			wantErr: false,
			prepare: func(f *mock_storage.MockIBookRepository) {
				f.EXPECT().Create(
					gomock.Any(),
					model.Book{
						ID:     1,
						Title:  "Big Dream",
						Author: "Haruki",
					},
				).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		bookRepo := mock_storage.NewMockIBookRepository(ctrl)
		tt.prepare(bookRepo)
		s := NewBookService(&storage.Storage{Book: bookRepo})
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Create(tt.args.ctx, tt.args.book)
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

func TestBookService_Create1(t *testing.T) {
	type fields struct {
		repo *storage.Storage
	}
	type args struct {
		ctx  context.Context
		book model.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := BookService{
				repo: tt.fields.repo,
			}
			got, err := s.Create(tt.args.ctx, tt.args.book)
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
