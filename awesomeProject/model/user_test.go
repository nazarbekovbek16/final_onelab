package model

import "testing"

func TestUser_SetPassword(t *testing.T) {
	type fields struct {
		ID       int
		Name     string
		Surname  string
		Email    string
		Password string
	}
	type args struct {
		plaintextPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				ID:       1,
				Name:     "bek",
				Surname:  "nazarbekov",
				Email:    "bek@mail.ru",
				Password: "1234",
			},
			args:    args{plaintextPassword: "1234"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &User{
				ID:       tt.fields.ID,
				Name:     tt.fields.Name,
				Surname:  tt.fields.Surname,
				Email:    tt.fields.Email,
				Password: tt.fields.Password,
			}
			if err := p.SetPassword(tt.args.plaintextPassword); (err != nil) != tt.wantErr {
				t.Errorf("SetPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_MatchesPassword(t *testing.T) {
	type fields struct {
		ID       int
		Name     string
		Surname  string
		Email    string
		Password string
	}
	type args struct {
		plaintextPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				ID:       1,
				Name:     "bek",
				Surname:  "nazarbekov",
				Email:    "bek@mail.ru",
				Password: "$2a$12$4xOfXywc4mKYM.Dka4AdquhJxcDzunRvhbURQxI5XMlebSeWausdK",
			},
			args:    args{plaintextPassword: "12345"},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &User{
				ID:       tt.fields.ID,
				Name:     tt.fields.Name,
				Surname:  tt.fields.Surname,
				Email:    tt.fields.Email,
				Password: tt.fields.Password,
			}
			got, err := p.MatchesPassword(tt.args.plaintextPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatchesPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MatchesPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_MatchesPassword1(t *testing.T) {
	type fields struct {
		ID       int
		Name     string
		Surname  string
		Email    string
		Password string
	}
	type args struct {
		plaintextPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				ID:       1,
				Name:     "bek",
				Surname:  "nazarbekov",
				Email:    "bek@mail.ru",
				Password: "$2a$12$4xOfXywc4mKYM.Dka4AdquhJxcDzunRvhbURQxI5XMlebSeWausdK",
			},
			args:    args{plaintextPassword: "12345"},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &User{
				ID:       tt.fields.ID,
				Name:     tt.fields.Name,
				Surname:  tt.fields.Surname,
				Email:    tt.fields.Email,
				Password: tt.fields.Password,
			}
			got, err := p.MatchesPassword(tt.args.plaintextPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatchesPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MatchesPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}
