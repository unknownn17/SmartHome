package devicedatabase

import (
	"context"
	"device/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LIST struct {
	D *mongo.Collection
	C context.Context
}

func (u *LIST) Findone(req string) (*models.Devicelist, error) {
	var res models.Devicelist
	if err := u.D.FindOne(u.C, bson.M{"name": req}).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *LIST) FindAll() ([]models.Devicelist, error) {
	var res []models.Devicelist

	resp, err := u.D.Find(u.C, bson.M{})
	if err != nil {
		return nil, err
	}
	for resp.Next(u.C) {
		var all models.Devicelist
		if err := resp.Decode(&all); err != nil {
			return nil, err
		}
		res = append(res, all)
	}
	return res, nil
}
