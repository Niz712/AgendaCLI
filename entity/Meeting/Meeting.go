package Meeting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"AgendaCLI/entity/User"
)

type Meeting struct {
	Title        string
	Sponsor      string
	Participants []string
	StartTime    time.Time
	EndTime      time.Time
	Id           string
}