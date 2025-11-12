package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/shagabiev/Go-QA-Api/internal/models"
	"github.com/shagabiev/Go-QA-Api/internal/repository"
	"github.com/sirupsen/logrus"
)

type QuestionHandler struct {
	repo *repository.QuestionRepository
}

func NewQuestionHandler(repo *repository.QuestionRepository) *QuestionHandler {
	return &QuestionHandler{repo: repo}
}

func (h *QuestionHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	questions, err := h.repo.FindAll()
	if err != nil {
		logrus.WithError(err).Error("failed to fetch questions")
		respondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	respondJSON(w, http.StatusOK, questions)
}

func (h *QuestionHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var input struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if input.Text == "" {
		respondError(w, http.StatusBadRequest, "Field 'text' is required")
		return
	}

	question := models.Question{
		Text:      input.Text,
		CreatedAt: time.Now(),
	}

	if err := h.repo.Create(&question); err != nil {
		logrus.WithError(err).Error("failed to create question")
		respondError(w, http.StatusInternalServerError, "Failed to create question")
		return
	}

	respondJSON(w, http.StatusCreated, question)
}

func (h *QuestionHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || id == 0 {
		respondError(w, http.StatusBadRequest, "Invalid question ID")
		return
	}

	question, err := h.repo.FindByIDWithAnswers(uint(id))
	if err != nil {
		respondError(w, http.StatusNotFound, "Question not found")
		return
	}

	respondJSON(w, http.StatusOK, question)
}

func (h *QuestionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || id == 0 {
		respondError(w, http.StatusBadRequest, "Invalid question ID")
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		respondError(w, http.StatusNotFound, "Question not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
