package post

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type API struct {
	repository *Repository
}

func New(db *gorm.DB) *API {
	return &API{
		repository: NewRepository(db),
	}
}

// List godoc
//
//	@summary		List posts
//	@description	List posts
//	@tags			posts
//	@accept			json
//	@produce		json
//	@success		200	{array}		DTO
//	@failure		500	{object}	err.Error
//	@router			/posts [get]
func (a *API) List(w http.ResponseWriter, r *http.Request) {
	posts, err := a.repository.List()
	if err != nil {
		// TODO: handle later
		return
	}

	if len(posts) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(posts.ToDto()); err != nil {
		// TODO: handle later
		return
	}
}

// Create godoc
//
//	@summary		Create post
//	@description	Create post
//	@tags			post
//	@accept			json
//	@produce		json
//	@param			body	body	Form	true	"Post form"
//	@success		201
//	@failure		400	{object}	err.Error
//	@failure		422	{object}	err.Errors
//	@failure		500	{object}	err.Error
//	@router			/post [post]
func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	form := &Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		// TODO: handle later
		return
	}

	newPost := form.ToModel()
	newPost.ID = uuid.New()

	_, err := a.repository.Create(newPost)
	if err != nil {
		// TODO: handle later
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Read godoc
//
//	@summary		Read post
//	@description	Read post
//	@tags			post
//	@accept			json
//	@produce		json
//	@param			id	path		string	true	"Post ID"
//	@success		200	{object}	DTO
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		500	{object}	err.Error
//	@router			/post/{id} [get]
func (a *API) Read(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// TODO: handle later
		return
	}

	post, err := a.repository.Read(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// TODO: handle later
		return
	}

	dto := post.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		// TODO: handle later
		return
	}
}

// Update godoc
//
//	@summary		Update post
//	@description	Update post
//	@tags			post
//	@accept			json
//	@produce		json
//	@param			id		path	string	true	"Post ID"
//	@param			body	body	Form	true	"Post form"
//	@success		200
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		422	{object}	err.Errors
//	@failure		500	{object}	err.Error
//	@router			/post/{id} [put]
func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// TODO: handle later
		return
	}

	form := &Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		// TODO: handle later
		return
	}

	post := form.ToModel()
	post.ID = id

	rows, err := a.repository.Update(post)
	if err != nil {
		// TODO: handle later
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// Delete godoc
//
//	@summary		Delete post
//	@description	Delete post
//	@tags			post
//	@accept			json
//	@produce		json
//	@param			id	path	string	true	"Post ID"
//	@success		200
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		500	{object}	err.Error
//	@router			/post/{id} [delete]
func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// TODO: handle later
		return
	}

	rows, err := a.repository.Delete(id)
	if err != nil {
		// TODO: handle later
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
