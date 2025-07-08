package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"john_doe":  "password1",
	"alice_bob": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Utility to write JSON or plain responses
func writeResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}

// Utility to extract and validate JWT token
func getClaimsFromToken(r *http.Request) (*Claims, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}

	tokenStr := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

// Home handler - Auth protected
func Home(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaimsFromToken(r)
	if err != nil {
		writeResponse(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	writeResponse(w, http.StatusOK, "Welcome "+claims.Username)
}

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		writeResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "Could not create token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenStr,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   true,
	})

	writeResponse(w, http.StatusOK, "Login successful")
}

// Refresh handler
func Refresh(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaimsFromToken(r)
	if err != nil {
		writeResponse(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	// Only allow refresh if token is close to expiring (within 30 seconds)
	if time.Until(time.Unix(claims.ExpiresAt, 0)) > 30*time.Second {
		writeResponse(w, http.StatusBadRequest, "Token not eligible for refresh yet")
		return
	}

	claims.ExpiresAt = time.Now().Add(5 * time.Minute).Unix()
	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "Could not refresh token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenStr,
		Expires:  time.Unix(claims.ExpiresAt, 0),
		HttpOnly: true,
		Secure:   true,
	})

	writeResponse(w, http.StatusOK, "Token refreshed successfully")
}

// Logout handler
func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
	})
	writeResponse(w, http.StatusOK, "Logout successful")
}
