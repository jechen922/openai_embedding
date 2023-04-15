package seedhandler

import (
	"github.com/gofiber/fiber/v2"
	"openaigo/openai/embedding"
	"openaigo/src/database/postgres"
	"openaigo/src/lib/file"
	"openaigo/src/service/seedservice"
)

//type seedHandler struct {
//}
// func (h *seedHandler) Seed() error {

func Seed(ctx *fiber.Ctx) error {
	csvRecords := file.ReadCSVByFields("./resources/yile/yile.csv", "title", "heading", "content")
	sections := make([]embedding.Section, 0, len(csvRecords))
	for _, record := range csvRecords {
		sections = append(sections, embedding.Section{
			Title:   record["title"],
			Heading: record["heading"],
			Content: record["content"],
		})
	}

	db, _ := postgres.New()
	ss := seedservice.NewSeed(db)
	if err := ss.SaveSections(sections); err != nil {
		return err
	}
	return ctx.Send([]byte("OK!"))
}
