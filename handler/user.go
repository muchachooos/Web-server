package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func (s *Server) LoginHandler(context *gin.Context) {

	log, ok := context.GetQuery("login")
	if log == "" || !ok {
		context.Writer.WriteString("No login")
		return
	}

	pass, ok := context.GetQuery("password")
	if pass == "" || !ok {
		context.Writer.WriteString("No password")
		return
	}

	resultTable, err := s.Storage.GetLoginUserInDB(log, pass)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if len(resultTable) == 0 {
		context.Status(404)
		context.Writer.WriteString("Wrong login or password. Try again")
		return
	}

	jsonInByte, err := json.Marshal(resultTable)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}

func (s *Server) RegistrationHandler(context *gin.Context) {

	log, ok := context.GetQuery("login")
	if log == "" || !ok {
		context.Writer.WriteString("No login")
		return
	}

	pass, ok := context.GetQuery("password")
	if pass == "" || !ok {
		context.Writer.WriteString("No password")
		return
	}

	err := s.Storage.RegistrationUserInBD(log, pass)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) DeleteHandler(context *gin.Context) {

	log, ok := context.GetQuery("login")
	if log == "" || !ok {
		context.Writer.WriteString("No login")
		return
	}

	pass, ok := context.GetQuery("password")
	if pass == "" || !ok {
		context.Writer.WriteString("No password")
		return
	}

	isDeleted, err := s.Storage.DeleteUserFromDB(log, pass)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if isDeleted == false {
		context.Writer.WriteString("Something went wrong")
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) ChangeHandler(context *gin.Context) {

	log, ok := context.GetQuery("login")
	if log == "" || !ok {
		context.Writer.WriteString("No login")
		return
	}

	pass, ok := context.GetQuery("password")
	if pass == "" || !ok {
		context.Writer.WriteString("No password")
		return
	}

	newPass, ok := context.GetQuery("newPassword")
	if newPass == "" || !ok {
		context.Writer.WriteString("No new password")
		return
	}

	if pass == newPass {
		context.Writer.WriteString("Incorrect new password")
		return
	}

	isChanged, err := s.Storage.ChangePassUserInDB(log, pass, newPass)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if isChanged == false {
		context.Writer.WriteString("Something went wrong")
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}
