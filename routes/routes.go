package routes

import (
	controller "github.com/degarzonm/go-back-portfolio/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
func UserRegistryRoutes(incomingRoutes *gin.Engine) {
	// Unknown users
	incomingRoutes.POST("/register", controller.SignUp())
	incomingRoutes.POST("/login", controller.Login())

}

func SoldierRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/my-fort", controller.GetMyFort())
	incomingRoutes.GET("/my-commander", controller.GetMyCommander())
	//incomingRoutes.GET("/today-agenda", controller.GetAgendaOfDay())
}

func OfficerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/officer/modify-agenda", controller.ModifyAgenda())
	//incomingRoutes.GET("/officer/check-general-plan", controller.CheckGeneralPlan())
	//incomingRoutes.GET("/officer/my-troops", controller.GetMyTroops())
}

func GeneralRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/general/create-fort", controller.CreateFort())
	incomingRoutes.POST("/general/set-fort-commander", controller.SetFortCommander())
	//incomingRoutes.POST("/general/create-plan", controller.CreateOrUpdateGeneralPlan())
	//incomingRoutes.PUT("/general/edit-fort/:id", controller.UpdateFort())
	//incomingRoutes.GET("/general/my-forts", controller.GetMyForts())
	//incomingRoutes.PUT("/general/transfer-fort/:general", controller.TransferFort())
	//incomingRoutes.GET("/general/my-troops/:fort", controller.GetMyTroopsByFort())
	//incomingRoutes.DELETE("/general/lost-fort/:id", controller.DeleteFort())
}

func RecruiterRoutes(incomingRoutes *gin.Engine) {
	// Recruiter
	incomingRoutes.PUT("/recruiter/ascend/", controller.AscendSoldier())
	incomingRoutes.DELETE("/recruiter/jubilate", controller.Jubilate())
	//incomingRoutes.PUT("/recruiter/transfer", controller.Transfer())
	//incomingRoutes.PUT("/recruiter/edit-soldier", controller.EditSoldier())
	//incomingRoutes.PUT("/recruiter/release-me", controller.ReleaseMe())
	//incomingRoutes.PUT("/recruiter/new-recruiter/:officer", controller.NewRecruiter())
}
