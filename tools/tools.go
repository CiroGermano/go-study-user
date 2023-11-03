package tools

import (
	"time"
	"fmt"	
)

func CloseMySQL() string {
	t := time.Now()
	return fmt.Sprintf("MySQL connection closed at %s", t.Format("2006-01-02 15:04:05"))
}