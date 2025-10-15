package users

import (
	"errors"
	"natasha/src/database"
	"natasha/src/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func isEmailExists(email string) (bool, error) {
	var count int64
	if err := database.Repo.Model(&User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func handleRegisterUser(req *RegisterUserRequest) error {
	// check email is exists
	exists, err := isEmailExists(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}

	// hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	// create code
	code, err := utils.GenerateCode6Digits()
	if err != nil {
		return err
	}

	// create time expiry
	expiry := time.Now().Add(time.Minute * 5)

	// last send code at
	lastSend := time.Now()

	userData := User{
		Email:          req.Email,
		FullName:       req.FullName,
		Password:       string(hashPassword),
		Role:           "user",
		IsActive:       false,
		CodeExpiry:     &expiry,
		CodeActive:     &code,
		LastSendCodeAt: &lastSend,
	}
	err = database.Repo.Create(&userData).Error
	if err != nil {
		return err
	}

	// email content and send email
	emailData := utils.EmailData{
		ToEmail:      req.Email,
		Subject:      "Xác thực tài khoản",
		TemplatePath: "src/templates/authTemplate.html",
		TemplateData: map[string]string{
			"Name": req.Email,
			"Code": code,
		},
	}
	err = utils.SendingEmail(&emailData)
	if err != nil {
		return err
	}

	return nil
}
