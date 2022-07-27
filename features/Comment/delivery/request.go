package delivery

import "backend/domain"

type CommentFormat struct {
	Userid  int
	Postid  int
	Comment string `json:"comment" form:"comment" validate:"required"`
}

func (i *CommentFormat) ToModel() domain.Comment {
	return domain.Comment{
		Userid:  i.Userid,
		Postid:  i.Postid,
		Comment: i.Comment,
	}
}
