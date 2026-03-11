package user

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	resource *Resource
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{
		resource: NewResource(db),
	}
}

// GetAllUsers handles GET /api/users
func (h *Handler) GetAllUsers(c fiber.Ctx) error {
	log.Info("Fetching all users")
	
	users, err := h.resource.GetAllUsers()
	if err != nil {
		log.Errorf("Failed to fetch users: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
			"details": err.Error(),
		})
	}

	log.Infof("Successfully fetched %d users", len(users))
	return c.JSON(fiber.Map{
		"success": true,
		"data":    users,
		"count":   len(users),
	})
}

// GetUserByID handles GET /api/users/:id
func (h *Handler) GetUserByID(c fiber.Ctx) error {
	idParam := c.Params("id")
	log.Infof("Fetching user with ID: %s", idParam)
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("Invalid user ID format: %s", idParam)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := h.resource.GetUserByID(id)
	if err != nil {
		log.Errorf("User not found with ID %d: %v", id, err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	log.Infof("Successfully fetched user: %s", user.Email)
	return c.JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

// CreateUser handles POST /api/users
func (h *Handler) CreateUser(c fiber.Ctx) error {
	var user User
	
	if err := c.Bind().Body(&user); err != nil {
		log.Errorf("Invalid request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
			"details": err.Error(),
		})
	}

	log.Infof("Creating new user: %s (%s)", user.Name, user.Email)

	if err := h.resource.CreateUser(&user); err != nil {
		log.Errorf("Failed to create user: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to create user",
			"details": err.Error(),
		})
	}

	log.Infof("Successfully created user with ID: %d", user.ID)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"data":    user,
	})
}

// UpdateUser handles PUT /api/users/:id
func (h *Handler) UpdateUser(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("Invalid user ID format: %s", idParam)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var user User
	if err := c.Bind().Body(&user); err != nil {
		log.Errorf("Invalid request body for user update: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
			"details": err.Error(),
		})
	}

	log.Infof("Updating user with ID: %d", id)

	if err := h.resource.UpdateUser(id, &user); err != nil {
		log.Errorf("Failed to update user %d: %v", id, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to update user",
			"details": err.Error(),
		})
	}

	// Fetch updated user to return
	updatedUser, err := h.resource.GetUserByID(id)
	if err != nil {
		log.Warnf("User updated but failed to fetch updated data: %v", err)
		return c.JSON(fiber.Map{
			"success": true,
			"message": "User updated successfully",
		})
	}

	log.Infof("Successfully updated user: %s", updatedUser.Email)
	return c.JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",
		"data":    updatedUser,
	})
}

// DeleteUser handles DELETE /api/users/:id
func (h *Handler) DeleteUser(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("Invalid user ID format: %s", idParam)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	log.Infof("Deleting user with ID: %d", id)

	if err := h.resource.DeleteUser(id); err != nil {
		log.Errorf("Failed to delete user %d: %v", id, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
			"details": err.Error(),
		})
	}

	log.Infof("Successfully deleted user with ID: %d", id)
	return c.JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully",
	})
}

// SetupRoutes configures all user routes
func SetupRoutes(app fiber.Router, db *sqlx.DB) {
	handler := NewHandler(db)
	
	userGroup := app.Group("/users")
	
	userGroup.Get("/", handler.GetAllUsers)
	userGroup.Get("/:id", handler.GetUserByID)
	userGroup.Post("/", handler.CreateUser)
	userGroup.Put("/:id", handler.UpdateUser)
	userGroup.Delete("/:id", handler.DeleteUser)
}