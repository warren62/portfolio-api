package post

import (
	"time"

	"github.com/google/uuid"
)

type PostSubject int

const (
	Choose PostSubject = iota
	Feedback
	Report
	Other
)

type DTO struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Message     string      `json:"message"`
	Description string      `json:"description"`
	Subject     PostSubject `json:"subject"`
}

type Form struct {
	Name        string `json:"name" form:"required,max=255"`
	Email       string `json:"email" form:"required,alpha_space,max=255"`
	Message     string `json:"message" form:"required,datetime=2006-01-02"`
	Description string `json:"description" form:"required,max=255"`
	Subject     string `json:"subject"  form:"required"`
}

type Post struct {
	ID          uuid.UUID `gorm:"primarykey"`
	Name        string
	Email       string
	Message     string
	Description string
	Subject     PostSubject
	CreatedDate time.Time
}

type Posts []*Post

func (b *Post) ToDto() *DTO {
	return &DTO{
		ID:          b.ID.String(),
		Name:        b.Name,
		Email:       b.Email,
		Message:     b.Message,
		Description: b.Description,
		Subject:     b.Subject,
	}
}

func (bs Posts) ToDto() []*DTO {
	dtos := make([]*DTO, len(bs))
	for i, v := range bs {
		dtos[i] = v.ToDto()
	}

	return dtos
}

func (f *Form) ToModel() *Post {

	return &Post{
		Name:        f.Name,
		Email:       f.Email,
		Message:     f.Message,
		Description: f.Description,
		Subject:     0,
	}
}
