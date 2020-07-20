package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id uuid.UUID `gorm:"unique;not null"`

	// custom fields
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Role      string `gorm:"not null"`

	// authboss.User fields
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`

	// ConfirmedUser fields
	Confirmed         bool `gorm:"not null"`
	ConfirmedSelector string
	ConfirmedVerifier string

	// LockableUser fields
	AttemptCount    int
	LastAttemptTime time.Time
	LockedTime      time.Time
}

func (u User) GetAttemptCount() (attempts int) {
	return u.AttemptCount
}

func (u User) GetLastAttempt() (last time.Time) {
	return u.LastAttemptTime
}

func (u User) GetLocked() (locked time.Time) {
	return u.LockedTime
}

func (u *User) PutAttemptCount(attempts int) {
	u.AttemptCount = attempts
}

func (u *User) PutLastAttempt(last time.Time) {
	u.LastAttemptTime = last
}

func (u User) PutLocked(locked time.Time) {
	u.LockedTime = locked
}

func (u User) GetEmail() (email string) {
	return u.Email
}

func (u User) GetConfirmed() (confirmed bool) {
	return u.Confirmed
}

func (u User) GetConfirmSelector() (selector string) {
	return u.ConfirmedSelector
}

func (u User) GetConfirmVerifier() (verifier string) {
	return u.ConfirmedVerifier
}

func (u *User) PutEmail(email string) {
	u.Email = email
}

func (u *User) PutConfirmed(confirmed bool) {
	u.Confirmed = confirmed
}

func (u *User) PutConfirmSelector(selector string) {
	u.ConfirmedSelector = selector
}

func (u* User) PutConfirmVerifier(verifier string) {
	u.ConfirmedVerifier = verifier
}

func (u User) GetArbitrary() (r map[string]string) {
	r = make(map[string]string)
	r["first_name"] = u.FirstName
	r["last_name"] = u.LastName

	return
}

func (u User) GetPassword() (password string) {
	return u.Password
}

func (u *User) PutPassword(password string) {
	u.Password = password
}

func (u User) GetPID() string {
	return u.Email
}

func (u *User) PutPID(pid string) {
	u.Email = pid
}

func (u *User) PutArbitrary(arbitrary map[string]string) {
	if v, ok := arbitrary["first_name"]; ok {
		u.FirstName = v
	}
	if v, ok := arbitrary["last_name"]; ok {
		u.LastName = v
	}
}

const (
	RoleMember = "Member"
	RoleAdmin  = "Admin"
)
