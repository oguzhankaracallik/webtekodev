package controllers

import (
	"github.com/gofiber/fiber/v2"
	"oguzodev/database"
	"oguzodev/models"
)

type CommentC struct{}

func (c CommentC) GetComments(ctx *fiber.Ctx) error {
	var comments []models.Comment
	id := ctx.Params("post-id")
	err := database.Conn.Find(&comments).Where("post_id = ?", id).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(comments)
	return nil
}

func (c CommentC) CreateComment(ctx *fiber.Ctx) error {
	var comment models.Comment
	err := ctx.BodyParser(&comment)
	if err != nil {
		panic(err.Error())
	}
	err = database.Conn.Create(&comment).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(comment)
	return nil
}

func (c CommentC) DeleteComment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var comment models.Comment
	err := database.Conn.Delete(&comment, id).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(comment)
	return nil
}
