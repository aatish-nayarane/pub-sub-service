package mango

import (
	"pub-sub-service/internal/models"
	"context"
)

type SwitchRepo struct {
	coll *Collection
}

func NewSwitchRepo(coll *Collection) *SwitchRepo {
	return &SwitchRepo{
		coll: coll,
	}
}

func (r *SwitchRepo) Create(ctx context.Context, switchData *models.Switch) error {
	_, err := r.coll.Switch.InsertOne(ctx, switchData)
	return err
}

// func (r *SwitchRepo) GetSwitchByProjectID(ctx context.Context, projectID string) (*models.Switch, error) {
	// var switchData models.Switch
	// err := r.coll.Switch.FindOne(ctx, map[string]interface{}{"project_id": projectID}).Decode(&switchData)
	// if err != nil {
	// 	return nil, err
	// }