package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/pborman/uuid"
)

// ServeJson write data as the type of json to a http response writer
// and it returns an error when the json encoding failed.
func ServeJson(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

// ExtractJson decode a json type http request body to a golang interface{} dst
// and to bind the decoded data to this param
// the dst must be an point.
func ExtractJson(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

// EnterLog log a base info of a http request
func EnterLog(r *http.Request) {
	log.Printf("|---REQUEST---[%-4s] %-10s | %-20s |\n", r.Method, r.Host, r.URL.Path)
}

// Password first get a salt, then crypto the user password with this salt, and
// returns the crypto password and the salt
func Password(password string) (pwd string, salt string) {
	salt = uuid.New()
	pwd = Hash(password, salt)
	return
}

// Hash return a crypto string whitch is cryptoed from the param password
// and the param salt
func Hash(password, salt string) string {
	return Md5(strings.NewReader(utils.Md5(strings.NewReader(password), true)+salt), true)
}

// Md5 function
func Md5(reader io.Reader, upper bool) string {
	hash := md5.New()
	io.Copy(hash, reader)
	if upper {
		return fmt.Sprintf("%X", hash.Sum(nil))
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}
