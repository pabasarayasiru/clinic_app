package controllers

import (
	"clinic_app/models"
	"clinic_app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateVisit(c *gin.Context) {
	var visit models.Visit

	if err := c.ShouldBindJSON(&visit); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Get doctor from JWT
	doctorID := c.GetUint("user_id")
	visit.DoctorID = doctorID

	// Save Visit FIRST
	err := services.CreateVisit(DB, visit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// AI Parsing
	aiResult, err := services.ParseWithAI(visit.NotesRawInput)
	if err != nil {
		c.JSON(500, gin.H{"error": "AI parsing failed"})
		return
	}

	// Save AI Logs (IMPORTANT for assignment)
	services.AddAIParsingLog(DB, models.AIParsingLog{
		VisitID:  visit.ID,
		RawInput: visit.NotesRawInput,
	})

	// Save Notes
	for _, note := range aiResult.Notes {
		services.AddClinicNote(DB, models.ClinicNote{
			VisitID:  visit.ID,
			NoteText: note,
		})
	}

	// Save Drugs
	for _, d := range aiResult.Drugs {
		services.AddDrugPrescription(DB, models.DrugPrescription{
			VisitID:  visit.ID,
			DrugName: d.Name,
			Dosage:   d.Dosage,
		})
	}

	// Save Labs
	for _, l := range aiResult.Labs {
		services.AddLabTest(DB, models.LabTest{
			VisitID:  visit.ID,
			TestName: l,
		})
	}

	c.JSON(200, gin.H{
		"message": "Visit created with AI parsing",
	})
}

func GetVisit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	visit, err := services.GetVisitByID(DB, uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	c.JSON(200, visit)
}

func GetDoctorVisits(c *gin.Context) {
	doctorID := c.GetUint("user_id")

	visits, _ := services.GetVisitsByDoctor(DB, doctorID)
	c.JSON(200, visits)
}
