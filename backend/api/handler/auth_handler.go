// api/handler/auth_handler.go
package handler // パッケージ名を handler にします

import (
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
 )

// --- グローバルではなく、ハンドラー構造体の一部として定義 ---
type AuthHandler struct {
	// 将来的にDB接続などが必要になったら、ここにフィールドを追加する
	// db *sql.DB
}

// NewAuthHandler はAuthHandlerのインスタンスを生成します
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// --- データとJWT関連の定義 (本来はより適切な場所に移動) ---
var users = make(map[string]string)
var jwtKey = []byte("my_secret_key")
var blacklistedTokens = make(map[string]bool) // ブラックリストされたトークンを保存

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
func (h *AuthHandler) Register(c *gin.Context) {
	// (処理内容は同じ)
	var req struct { Username string `json:"username"`; Password string `json:"password"` }
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"} ); return }
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"} ); return }
	users[req.Username] = string(hashedPassword)
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"} )
}

// Login はログイン処理を行います
func (h *AuthHandler) Login(c *gin.Context) {
	// (処理内容は同じ)
	var req struct { Username string `json:"username"`; Password string `json:"password"` }
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"} ); return }
	storedPassword, ok := users[req.Username]
	if !ok { c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"} ); return }
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.Password)); err != nil { c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"} ); return }
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{ Username: req.Username, RegisteredClaims: jwt.RegisteredClaims{ ExpiresAt: jwt.NewNumericDate(expirationTime), }, }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"} ); return }
	c.JSON(http.StatusOK, gin.H{"token": tokenString} )
}

// Logout はログアウト処理を行います
func (h *AuthHandler) Logout(c *gin.Context) {
	// Authorizationヘッダーからトークンを取得
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		return
	}

	// "Bearer "プレフィックスを確認
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header format must be Bearer {token}"})
		return
	}

	tokenString := parts[1]

	// トークンを検証
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// トークンをブラックリストに追加
	blacklistedTokens[tokenString] = true

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// Profile はプロフィール表示処理を行います
func (h *AuthHandler) Profile(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{ "message":  "This is a protected route", "username": username, } )
}
