package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openai_golang/src/model/po"
)

type EmbeddingRepo struct{}

//func NewEmbedding() *EmbeddingRepo {
//	return &EmbeddingRepo{}
//}

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
