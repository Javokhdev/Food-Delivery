package postgres

import (
	"database/sql"
	"log/slog"

	pb "auth/genproto/auth"

	"github.com/google/uuid"
)

type AuthStorage struct {
	db *sql.DB
}

func NewAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{db: db}
}

func (p *AuthStorage) Register(req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userId := uuid.NewString()
	query := `
		INSERT INTO users (id, username, email, password_hash, name, phone, address, role
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := p.db.Exec(query, userId, req.Username, req.Email, req.Password, req.Name, req.Phone, req.Address, req.Role)
	if err != nil {
		slog.Info(err.Error())
		return nil, err
	}

	tokenId := uuid.NewString()
	tokenQuery := `
		INSERT INTO tokens (id, token, user_id)
		VALUES ($1, $2, $3)
	`
	_, err = p.db.Exec(tokenQuery, tokenId, req.Token, userId)
	if err != nil {
		slog.Info(err.Error())
		return nil, err
	}

	return &pb.RegisterResponse{Id: userId, Message: "User created successfully", Success: true}, nil
}

func (p *AuthStorage) Login(req *pb.LoginRequest) (*pb.LoginResponse, error) {
	query := `
		SELECT id
		FROM users
		WHERE username = $1 AND password_hash = $2
	`
	var userId string
	err := p.db.QueryRow(query, req.Username, req.Password).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.LoginResponse{Message: "Invalid username or password", Success: false}, nil
		}
		return nil, err
	}
	var token string
	getTokenQuery := `
		SELECT token
		FROM tokens
		WHERE user_id = $1
	`
	err = p.db.QueryRow(getTokenQuery, userId).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.LoginResponse{Message: "Token not found", Success: false}, nil
		}
		return nil, err
	}
	return &pb.LoginResponse{Token: token, Message: "Login successful", Success: true}, nil
}

func (p *AuthStorage) Logout(req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	query := `
		DELETE FROM tokens
		WHERE token = $1
	`
	_, err := p.db.Exec(query, req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.LogoutResponse{Message: "Logged out successfully"}, nil
}

func (p *AuthStorage) ResetPassword(req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	query := `
		UPDATE users
		SET password_hash = $1
		WHERE email = $2 and username = $3
	`
	_, err := p.db.Exec(query, req.NewPassword, req.Email, req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.ResetPasswordResponse{Message: "Password reset successfully"}, nil
}
