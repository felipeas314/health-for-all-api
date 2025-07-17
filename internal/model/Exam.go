package model

type Exam struct {
	ID        string `json:"id"`
	UserEmail string `json:"user_email"`
	FileName  string `json:"file_name"`
	Type      string `json:"type"`
	Result    string `json:"result"`
	CreatedAt string `json:"created_at"`
}
