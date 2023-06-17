package authRepo

import (
	"aramen-api/cmd/api/ent"
	"aramen-api/cmd/api/ent/user"
	"aramen-api/cmd/api/ent/userotp"
	authModel "aramen-api/cmd/api/model/auth"
	"aramen-api/srv/hash"
	mailbit "aramen-api/srv/mailbit"
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type AuthRepoInterface interface {
	Login() (*int, error)
	CreateUser() error
	RequestOtp(phoneNumber string) (*string, error)
}

type AuthRepo struct {
	db *ent.Client
}

func NewAuthRepo(db *ent.Client) AuthRepo {
	return AuthRepo{db}
}

func (r AuthRepo) Login(login authModel.LoginModel) (*int, error) {
	otp, err := r.db.UserOtp.Query().Where(userotp.ReferenceEQ(login.Reference)).WithUser().First(context.Background())
	if err != nil {
		return nil, errors.New("invalid user")
	}
	if hash.Compare(login.Otp, otp.HashedOtp) {
		return &otp.UserID, nil
	}

	return nil, errors.New("invalid otp")
}

func (r AuthRepo) CreateUser(user authModel.CreateUserModel) error {
	_, err := r.db.User.Create().SetPhoneNumber(user.PhoneNumber).SetFirstName(user.Firstname).SetLastName(user.Lastname).SetEmail(user.Email).SetBirthdate(user.Birthdate).Save(context.Background())
	return err
}

func (r AuthRepo) RequestOtp(phoneNumber string) (*string, error) {
	user, err := r.db.User.Query().Where(user.PhoneNumberEQ(phoneNumber)).First(context.Background())
	if err != nil {
		return nil, err
	}
	if user == nil {
		//create new user
		user, err = r.db.User.Create().SetPhoneNumber(phoneNumber).Save(context.Background())
		if err != nil {
			return nil, err
		}
	}

	ref, _ := mailbit.Generate(6, true)
	otp, _ := mailbit.Generate(6, false)
	fmt.Println(otp)
	hashedOtp := hash.Hash(otp)
	_, err = r.db.UserOtp.Create().SetUserID(user.ID).SetReference(ref).SetHashedOtp(hashedOtp).Save(context.Background())
	// mailbit.SendOtp(ref, otp, phoneNumber)
	if err != nil {
		return nil, err
	}
	return &ref, nil
}
