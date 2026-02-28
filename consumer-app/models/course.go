package models

type Course struct {
    // Change ID to string to match Java's private String courseId
    ID          string  `json:"courseId"` 
    Title       string  `json:"title"`
    // Rename this to Trainer to match Java's trainer field
    Trainer     string  `json:"trainer"` 
    Price       float64 `json:"price"`
}
