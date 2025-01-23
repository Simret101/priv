package controller

import (
    "net/http"

    "auth/config" // Import the config package for error handling
    "user/domain"

    "go.mongodb.org/mongo-driver/bson/primitive"

    "github.com/gin-gonic/gin"
)

type UserController struct {
    UserUsecase domain.User_Usecase_interface
}

func NewUserController(usecase domain.User_Usecase_interface) *UserController {
    return &UserController{UserUsecase: usecase}
}

// GetOneUser retrieves a single user by ID
// @Summary Get a single user by ID
// @Description Get a single user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/user/{id} [get]
func (controller *UserController) GetOneUser() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        id := ctx.Param("id")

        user, err := controller.UserUsecase.GetOneUser(id)
        if err != nil {
            ctx.IndentedJSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
            return
        }

        ctx.IndentedJSON(http.StatusOK, gin.H{"data": user})
    }
}

// GetUsers retrieves all users
// @Summary Get all users
// @Description Get all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} domain.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/user [get]
func (controller *UserController) GetUsers() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        users, err := controller.UserUsecase.GetUsers()

        if err != nil {
            ctx.IndentedJSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
            return
        }
        ctx.IndentedJSON(http.StatusOK, gin.H{"data": users})
    }
}

// UpdateUser updates a user by ID
// @Summary Update a user by ID
// @Description Update a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body domain.User true "User"
// @Success 200 {object} domain.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/user/{id} [put]
func (controller *UserController) UpdateUser() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        update_user := domain.UpdateUser{}
        idstr := ctx.Param("id")
        id, err := primitive.ObjectIDFromHex(idstr)

        if err != nil {
            ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": config.ErrInvalidToken.Error()})
            return
        }
        var user domain.User
        if err := ctx.BindJSON(&user); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request: " + err.Error()})
            return
        }
        user.ID = id

        updateuser, err := controller.UserUsecase.UpdateUser(id.Hex(), update_user)
        if err != nil {
            ctx.JSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
            return
        }

        ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully!", "user": updateuser})
    }
}

// DeleteUser deletes a user by ID
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 202 {string} string "accepted!"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/user/{id} [delete]
func (controller *UserController) DeleteUser() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        id := ctx.Param("id")
        err := controller.UserUsecase.DeleteUser(id)
        if err != nil {
            ctx.IndentedJSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
            return
        }
        ctx.IndentedJSON(http.StatusAccepted, gin.H{"message": "accepted!"})
    }
}

// FilterUser filters users based on query parameters
// @Summary Filter users
// @Description Filter users based on query parameters
// @Tags Users
// @Accept json
// @Produce json
// @Param filter query string false "Filter"
// @Success 200 {array} domain.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/users/filter [get]
func (controller *UserController) FilterUser() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        query := ctx.Request.URL.Query()
        filter := make(map[string]string)
        for key, value := range query {
            filter[key] = value[0]
        }

        users, err := controller.UserUsecase.FilterUser(filter)
        if err != nil {
            ctx.IndentedJSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
            return
        }
        ctx.IndentedJSON(http.StatusOK, gin.H{"data": users})
    }
}

// PromoteUser promotes a user by ID
// @Summary Promote a user by ID
// @Description Promote a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/promote/{id} [put]
func (controller *UserController) PromoteUser() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        id := ctx.Param("id")

        user, err := controller.UserUsecase.PromoteUser(id)
        if err != nil {
            ctx.IndentedJSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
            return
        }

        ctx.IndentedJSON(http.StatusOK, gin.H{"data": user})
    }
}

// DemoteUser demotes a user by ID
// @Summary Demote a user by ID
// @Description Demote a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/demote/{id} [put]
func (controller *UserController) DemoteUser() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        id := ctx.Param("id")

        user, err := controller.UserUsecase.DemoteUser(id)
        if err != nil {
            ctx.IndentedJSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
            return
        }

        ctx.IndentedJSON(http.StatusOK, gin.H{"data": user})
    }
}

// UpdatePassword updates a user's password by ID
// @Summary Update a user's password by ID
// @Description Update a user's password by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param password body domain.UpdatePassword true "Password"
// @Success 202 {object} domain.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/user/{id}/password [put]
func (controller *UserController) UpdatePassword() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var update_password domain.UpdatePassword
        id := ctx.Param("id")
        if err := ctx.BindJSON(&update_password); err != nil {
            ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        response_user, err := controller.UserUsecase.UpdatePassword(id, update_password)
        if err != nil {
            ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        ctx.IndentedJSON(http.StatusAccepted, gin.H{"data": response_user})
    }
}