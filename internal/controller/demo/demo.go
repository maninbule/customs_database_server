package demo

import (
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/internal/dao/mysql/FaceEmbedding"
	"github.com/customs_database_server/internal/dao/mysql/FaceResult"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func ShowAllFaceResult(c *gin.Context) {
	faces := mysqlFaceResult.GetAllFace()

	var htmlBuilder strings.Builder
	htmlBuilder.WriteString("<html><body><table>")
	for _, face := range faces {
		htmlBuilder.WriteString("<tr>")
		htmlBuilder.WriteString("<td>" + *face.Name + "</td>")
		htmlBuilder.WriteString("<td>" + face.FaceTime.String() + "</td>")
		htmlBuilder.WriteString("<td>" + *face.CameraID + "</td>")
		htmlBuilder.WriteString("<td><img src=\"" + config.HttpIP + *face.FaceImgCorrect + "\" width=\"200\" height=\"200\"></td>")
		htmlBuilder.WriteString("<td><img src=\"" + config.HttpIP + *face.FaceImgPredict + "\" width=\"200\" height=\"200\"></td>")
		htmlBuilder.WriteString("</tr>")
	}
	htmlBuilder.WriteString("</table></body></html>")

	// 返回 HTML
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, htmlBuilder.String())
}

func ShowAllFaceEmbedding(c *gin.Context) {
	faces := mysqlFaceEmbedding.GetAllFace()

	var htmlBuilder strings.Builder
	htmlBuilder.WriteString("<html><body><table>")
	for _, face := range faces {
		htmlBuilder.WriteString("<tr>")
		htmlBuilder.WriteString("<td>" + strconv.Itoa(int(*face.FaceId)) + "</td>")
		htmlBuilder.WriteString("<td>" + *face.Name + "</td>")
		htmlBuilder.WriteString("<td><img src=\"" + config.HttpIP + *face.FaceImgURL + "\" width=\"200\" height=\"200\"></td>")
		htmlBuilder.WriteString("</tr>")
	}
	htmlBuilder.WriteString("</table></body></html>")

	// 返回 HTML
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, htmlBuilder.String())
}

func ShowHTMLFaceResult(c *gin.Context) {
	c.HTML(http.StatusOK, "demo_face_result.html", nil)
}
