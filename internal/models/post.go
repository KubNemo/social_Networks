package models

import "time"

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`   // ID пользователя (внешний ключ)
	Content   string    `json:"content"`   // Текст поста
	ImageURL  string    `json:"image_url"` // Ссылка на изображение (если есть)
	Likes     int       `json:"likes"`     // Количество лайков
	Comments  int       `json:"comments"`  // Кол-во комментариев (можно кэшировать)
	IsEdited  bool      `json:"is_edited"` // Был ли пост отредактирован
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
