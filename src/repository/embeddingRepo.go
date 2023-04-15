package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openaigo/src/model/po"
)

type EmbeddingRepo struct{}

//func NewEmbedding() *EmbeddingRepo {
//	return &EmbeddingRepo{}
//}

func (r *EmbeddingRepo) All(db *gorm.DB) ([]po.Embedding, error) {
	embeddings := make([]po.Embedding, 0)
	if err := db.Table(po.TableNameEmbedding).Find(&embeddings).Error; err != nil {
		return nil, err
	}
	return embeddings, nil
}

func (r *EmbeddingRepo) AllByTitle(db *gorm.DB, title string) ([]po.Embedding, error) {
	embeddings := make([]po.Embedding, 0)
	if err := db.Table(po.TableNameEmbedding).
		Where("title", title).
		Find(&embeddings).Error; err != nil {
		return nil, err
	}
	return embeddings, nil
}

func (r *EmbeddingRepo) Save(db *gorm.DB, embeddings ...po.Embedding) error {
	return db.Table(po.TableNameEmbedding).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "title"},
				{Name: "heading"},
			},
			DoUpdates: clause.AssignmentColumns([]string{"content", "vectors"}),
		}).
		Create(&embeddings).Error
}
