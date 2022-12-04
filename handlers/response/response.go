package response

type UserResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	CompanyName string `json:"company_name"`
	JobTitle    string `json:"job_title"`
	Active      bool   `json:"active"`
}