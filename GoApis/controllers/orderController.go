package controllers

import (
	"fmt"

	"testApi/database"
	"testApi/models"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	fmt.Printf("The order is %v",order)
	database.InsertOrder(order)

	return c.Status(200).JSON("Inserted")
}


func GetOrders(c *fiber.Ctx) error {
	allOrders := database.GetAll();
	if allOrders == nil{
		return c.Status(404).JSON("No orders")
	}
	return c.Status(200).JSON(allOrders)
}

func GetOrder(c *fiber.Ctx) error {
	orderId := c.Params("id")
	fmt.Printf("The GetByID IS %v",orderId)
	result, err := database.GetById(orderId)
	fmt.Printf("the the result of get by id %v is ",result)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}
	
	var order bson.M
	err = result.Decode(&order)
	fmt.Printf("the order struct is %v",order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error decoding order",
		})
	}
	
	return c.JSON(order)
}


func UpdateOrder(c *fiber.Ctx)error{
	id := c.Params("id");
	var order models.Order
	c.BodyParser(&order)
	res,err := database.UpdateOrder(id,order);
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":"not updated",
		})
	}
	return c.Status(200).JSON(res)
}

func DeleteOrders(c *fiber.Ctx) error{

	isDeleted := database.Delete(c.Params("id"));
	if !isDeleted{
		return c.Status(404).JSON("Not deleted")
	}
	return c.Status(200).JSON("Deleted");
}