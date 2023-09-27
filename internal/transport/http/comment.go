package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/souviks72/go-rest-api/internal/comment"
)

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(context.Context, string) (comment.Comment, error)
	UpdateComment(context.Context, string, comment.Comment) (comment.Comment, error)
	DeleteComment(context.Context, string) error
}

type Response struct {
	Message string
}

type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"Author" validate:"required"`
	Body   string `json:"Body" validate:"required"`
}

func convertPostCommentRequestToComment(c PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   c.Slug,
		Author: c.Author,
		Body:   c.Body,
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var reqCmt PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&reqCmt); err != nil {
		log.Println(err)
		return
	}

	validate := validator.New()
	err := validate.Struct(reqCmt)
	if err != nil {
		http.Error(w, "not a valid comment", http.StatusBadRequest)
	}

	convertedCmt := convertPostCommentRequestToComment(reqCmt)
	cmt, err := h.Service.PostComment(r.Context(), convertedCmt)
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		log.Println(err)
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		log.Println(err)
		return
	}

	updatedCmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(updatedCmt); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteComment(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
