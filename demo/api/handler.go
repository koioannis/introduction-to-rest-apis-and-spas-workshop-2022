package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterHandler(e *echo.Echo, r Repository) {
	h := handler{notes: r}

	e.GET("/notes", h.GetAll)
	e.GET("/notes/:id", h.Get)
	e.POST("/notes", h.Create)
	e.PUT("/notes/:id", h.Update)
	e.DELETE("/notes/:id", h.Delete)
}

type handler struct {
	notes Repository
}

func (h *handler) Get(ctx echo.Context) error {
	param := ctx.Param("id")
	if param == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is missing")
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a valid integer")
	}

	note, err := h.notes.Get(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("note with id %d not found", id))
		}

		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return ctx.JSON(http.StatusOK, note)
}

func (h *handler) GetAll(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.notes.GetAll())
}

func (h *handler) Update(ctx echo.Context) error {
	param := ctx.Param("id")
	if param == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is missing")
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a valid integer")
	}

	request := struct {
		Text string `json:"text" validate:"required"`
	}{}
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err = ctx.Validate(request); err != nil {
		return err
	}

	note := Note{
		ID:   id,
		Text: request.Text,
	}
	if err := h.notes.Update(id, note); err != nil {
		if errors.Is(err, ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("note with id %d not found", id))
		}

		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *handler) Create(ctx echo.Context) error {
	request := struct {
		Text string `json:"text" validate:"required"`
	}{}
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := ctx.Validate(request); err != nil {
		return err
	}

	note := Note{
		Text: request.Text,
	}
	note = h.notes.Create(note)

	return ctx.JSON(http.StatusCreated, note)
}

func (h *handler) Delete(ctx echo.Context) error {
	param := ctx.Param("id")
	if param == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is missing")
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a valid integer")
	}

	if err := h.notes.Delete(id); err != nil {
		if errors.Is(err, ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("note with id %d not found", id))
		}

		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return ctx.NoContent(http.StatusOK)
}
