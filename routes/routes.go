package routes

import (
    controller "github.com/degarzonm/go-back-portfolio/controllers"

    "github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
    // Unknown users
    incomingRoutes.POST("/register", controller.SignUp())
    incomingRoutes.POST("/login", controller.Login())
/*
    // Soldier | Officer
    incomingRoutes.GET("/my-fort", controller.GetMyFort())
    incomingRoutes.GET("/my-commander", controller.GetMyCommander())
    incomingRoutes.GET("/check-agenda", controller.GetAgendaOfDay())

    // Officer
    incomingRoutes.PUT("/officer/modify-agenda", controller.ModifyAgenda())
    incomingRoutes.GET("/officer/check-general-plan", controller.CheckGeneralPlan())
    incomingRoutes.GET("/officer/my-troops", controller.GetMyTroops())

    // General
    incomingRoutes.POST("/general/create-plan", controller.CreateOrUpdateGeneralPlan())
    incomingRoutes.POST("/general/create-fort", controller.CreateFort())
    incomingRoutes.PUT("/general/edit-fort/:id", controller.UpdateFort())
    incomingRoutes.GET("/general/my-forts", controller.GetMyForts())
    incomingRoutes.PUT("/general/set-fort-commander/:officer/:fort", controller.SetFortCommander())
    incomingRoutes.PUT("/general/transfer-fort/:general", controller.TransferFort())
    incomingRoutes.GET("/general/my-troops/:fort", controller.GetMyTroopsByFort())
    incomingRoutes.DELETE("/general/lost-fort/:id", controller.DeleteFort())

    // Recruiter
    incomingRoutes.PUT("/recruiter/ascend/:role", controller.Ascend())
    incomingRoutes.PUT("/recruiter/transfer/:role/:fort", controller.Transfer())
    incomingRoutes.DELETE("/recruiter/jubilate/:role", controller.Jubilate())
    incomingRoutes.PUT("/recruiter/edit-soldier/:role", controller.EditSoldier())
    incomingRoutes.PUT("/recruiter/release-me", controller.ReleaseMe())
    incomingRoutes.PUT("/recruiter/new-recruiter/:officer", controller.NewRecruiter())*/
}