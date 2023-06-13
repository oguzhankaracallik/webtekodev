package controllers

import (
	"github.com/gofiber/fiber/v2"
	"oguzodev/database"
	"oguzodev/models"
)

type PostC struct{}

func (c PostC) GetPosts(ctx *fiber.Ctx) error {
	var posts []models.Post
	err := database.Conn.Find(&posts).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(posts)
	return nil
}

func (c PostC) GetPost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var post models.Post
	err := database.Conn.Find(&post, id).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(post)
	return nil
}

func (c PostC) CreatePost(ctx *fiber.Ctx) error {
	var post models.Post
	err := ctx.BodyParser(&post)
	if err != nil {
		panic(err.Error())
	}
	err = database.Conn.Create(&post).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(post)
	return nil
}

func (c PostC) UpdatePost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var post models.Post
	err := ctx.BodyParser(&post)
	if err != nil {
		panic(err.Error())
	}
	err = database.Conn.Model(&post).Where("post_id = ?", id).Updates(&post).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(post)
	return nil
}

func (c PostC) DeletePost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var post models.Post
	err := database.Conn.Delete(&post, id).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(post)
	return nil
}

func (c PostC) LikePost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var post models.Post
	err := database.Conn.Find(&post, id).Error
	if err != nil {
		panic(err.Error())
	}
	post.PostLike += 1
	err = database.Conn.Model(&post).Where("post_id = ?", id).Updates(&post).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(post)
	return nil
}
