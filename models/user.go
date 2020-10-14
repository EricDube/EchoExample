// This package name isnt the best but its an example of a way
// to group parts of the application by packages
package models

import (
	"github.com/ericdube/echoexample/logger"
)

// User struct
type User struct {
	Name string `json:"name"`
	UserID int64 `json:"userID"`
	IsSaved bool
}

// SaveUser
func SaveUser(u *User){
	//Save user to database
	u.IsSaved = true

	// This is why naming is important
	// Instead of putting the Logger inside a struct called Logger
	// and in a package called logger, put the logger in a separate repo
	// that would contain common logging logic to be reused by other
	// applications.
	logger.Logger.Logger.Info(u)
}