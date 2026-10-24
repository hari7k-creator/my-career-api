package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hari7k-creator/my-career-api/internal/models"

	"github.com/gorilla/mux"
)

// CareerHandler handles career API requests
type CareerHandler struct {
	repository models.CareerRepository
}

// NewCareerHandler creates a new handler instance
func NewCareerHandler(repo models.CareerRepository) *CareerHandler {
	return &CareerHandler{repository: repo}
}

// GetCareers returns all careers with optional filtering
func (h *CareerHandler) GetCareers(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	category := r.URL.Query().Get("category")
	tags := r.URL.Query().Get("tags")
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("per_page")

	page := 1
	perPage := 20

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if p, err := strconv.Atoi(perPageStr); err == nil && p > 0 && p <= 100 {
		perPage = p
	}

	var careers []models.Career
	var err error
	var total int64

	if category != "" {
		careers, err = h.repository.FindByCategory(category)
	} else if tags != "" {
		tagList := splitTags(tags)
		careers, err = h.repository.FindByTags(tagList)
	} else {
		careers, err = h.repository.FindAll()
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	total, err = h.repository.Count()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Calculate pagination
	totalPages := int(total) / perPage
	if int(total)%perPage > 0 {
		totalPages++
	}

	response := models.Response{
		Success: true,
		Message: "Careers fetched successfully",
		Data: models.PaginatedResponse{
			Items:      careers,
			Total:      total,
			Page:       page,
			PerPage:    perPage,
			TotalPages: totalPages,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetCareerByID returns a single career by ID
func (h *CareerHandler) GetCareerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	career, err := h.repository.FindByID(id)
	if err != nil {
		if err.Error() == "career not found" {
			http.Error(w, "Career not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Success: true,
		Message: "Career fetched successfully",
		Data:    career,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetCareerBySlug returns a single career by slug
func (h *CareerHandler) GetCareerBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	career, err := h.repository.FindBySlug(slug)
	if err != nil {
		if err.Error() == "career not found" {
			http.Error(w, "Career not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Success: true,
		Message: "Career fetched successfully",
		Data:    career,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateCareer creates a new career
func (h *CareerHandler) CreateCareer(w http.ResponseWriter, r *http.Request) {
	var career models.Career

	if err := json.NewDecoder(r.Body).Decode(&career); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := career.Validate(); err != nil {
		response := models.Response{
			Success: false,
			Message: "Validation failed",
			Error:   err,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := h.repository.Create(&career); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Success: true,
		Message: "Career created successfully",
		Data:    career,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// UpdateCareer updates an existing career
func (h *CareerHandler) UpdateCareer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Find existing career
	existingCareer, err := h.repository.FindByID(id)
	if err != nil {
		if err.Error() == "career not found" {
			http.Error(w, "Career not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updatedCareer models.Career
	if err := json.NewDecoder(r.Body).Decode(&updatedCareer); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Preserve ID and timestamps
	updatedCareer.ID = existingCareer.ID
	updatedCareer.CreatedAt = existingCareer.CreatedAt

	if err := updatedCareer.Validate(); err != nil {
		response := models.Response{
			Success: false,
			Message: "Validation failed",
			Error:   err,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := h.repository.Update(&updatedCareer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Success: true,
		Message: "Career updated successfully",
		Data:    updatedCareer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteCareer deletes a career by ID
func (h *CareerHandler) DeleteCareer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.repository.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Success: true,
		Message: "Career deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetCategories returns all available categories
func (h *CareerHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories := []string{
		"civil",   // Civil & Public Policy
		"tech",    // Software & AI Engineering
		"finance", // Finance & Strategy
		"design",  // Product & Design
		"data",    // Data & Quantitative
		"health",  // Healthcare & Life Sciences
	}

	response := models.Response{
		Success: true,
		Message: "Categories fetched successfully",
		Data:    categories,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetDifficulties returns all available difficulty levels
func (h *CareerHandler) GetDifficulties(w http.ResponseWriter, r *http.Request) {
	difficulties := []string{
		"beginner",
		"intermediate",
		"advanced",
		"elite",
	}

	response := models.Response{
		Success: true,
		Message: "Difficulty levels fetched successfully",
		Data:    difficulties,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func splitTags(tagsStr string) []string {
	tags := []string{}
	for _, tag := range []string{"tech", "finance", "design", "data", "health"} {
		if containsString(tagsStr, tag) {
			tags = append(tags, tag)
		}
	}
	return tags
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsSubstring(s, substr)))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
