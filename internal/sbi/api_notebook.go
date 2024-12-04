package sbi

import (
	"net/http"

	"github.com/andy89923/nf-example/internal/logger"
	"github.com/gin-gonic/gin"
)

func (s *Server) getNotebookRoute() []Route {
	return []Route{
		{
			Name:    "Show Note",
			Method:  http.MethodGet,
			Pattern: "/:Title",
			APIFunc: s.HTTPShowNote,
			// Use
			// curl -X GET http://127.0.0.163:8000/spyfamily/ -w "\n"
		},
		{
			Name:    "SPYxFAMILY Character",
			Method:  http.MethodGet,
			Pattern: "/character/:Name",
			APIFunc: s.HTTPSerchSpyFamilyCharacter,
			// Use
			// curl -X GET http://127.0.0.163:8000/spyfamily/Anya -w "\n"
			// "Character: Anya Forger"
		},
	}
}

func (s *Server) HTTPShowNote(c *gin.Context) {
	logger.SBILog.Infof("In HTTPShowNote")

	targetName := c.Param("Title")
	if targetName == "" {
		c.String(http.StatusBadRequest, "No name provided")
		return
	}

	s.Processor().FindSpyFamilyCharacterName(c, targetName)
}

func (s *Server) HTTPSerchSpyFamilyCharacter(c *gin.Context) {
	logger.SBILog.Infof("In HTTPSerchCharacter")

	targetName := c.Param("Name")
	if targetName == "" {
		c.String(http.StatusBadRequest, "No name provided")
		return
	}

	s.Processor().FindSpyFamilyCharacterName(c, targetName)
}
