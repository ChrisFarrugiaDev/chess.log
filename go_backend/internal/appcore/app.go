package appcore

import (
	"chess_log/go_backend/internal/mailer"
	"chess_log/go_backend/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	DB       *pgxpool.Pool
	Models   models.Models
	Mailer   *mailer.SmtpMailer
	Validate *validator.Validate
}
