package main

import (
	"fmt"
	"strconv"
)

func GetQuiz(input string) (Quiz, string) {
	var quiz Quiz

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return quiz, "invalid input"
	}

	db, err := GetDB()
	if db == nil {
		fmt.Println(err)
		return quiz, "error connecting to db"
	}

	query := "SELECT id, name, description FROM quiz WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&quiz.Id, &quiz.Name, &quiz.Desc)
	if err != nil {
		fmt.Println(err)
		return quiz, err.Error()
	}

	fmt.Println("Response is {}", quiz)
	return quiz, ""
}

func CreateQuiz(quiz Quiz) Quiz {

	db, _ := GetDB()
	query := "INSERT INTO quiz(name, description) VALUES('" + quiz.Name + "', '" + quiz.Desc + "')"
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return quiz
	}

	id, _ := res.LastInsertId()
	quiz.Id = int(id)

	return quiz
}

func EnterQuestion(input Question) (bool, int) {

	db, err := GetDB()
	if err != nil {
		fmt.Println(err)
		return false, 0
	}

	query := "insert into question(name, options, correct_option, quiz, points) values(?, ?, ?, ?, ?)"
	res, err := db.Exec(query, input.Name, input.Options, input.CorrectOption, input.Quiz, input.Points)
	if err != nil {
		fmt.Println(err)
		return false, 0
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return false, 0
	}

	return true, int(id)
}

func GetQuestion(input string) (Question, string) {
	var quiz Question

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return quiz, "invalid input"
	}

	db, err := GetDB()
	if db == nil {
		fmt.Println(err)
		return quiz, "error connecting to db"
	}

	query := "SELECT id, name, options, correct_option as correctOption, quiz, points FROM question WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&quiz.Id, &quiz.Name, &quiz.Options, &quiz.CorrectOption, &quiz.Quiz, &quiz.Points)
	if err != nil {
		fmt.Println(err)
		return quiz, err.Error()
	}

	fmt.Println("Response is {}", quiz)
	return quiz, ""
}

func GetAllQuestion(input string) ([]Question, string) {
	var quez []Question

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return quez, "invalid input"
	}

	db, err := GetDB()
	if db == nil {
		fmt.Println(err)
		return quez, "error connecting to db"
	}

	query := "SELECT question.id, question.name, question.options, question.correct_option as correctOption, question.quiz, question.points FROM quiz,question WHERE quiz.id = question.id and question.quiz = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println(err)
		return quez, err.Error()
	}

	for rows.Next() {
		var quiz Question
		err = rows.Scan(&quiz.Id, &quiz.Name, &quiz.Options, &quiz.CorrectOption, &quiz.Quiz, &quiz.Points)
		if err != nil {
			fmt.Println(err)
			return quez, err.Error()
		}
		quez = append(quez, quiz)
	}

	fmt.Println("Response is {}", quez)
	return quez, ""
}
