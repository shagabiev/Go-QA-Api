package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/shagabiev/Go-QA-Api/internal/models"
	"github.com/shagabiev/Go-QA-Api/internal/repository"
	"github.com/sirupsen/logrus"
)

type AnswerHandler struct {
	qRepo *repository.QuestionRepository
	aRepo *repository.AnswerRepository
}

func NewAnswerHandler(qRepo *repository.QuestionRepository, aRepo *repository.AnswerRepository) *AnswerHandler {
	return &AnswerHandler{qRepo: qRepo, aRepo: aRepo}
}

func (h *AnswerHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	qidStr := r.PathValue("id")
	qid, err := strconv.ParseUint(qidStr, 10, 32)
	if err != nil || qid == 0 {
		respondError(w, http.StatusBadRequest, "Invalid question ID")
		return
	}

	if _, err := h.qRepo.FindByID(uint(qid)); err != nil {
		respondError(w, http.StatusNotFound, "Question not found")
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
		respondError(w, http.StatusBadRequest, "text is required")
		return
	}

	userID := uuid.New()

	answer := models.Answer{
		QuestionID: uint(qid),
		UserID:     userID,
		Text:       input.Text,
		CreatedAt:  time.Now(),
	}

	if err := h.aRepo.Create(&answer); err != nil {
		logrus.WithError(err).Error("failed to create answer")
		respondError(w, http.StatusInternalServerError, "Failed to create answer")
		return
	}

	respondJSON(w, http.StatusCreated, answer)
}

func (h *AnswerHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || id == 0 {
		respondError(w, http.StatusBadRequest, "Invalid answer ID")
		return
	}

	answer, err := h.aRepo.FindByID(uint(id))
	if err != nil {
		respondError(w, http.StatusNotFound, "Answer not found")
		return
	}

	respondJSON(w, http.StatusOK, answer)
}

func (h *AnswerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || id == 0 {
		respondError(w, http.StatusBadRequest, "Invalid answer ID")
		return
	}

	if err := h.aRepo.Delete(uint(id)); err != nil {
		respondError(w, http.StatusNotFound, "Answer not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
