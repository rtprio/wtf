package wtf

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

// DateFormat defines the format we expect to receive dates from BambooHR in
const DateFormat = "2006-01-02"

func CenterText(str string, width int) string {
	return fmt.Sprintf("%[1]*s", -width, fmt.Sprintf("%[1]*s", (width+len(str))/2, str))
}

func ExecuteCommand(cmd *exec.Cmd) string {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Sprintf("A: %v\n", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Sprintf("B: %v\n", err)
	}

	var str string
	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += string(b)
	}

	cmd.Wait()

	return str
}

func Exclude(strs []string, val string) bool {
	for _, str := range strs {
		if val == str {
			return false
		}
	}
	return true
}

func NameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	return strings.Title(strings.Replace(parts[0], ".", " ", -1))
}

func NamesFromEmails(emails []string) []string {
	names := []string{}

	for _, email := range emails {
		names = append(names, NameFromEmail(email))
	}

	return names
}

func PrettyDate(dateStr string) string {
	newTime, _ := time.Parse(DateFormat, dateStr)
	return fmt.Sprint(newTime.Format("Jan 2, 2006"))
}

func Today() string {
	localNow := time.Now().Local()
	return localNow.Format("2006-01-02")
}

func UnixTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}
