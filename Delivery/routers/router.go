package routers

import (
	infrastructure "LoanTrackerAPI/Infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(R *gin.Engine, env infrastructure.Config, client *mongo.Database) {

	userroute := R.Group("")
	UserRouter(userroute, *client, env)

}
