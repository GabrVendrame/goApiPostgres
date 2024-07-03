package models

type Todo struct {
    Id int64 `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    Done bool `json:"done"`

}
