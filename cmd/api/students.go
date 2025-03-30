package main

import (
	"context"
	"net/http"

	"eduflow.dijam.io/internal/data/student"
	"eduflow.dijam.io/internal/validator"
)

func (app *application) CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Firstname string `json:"first_name"`
		Lastname  string `json:"last_name"`
		Email     string `json:"email"`
		Gender    string `json:"gender"`
		Image_url string `json:"image_url"`
		Status    string `json:"status"`
		ClassID   int64  `json:"class_id"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	s := student.Student{
		FirstName: input.Firstname,
		LastName:  input.Lastname,
		Email:     input.Email,
		Gender:    input.Gender,
		ImageURL:  input.Image_url,
		Status:    input.Status,
		ClassID:   input.ClassID,
	}

	v := validator.New()
	if student.ValidateStudent(v, &s); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.Models.StudentModel.AddStudent(context.Background(), &s)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	app.writeJSON(w, http.StatusCreated, envelope{"message": "student successfully added"}, nil)

}

func (app *application) GetStudentByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.getID(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	student, err := app.Models.StudentModel.GetStudentByID(context.Background(), id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusFound, envelope{"student": student}, nil)

}

func (app *application) GetAllStudentsHandler(w http.ResponseWriter, r *http.Request) {
	students, err := app.Models.StudentModel.GetAllStudents(context.Background())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"students": students}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
