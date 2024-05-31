// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/chenmingyong0423/fnote/server/internal/website_config/internal/domain"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/chenmingyong0423/go-mongox/builder/query"

	"github.com/chenmingyong0423/go-mongox"
	"github.com/chenmingyong0423/go-mongox/bsonx"
	"github.com/chenmingyong0423/go-mongox/builder/update"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

// WebsiteConfig defines for the MongoDB Collection "website_config"
type WebsiteConfig struct {
	mongox.Model `bson:",inline"`
	Props        any    `bson:"props"`
	Typ          string `bson:"typ"`
}

type IWebsiteConfigDao interface {
	FindByTyp(ctx context.Context, typ string) (*WebsiteConfig, error)
	Increase(ctx context.Context, field string) error
	GetByTypes(ctx context.Context, types ...string) ([]*WebsiteConfig, error)
	Decrease(ctx context.Context, field string) error
	UpdateByConditionAndUpdates(ctx context.Context, cond bson.D, updates bson.D) error
	UpdatePropsByTyp(ctx context.Context, typ string, cfg any, now time.Time) error
	AddTPSVConfig(ctx context.Context, tpsv domain.TPSV) error
	DeleteTPSVConfigByKey(ctx context.Context, key string) error
	UpdatePostIndexProps(ctx context.Context, updates bson.D) error
	PushCarouselConfig(ctx context.Context, carouselElem domain.CarouselElem) error
	FindByFilter(ctx context.Context, filter bson.D) (*WebsiteConfig, error)
}

var _ IWebsiteConfigDao = (*WebsiteConfigDao)(nil)

func NewWebsiteConfigDao(db *mongo.Database) *WebsiteConfigDao {
	return &WebsiteConfigDao{
		coll: mongox.NewCollection[WebsiteConfig](db.Collection("configs")),
	}
}

type WebsiteConfigDao struct {
	coll *mongox.Collection[WebsiteConfig]
}

func (d *WebsiteConfigDao) FindByFilter(ctx context.Context, filter bson.D) (*WebsiteConfig, error) {
	return d.coll.Finder().Filter(filter).FindOne(ctx)
}

func (d *WebsiteConfigDao) PushCarouselConfig(ctx context.Context, carouselElem domain.CarouselElem) error {
	updateResult, err := d.coll.Updater().Filter(query.Eq("typ", "carousel")).Updates(update.Push("props.list", carouselElem)).UpdateOne(ctx)
	if err != nil {
		return errors.Wrapf(err, "fails to push carousel config, carouselElem=%v", carouselElem)
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("ModifiedCount=0, fails to push carousel config, carouselElem=%v", carouselElem)
	}
	return nil
}

func (d *WebsiteConfigDao) UpdatePostIndexProps(ctx context.Context, updates bson.D) error {
	updateResult, err := d.coll.Updater().Filter(query.Eq("typ", "post index")).Updates(updates).UpdateOne(ctx)
	if err != nil {
		return errors.Wrapf(err, "fails to update website_config, typ=post index, updates=%v", updates)
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("ModifiedCount=0, fails to update website_config, typ=post index, updates=%v", updates)
	}
	return nil
}

func (d *WebsiteConfigDao) DeleteTPSVConfigByKey(ctx context.Context, key string) error {
	updateResult, err := d.coll.Updater().
		Filter(query.Eq("typ", "third party site verification")).
		Updates(update.Pull("props.list", bsonx.M("key", key))).
		UpdateOne(ctx)
	if err != nil {
		return errors.Wrapf(err, "fails to delete tpsv config, key=%s", key)
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("DeletedCount=0, fails to delete tpsv config, key=%s", key)
	}
	return nil
}

func (d *WebsiteConfigDao) AddTPSVConfig(ctx context.Context, tpsv domain.TPSV) error {
	updateResult, err := d.coll.Updater().Filter(query.Eq("typ", "third party site verification")).Updates(
		update.NewBuilder().Push("props.list", tpsv).Set("updated_at", time.Now().Local()).Build(),
	).UpdateOne(ctx)
	if err != nil {
		return errors.Wrapf(err, "fails to add tpsv config, tpsv=%v", tpsv)
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("ModifiedCount=0, fails to add tpsv config, tpsv=%v", tpsv)
	}
	return nil
}

func (d *WebsiteConfigDao) UpdatePropsByTyp(ctx context.Context, typ string, cfg any, now time.Time) error {
	updateResult, err := d.coll.Updater().Filter(bsonx.M("typ", typ)).Updates(update.NewBuilder().Set("props", cfg).Set("updated_at", now).Build()).UpdateOne(ctx)
	if err != nil {
		return errors.Wrapf(err, "fails to update %s config, updates=%v", typ, cfg)
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("ModifiedCount=0, fails to update %s config, updates=%v", typ, cfg)
	}
	return nil
}

func (d *WebsiteConfigDao) UpdateByConditionAndUpdates(ctx context.Context, cond bson.D, updates bson.D) error {
	updateOne, err := d.coll.Updater().Filter(cond).Updates(updates).UpdateOne(ctx)
	if err != nil {
		return errors.Wrapf(err, "fails to update website_config, cond=%v, updates=%v", cond, updates)
	}
	if updateOne.ModifiedCount == 0 {
		return fmt.Errorf("ModifiedCount=0, fails to update website_config, cond=%v, updates=%v", cond, updates)
	}
	return nil
}

func (d *WebsiteConfigDao) Decrease(ctx context.Context, field string) error {
	field = fmt.Sprintf("props.%s", field)
	updateResult, err := d.coll.Updater().Filter(bsonx.M("typ", "website")).Updates(update.Inc(field, -1)).UpdateOne(ctx)
	if err != nil {
		return errors.Wrapf(err, "fails to increase %s", field)
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("ModifiedCount=0, fails to increase %s", field)
	}
	return nil
}

func (d *WebsiteConfigDao) GetByTypes(ctx context.Context, types ...string) ([]*WebsiteConfig, error) {
	configs, err := d.coll.Finder().Filter(query.In("typ", types...)).Find(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "fails to find configs by types, types=%v", types)
	}
	return configs, nil
}

func (d *WebsiteConfigDao) Increase(ctx context.Context, field string) error {
	field = fmt.Sprintf("props.%s", field)
	updateResult, err := d.coll.Updater().Filter(bsonx.M("typ", "website")).Updates(update.Inc(field, 1)).UpdateOne(ctx)
	if err != nil {
		return errors.Wrapf(err, "fails to increase %s", field)
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("ModifiedCount=0, fails to increase %s", field)
	}
	return nil
}

func (d *WebsiteConfigDao) FindByTyp(ctx context.Context, typ string) (*WebsiteConfig, error) {
	config, err := d.coll.Finder().Filter(bsonx.M("typ", typ)).FindOne(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Find website_config failed, typ=%s", typ)
	}
	return config, nil
}
