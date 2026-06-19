package security

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"justdrven.dev/storage/internal/api/common"
	"justdrven.dev/storage/shared/src/security/ratelimit"
)

var rateLimiter = ratelimit.NewRateLimiter()

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
	return rateLimiter.Allow(req.RemoteAddr)
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

		/**

		if !isAuthOk(r) {
			common.SetJsonType(w)

			json.NewEncoder(w).Encode(common.APIErrorResponse{
				Code:    403,
				Message: "Access denied!",
			})

			return
		}

			**/

		next.ServeHTTP(w, r)
	})
}
