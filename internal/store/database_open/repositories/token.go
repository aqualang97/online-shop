package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"online-shop/internal/models"
)

type TokenRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewTokenRepo(db *sql.DB) *TokenRepo {
	return &TokenRepo{DB: db}
}

func (t *TokenRepo) GetTokensByUserID(userID int) (*models.Token, error) {
	var tokenModel models.Token

	err := t.DB.QueryRow(
		"SELECT user_id, access_hash, refresh_hash, created_at, updated_at FROM tokens WHERE user_id = ?",
		userID).Scan(&tokenModel.UserID, &tokenModel.AccessTokenHash, &tokenModel.RefreshTokenHash,
		&tokenModel.CreatedAt, &tokenModel.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &tokenModel, nil
}

//func (t *TokenRepo) GetTokensByID(tokenID int) (*models.Token, error) {
//	var tokenModel models.Token
//
//	err := t.DB.QueryRow(
//		"SELECT id, user_id, access_hash, refresh_hash, created_at, updated_at FROM tokens WHERE id = ?",
//		tokenID).Scan(&tokenModel.ID, &tokenModel.UserID, &tokenModel.AccessTokenHash, &tokenModel.RefreshTokenHash,
//		&tokenModel.CreatedAt, &tokenModel.UpdatedAt)
//	if err != nil {
//		return nil, err
//	}
//
//	return &tokenModel, nil
//}

func (t *TokenRepo) CreateToken(tokenModel *models.Token) error {
	if tokenModel == nil {
		return errors.New("incorrect token model")
	}
	uid, err := tokenModel.UserID.MarshalBinary()
	if err != nil {
		return err
	}
	if t.TX != nil {
		prepare, err := t.TX.Prepare("INSERT INTO users_tokens( user_id, access_hash, refresh_hash) VALUES ($1,$2,$3)")
		if err != nil {
			return err
		}
		fmt.Println(123)
		_, err = prepare.Exec(uid, tokenModel.AccessTokenHash, tokenModel.RefreshTokenHash)
		if err != nil {
			return err
		}
		fmt.Println(456)
		return nil
	}
	prepare, err := t.DB.Prepare("INSERT INTO users_tokens( user_id, access_hash, refresh_hash) VALUES ($1,$2,$3)")
	if err != nil {
		return err
	}
	_, err = prepare.Exec(uid, tokenModel.AccessTokenHash, tokenModel.RefreshTokenHash)
	if err != nil {
		return err
	}
	return nil
}

func (t *TokenRepo) UpdateTokens(tokenModel *models.Token) error {
	if tokenModel == nil {
		return errors.New("incorrect data to update")
	}

	if t.TX != nil {
		prepare, err := t.TX.Prepare("UPDATE users_tokens SET access_hash=$2, refresh_hash=$3 WHERE user_id=$1")
		if err != nil {
			return err
		}
		_, err = prepare.Exec(tokenModel.UserID, tokenModel.AccessTokenHash, tokenModel.RefreshTokenHash)
		if err != nil {
			return err
		}
		return nil
	}
	prepare, err := t.DB.Prepare("UPDATE users_tokens SET access_hash=$2, refresh_hash=$3 WHERE user_id=$1")
	if err != nil {
		return err
	}
	_, err = prepare.Exec(tokenModel.UserID, tokenModel.AccessTokenHash, tokenModel.RefreshTokenHash)
	if err != nil {
		return err
	}
	return nil
}

func (t *TokenRepo) DeleteTokensByUserID(userID int) error {
	if t.TX != nil {
		_, err := t.TX.Exec("DELETE FROM users_tokens WHERE user_id=$1", userID)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := t.TX.Exec("DELETE FROM users_tokens WHERE user_id=$1", userID)
	if err != nil {
		return err
	}
	return nil
}

//func (t *TokenRepo) DeleteTokensByID(userID int) error {
//	if t.TX != nil {
//		_, err := t.TX.Exec("DELETE FROM users_tokens WHERE id=$1", userID)
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//	_, err := t.TX.Exec("DELETE FROM users_tokens WHERE id=$1", userID)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (t *TokenRepo) BeginTX() error {
	tx, err := t.DB.Begin()
	if err != nil {
		return err
	}
	t.TX = tx
	return nil
}
func (t *TokenRepo) RollbackTX() error {
	defer func() {
		t.TX = nil
	}()
	if t.TX != nil {
		return t.TX.Rollback()
	}
	return nil
}
func (t *TokenRepo) CommitTX() error {
	defer func() {
		t.TX = nil
	}()
	if t.TX != nil {
		return t.TX.Commit()
	}
	return nil
}
func (t *TokenRepo) SetTX(tx *sql.Tx) {
	t.TX = tx
}
func (t *TokenRepo) GetTX() *sql.Tx {
	return t.TX
}
