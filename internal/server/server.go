package server

import (
	"net/http"

	"github.com/shagabiev/Go-QA-Api/internal/handlers"
	"github.com/shagabiev/Go-QA-Api/internal/repository"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) http.Handler {
	mux := http.NewServeMux()

	qRepo := repository.NewQuestionRepository(db)
	aRepo := repository.NewAnswerRepository(db)

	qHandler := handlers.NewQuestionHandler(qRepo)
	aHandler := handlers.NewAnswerHandler(qRepo, aRepo)

	mux.HandleFunc("GET /questions", qHandler.List)
	mux.HandleFunc("POST /questions", qHandler.Create)
	mux.HandleFunc("GET /questions/{id}", qHandler.Get)
	mux.HandleFunc("DELETE /questions/{id}", qHandler.Delete)

	mux.HandleFunc("POST /questions/{id}/answers", aHandler.Create)
	mux.HandleFunc("GET /answers/{id}", aHandler.Get)
	mux.HandleFunc("DELETE /answers/{id}", aHandler.Delete)

	return mux
}
