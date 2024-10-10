package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "Rizki Ramadhan's Portfolio",
        "projects": []Project{
            {Name: "Project One", Description: "Description of project one.", Link: "https://github.com/yourusername/project-one"},
            {Name: "Project Two", Description: "Description of project two.", Link: "https://github.com/yourusername/project-two"},
        },
    })
}

type Project struct {
    Name        string
    Description string
    Link        string
}
