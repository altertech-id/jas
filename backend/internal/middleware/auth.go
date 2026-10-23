package middleware

import (
	"strings"

	"jayaabadisolusi/backend/internal/config"
	"jayaabadisolusi/backend/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// UserContextKey is key stored in c.Locals
const UserContextKey = "user_profile"

type TokenClaims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// AuthMiddleware verifies the JWT token from Supabase or allows dev bypass
func AuthMiddleware(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Header otorisasi diperlukan.",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Format token tidak valid. Gunakan format Bearer <token>.",
			})
		}

		tokenStr := parts[1]

		// Support local dev tokens for immediate testing without live Supabase
		if tokenStr == "dev-admin-token" {
			c.Locals(UserContextKey, &domain.Profile{
				ID:       "a0000000-0000-0000-0000-000000000001",
				Email:    "admin@jayaabadisolusi.com",
				Name:     "Administrator JAS",
				Role:     domain.RoleAdmin,
				IsActive: true,
			})
			return c.Next()
		} else if tokenStr == "dev-editor-token" {
			c.Locals(UserContextKey, &domain.Profile{
				ID:       "a0000000-0000-0000-0000-000000000002",
				Email:    "editor@jayaabadisolusi.com",
				Name:     "Tim Konten JAS",
				Role:     domain.RoleEditor,
				IsActive: true,
			})
			return c.Next()
		}

		// Parse and verify standard Supabase JWT
		token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(cfg.SupabaseJWTSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token tidak valid atau telah kedaluwarsa.",
			})
		}

		claims, ok := token.Claims.(*TokenClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Klaim token tidak terbaca.",
			})
		}

		userRole := domain.RoleEditor
		if claims.Role == string(domain.RoleAdmin) {
			userRole = domain.RoleAdmin
		}

		c.Locals(UserContextKey, &domain.Profile{
			ID:       claims.Sub,
			Email:    claims.Email,
			Role:     userRole,
			IsActive: true,
		})

		return c.Next()
	}
}

// RequireRole checks if authenticated user matches one of the required roles
func RequireRole(roles ...domain.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		profile, ok := c.Locals(UserContextKey).(*domain.Profile)
		if !ok || profile == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Pengguna belum diautentikasi.",
			})
		}

		for _, r := range roles {
			if profile.Role == r {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Akses ditolak. Anda tidak memiliki izin untuk fitur ini.",
		})
	}
}
