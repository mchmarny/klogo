package main

// ServiceRequest represents service input
type ServiceRequest struct {
	ID       string `json:"id"`
	ImageURL string `json:"url"`
}

// ServiceResponse represents service output
type ServiceResponse struct {
	ID          string `json:"id"`
	ImageURL    string `json:"url"`
	Description string `json:"desc"`
}
