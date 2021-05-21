package access

import (
	"context"
	"main/core"
	"main/models"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Challenge struct {
	DB             *firestore.Client
	CollectionName string
}

func (a *Challenge) Create(challenge *models.Challenge) (err error) {
	id := core.UseUtil().UUID()
	_, err = a.DB.Collection(a.CollectionName).Doc(id).Set(context.Background(), challenge)
	return
}

func (a *Challenge) Update(challenge *models.Challenge) (err error) {
	_, err = a.DB.Collection(a.CollectionName).Doc(challenge.Id).Set(context.Background(), challenge)
	return
}

func (a *Challenge) Delete(id string) (err error) {
	_, err = a.DB.Collection(a.CollectionName).Doc(id).Delete(context.Background())
	return
}

func (a *Challenge) GetById(id string) (challenge *models.Challenge, err error) {
	snapshot, err := a.DB.Collection(a.CollectionName).Doc(id).Get(context.Background())
	err = snapshot.DataTo(challenge)
	return
}

func (a *Challenge) GetByAuthor(authorId string) (challenges []*models.Challenge, err error) {
	iter := a.DB.Collection(a.CollectionName).Where("authorId", "==", authorId).Documents(context.Background())
	challenges = make([]*models.Challenge, 0)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var challenge models.Challenge
		doc.DataTo(&challenge)
		challenges = append(challenges, &challenge)
	}
	iter.Stop()
	return
}
