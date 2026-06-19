package security

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"justdrven.dev/storage/internal/api/common"
)

func Hash(value string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	if err != nil {
		return ""
	}
	return string(hash)
}

func Compare(hash, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}

func isRatelimitOk(req *http.Request) bool {
	return true
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !isRatelimitOk(r) {
			common.SetJsonType(w)

			json.NewEncoder(w).Encode(common.APIErrorResponse{
				Code:    429,
				Message: "Too many requests!",
			})

			return
		}

		next.ServeHTTP(w, r)
	})
}
