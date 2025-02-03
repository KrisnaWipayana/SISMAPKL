package auth_service

import (
	"context"
	"sismapkl/internal/pkg/jwt"
	"sismapkl/internal/pkg/schemas"

	"github.com/rotisserie/eris"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) LoginMahasiswa(ctx context.Context, NIM string, Password string) (
	*schemas.LoginInputMahasiswa, error) {

	mahasiswa, err := s.AuthRepository.CheckUserMahasiswa(ctx, nil, NIM)
	if err != nil {
		return nil, eris.Wrap(err, "NIM atau password tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(mahasiswa.Password), []byte(Password)); err != nil {
		return nil, eris.New("NIM atau password tidak ditemukan")
	}

	token, err := jwt.GenerateToken(mahasiswa.ID.String())
	if err != nil {
		eris.New("gagal mendapatkan token")
	}

	return &schemas.LoginInputMahasiswa{
		NIM:   mahasiswa.NIM,
		Token: token,
	}, nil
}
