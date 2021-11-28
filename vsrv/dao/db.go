package dao

import (
	"mithril-micro/dao"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewCl() *mongo.Collection {
	client := dao.NewMgoDB()
	cl := client.Database("mithril_video").Collection("video")
	return cl
}
