package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"

)

type Diamond struct {
	ID        bson.ObjectId   `bson:"_id" json:"_id"`
	CreatedDate *time.Time `bson:"created_date" json:"created_date"`
	UpdatedDate *time.Time `bson:"updated_date,omitempty" json:"updated_date,omitempty"`
	Url      string  `bson:"url,omitempty" json:"url,omitempty"`
	Lab      string  `bson:"lab,omitempty" json:"lab,omitempty"`
	Cert     int     `bson:"cert,omitempty" json:"cert,omitempty"`
	image 		string `bson:"img,omitempty" json:"img,omitempty"`
	Price    float32 `bson:"price,omitempty" json:"price,omitempty"`
	Carat    float32 `bson:"carat,omitempty" json:"carat,omitempty"`
	Cut      string  `bson:"cut,omitempty" json:"cut,omitempty"`
	Color    string  `bson:"color,omitempty" json:"color,omitempty"`
	Clarity  string  `bson:"clarity,omitempty" json:"clarity,omitempty"`
	Depth    float32 `bson:"depth,omitempty" json:"depth,omitempty"`
	Table    float32 `bson:"table,omitempty" json:"table,omitempty"`
	Crown    float32 `bson:"crown,omitempty" json:"crown,omitempty"`
	Pavilion float32 `bson:"pavilion,omitempty" json:"pavilion,omitempty"`
	Culet    float32 `bson:"culet,omitempty" json:"culet,omitempty"`
	Diameter float32 `bson:"diameter,omitempty" json:"diameter,omitempty"`
	HCA      float32 `bson:"hca,omitempty" json:"hca,omitempty"`
	CutScore float32 `bson:"cutscore,omitempty" json:"cutscore,omitempty"`
}
