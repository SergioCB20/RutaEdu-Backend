package handlers

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/SergioCB20/RutaEdu-Backend/pkg/models"
)

var courses []models.Course

func ReadCurriculumGrid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var requestData struct {
		CsvData string `json:"csvData"`
	}
	err = json.Unmarshal(bodyBytes, &requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	csvReader := csv.NewReader(strings.NewReader(requestData.CsvData))
	records, err := csvReader.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, record := range records {
		level, err := strconv.Atoi(record[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		credits, err := strconv.ParseFloat(record[3],64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		course := models.Course{
			ID:         record[1],
			Name:       record[2],
			Credits:    credits,
			Requisites: strings.Fields(record[4]),
			Level:      level,
		}
		courses = append(courses, course)
	}
	jsonCourses, _ := json.Marshal(courses)
    w.Write(jsonCourses)
}

/*Prompt: Este pdf contiene información de la malla curricular de una carrera
en base a eso, devuelveme los datos de cada curso de esta malla en este formato:
codigo del curso, nombre del curso, créditos, número de requisitos,código de los cursos que son requisitos o el requisito en sí*/
