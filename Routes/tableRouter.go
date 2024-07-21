package routes

import(
	"github.com/gin-gonic/gin"
	controller "restaurant-management/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/tables",controller.GetTables()) 
	// this is list of all tables, that consist below coloumns
	
	incomingRoutes.GET("/tables/:table_id",controller.GetTable())
	incomingRoutes.POST("/tables",controller.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id",controller.UpdateTable())
}
