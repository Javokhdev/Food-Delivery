package postgres

import (
	"database/sql"

	pb "auth/genproto/user"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (p *UserStorage) GetProfile(req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	var profile pb.GetProfileResponse
	query := `
		SELECT id, username, email, name, phone, address
		FROM users
		WHERE id = $1
	`
	err := p.db.QueryRow(query, req.Token).Scan(
		&profile.UserId, &profile.Username, &profile.Email,
		&profile.Name, &profile.Phone, &profile.Address)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (p *UserStorage) UpdateProfile(req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	query := `
		UPDATE users
		SET username = $1, email = $2, name = $3, phone = $4, address = $5, updated_at = now()
		WHERE id = $6
	`
	_, err := p.db.Exec(query, req.Username, req.Email, req.Name, req.Phone, req.Address, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProfileResponse{Message: "Profile successfully updated", Success: true}, nil
}

func (p *UserStorage) ChangePassword(req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	query := `
		UPDATE users
		SET password_hash = $1, updated_at = now()
		WHERE password_hash = $2 and id = $3
	`
	_, err := p.db.Exec(query, req.NewPassword, req.CurrentPassword, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.ChangePasswordResponse{Message: "Password successfully changed", Success: true}, nil
}
