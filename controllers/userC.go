package controllers

import (
	"github.com/gofiber/fiber/v2"
	"oguzodev/database"
	"oguzodev/models"
)

type UserC struct{}

func (c UserC) GetUsers(ctx *fiber.Ctx) error {
	var users []models.User

	err := database.Conn.Find(&users).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(users)
	return nil
}

func (c UserC) GetUser(ctx *fiber.Ctx) error {
	var user models.User
	err := database.Conn.Find(&user).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(user)
	return nil
}

func (c UserC) CreateUser(ctx *fiber.Ctx) error {
	var user models.User
	err := database.Conn.Create(&user).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(user)
	return nil
}

func (c UserC) FollowUser(ctx *fiber.Ctx) error {
	var user models.Follower
	err := database.Conn.Create(&user).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(user)
	return nil
}

func (c UserC) UpdateUser(ctx *fiber.Ctx) error {
	var user models.User
	err := database.Conn.Model(&user).Updates(&user).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(user)
	return nil
}

func (c UserC) DeleteUser(ctx *fiber.Ctx) error {
	var user models.User
	err := database.Conn.Delete(&user).Error
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(user)
	return nil
}
