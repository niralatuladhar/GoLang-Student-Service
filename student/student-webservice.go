package student

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/martini"
	"log"
)

// GetPath implements webservice.GetPath.
func (this *Student)  GetPath() string {
	// Associate this service with http://host:port/guestbook.
	return "/students"
}


// WebPost implements webservice.WebPost.
func (this Student) WebPost(params martini.Params,
req *http.Request) (int, string) {
	// Make sure Body is closed when we are done.
	defer req.Body.Close()

	// Read request body.
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusInternalServerError, "internal error"
	}

	if len(params) != 0 {
		// No keys in params. This is not supported.
		return http.StatusMethodNotAllowed, "method not allowed"
	}

	// Unmarshal entry sent by the user.
	var student Student
	err = json.Unmarshal(requestBody, &student)
	if err != nil {
		// Could not unmarshal entry.
		return http.StatusBadRequest, "invalid JSON data"
	}


	// Add entry provided by the user.

	this.AddStudents(student)

	// Everything is fine.
	return http.StatusOK, "new student created"
}



// WebGet implements webservice.WebGet.
func (this *Student) WebGet(params martini.Params) (int, string) {
	if len(params) == 0 {
		// No params. Return entire collection encoded as JSON.
		log.Println("calling GetAllStudents")
		encodedEntries, err := json.Marshal(this.GetAllStudents())
		if err != nil {
			// Failed encoding collection.
			return http.StatusInternalServerError, "internal error"
		}

		// Return encoded entries.
		return http.StatusOK, string(encodedEntries)
	} else {
		log.Println(params)
		// Convert id to integer.
		name := params["name"]
		log.Println("calling GetStudent")
		// Get entry identified by id.
		student, err := this.GetStudent(name)
		if err != nil {
			log.Println(err)
			// Entry not found.
			return http.StatusNotFound, "student not found"
		}

		// Encode entry in JSON.
		encodedEntry, err := json.Marshal(student)
		if err != nil {
			// Failed encoding entry.
			return http.StatusInternalServerError, "internal error"
		}

		// Return encoded entry.
		return http.StatusOK, string(encodedEntry)
	}
}

// WebDelete implements webservice.WebDelete.
func (g *Student) WebDelete(params martini.Params) (int, string) {

	return http.StatusOK, "entry deleted"
}