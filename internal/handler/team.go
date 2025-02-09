package handler

import (
	teamservice "github.com/Imtiaz246/Thesis-Management-System/internal/service/team"
)

type TeamHandler struct {
	Handler
	teamService teamservice.Service
}
