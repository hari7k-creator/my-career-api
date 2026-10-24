package models

import (
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate = validator.New()

// ContentBlock represents a reusable content component
type ContentBlock struct {
	Type    string      `bson:"type" json:"type" validate:"required"`
	Title   string      `bson:"title" json:"title" validate:"required"`
	Content interface{} `bson:"content" json:"content" validate:"required"`
}

// RoadmapStep represents a single step in a career roadmap
type RoadmapStep struct {
	ID              string `bson:"id,omitempty" json:"id,omitempty"`
	Title           string `bson:"title" json:"title" validate:"required"`
	Description     string `bson:"description" json:"description" validate:"required"`
	EstimatedMonths int    `bson:"estimated_months" json:"estimated_months" validate:"gte=0,lte=120"`
	Dependencies    []string `bson:"dependencies" json:"dependencies"`
	Resources       []Resource `bson:"resources" json:"resources"`
	Status          string `bson:"status" json:"status" validate:"oneof=todo in-progress completed"`
	Tags            []string `bson:"tags" json:"tags"`
	Prerequisites   []string `bson:"prerequisites" json:"prerequisites"`
}

// Resource represents a learning resource
type Resource struct {
	Title string `bson:"title" json:"title"`
	URL   string `bson:"url" json:"url"`
	Type  string `bson:"type" json:"type" validate:"oneof=video article course book documentation tool"`
}

// Career represents a career pathway or blueprint
type Career struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name            string             `bson:"name" json:"name" validate:"required,min=2,max=100"`
	Slug            string             `bson:"slug" json:"slug" validate:"required,slug"`
	Description     string             `bson:"description" json:"description" validate:"required,min=10"`
	Category        string             `bson:"category" json:"category" validate:"required,oneof=civil tech finance design data health"`
	Difficulty      string             `bson:"difficulty" json:"difficulty" validate:"required,oneof=beginner intermediate advanced elite"`
	Tags            []string           `bson:"tags" json:"tags"`
	Skills          []string           `bson:"skills" json:"skills"`
	ContentBlocks   []ContentBlock     `bson:"content_blocks" json:"content_blocks"`
	Roadmap         []RoadmapStep      `bson:"roadmap" json:"roadmap"`
	Metrics         Metrics            `bson:"metrics" json:"metrics"`
	CreatedAt       string             `bson:"created_at" json:"created_at"`
	UpdatedAt       string             `bson:"updated_at" json:"updated_at"`
}

// Metrics represents career-specific metrics
type Metrics struct {
	DurationMonths int     `bson:"duration_months" json:"duration_months"`
	EntryThreshold string  `bson:"entry_threshold" json:"entry_threshold"`
	Compensation   string  `bson:"compensation" json:"compensation"`
	SelectionRatio string  `bson:"selection_ratio" json:"selection_ratio"`
	TotalMarks     int     `bson:"total_marks" json:"total_marks" validate:"gte=0"`
}

// ValidationError represents validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// init registers custom validators
func init() {
	validate.RegisterValidation("slug", validateSlug)
}

// RegisterValidators is called to register validators
func RegisterValidators() {
	// Already registered in init()
}

// validateSlug validates that the string is a valid slug
func validateSlug(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	if field == "" {
		return false
	}
	pattern := regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	return pattern.MatchString(field)
}

// Validate validates the Career model
func (c *Career) Validate() []ValidationError {
	var errors []ValidationError

	if err := validate.Struct(c); err != nil {
		if validations, ok := err.(validator.ValidationErrors); ok {
			for _, ve := range validations {
				errors = append(errors, ValidationError{
					Field:   ve.Field(),
					Message: ve.Tag(),
				})
			}
		}
	}

	return errors
}
