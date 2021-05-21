package business

import (
	"errors"
	"main/access"
	"main/models"

	"firebase.google.com/go/v4/auth"
)

type Challenge struct {
	ChallengeAccess *access.Challenge
}

func (b *Challenge) GetById(auth *auth.Token, id string) (challenge *models.Challenge, err error) {
	challenge, err = b.ChallengeAccess.GetById(id)
	if err != nil {
		return nil, err
	}
	if !challenge.IsPublic && challenge.AuthorId != auth.UID {
		return nil, errors.New("Cannot get private or non-ownership challenge")
	}
	return
}

func (b *Challenge) GetByAuthor(auth *auth.Token) (challenges []*models.Challenge, err error) {
	return b.ChallengeAccess.GetByAuthor(auth.UID)
}

func (b *Challenge) Create(auth *auth.Token, challenge *models.Challenge) (err error) {
	challenge.AuthorId = auth.UID
	return b.ChallengeAccess.Create(challenge)
}

func (b *Challenge) Update(auth *auth.Token, challenge *models.Challenge) (err error) {
	if challenge.AuthorId != auth.UID {
		return errors.New("Cannot update non-ownership challenge")
	}
	return b.ChallengeAccess.Update(challenge)
}

func (b *Challenge) Delete(auth *auth.Token, id string) (err error) {
	_, err = b.GetById(auth, id)
	if err != nil {
		return errors.New("Cannot delete non-ownership challenge")
	}
	return b.ChallengeAccess.Delete(id)
}
