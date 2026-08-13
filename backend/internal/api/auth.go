package api

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/auth"
	"github.com/skriptra/skriptra/backend/internal/db"
	"github.com/skriptra/skriptra/backend/internal/domain"
)

// Authentication.
//
// The authorization checks were wired from the first endpoint and only the
// identity behind them was a stub. This is the identity.

const (
	// contextUserKey is where the middleware leaves the caller for handlers.
	contextUserKey = "skriptra.userID"

	// refreshCookie carries the refresh token.
	//
	// A cookie rather than a JSON field, and HttpOnly, so script running on the
	// page cannot read it. That is the whole point: an XSS bug can then steal
	// at most a 15-minute access token instead of a 30-day session. SameSite
	// Lax stops it riding along on a cross-site form post.
	refreshCookie = "skriptra_refresh"

	minPasswordLen = 10
)

type signupRequest struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	DisplayName string `json:"displayName"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type tokenResponse struct {
	AccessToken string      `json:"accessToken"`
	TokenType   string      `json:"tokenType"`
	ExpiresIn   int         `json:"expiresIn"`
	User        domain.User `json:"user"`
}

func (s *Server) signup(c *gin.Context) {
	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", err.Error())
		return
	}

	email := strings.TrimSpace(strings.ToLower(req.Email))
	if !looksLikeEmail(email) {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", "That does not look like an email address.")
		return
	}

	// A length floor and nothing else.
	//
	// Composition rules ("one uppercase, one symbol") push people towards
	// Password1! and are no longer recommended by NIST or OWASP. Length is the
	// property that actually costs an attacker work.
	if len([]rune(req.Password)) < minPasswordLen {
		fail(c, http.StatusUnprocessableEntity, "validation_failed",
			"Password must be at least 10 characters.")
		return
	}
	// Argon2 has no bcrypt-style truncation, but an unbounded password is an
	// unbounded amount of hashing work for anyone who posts one.
	if len(req.Password) > 1024 {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", "Password is too long.")
		return
	}

	displayName := strings.TrimSpace(req.DisplayName)
	if displayName == "" {
		displayName = email[:strings.Index(email, "@")]
	}

	hash, err := auth.HashPassword(req.Password)
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	account, err := s.store.CreateAccount(c, email, displayName, hash)
	if err != nil {
		if errors.Is(err, db.ErrEmailTaken) {
			// Registration cannot avoid confirming that an address is taken:
			// the alternative is accepting a signup that silently does
			// nothing. Login is where the distinction is hidden.
			fail(c, http.StatusConflict, "email_taken", "An account with that email already exists.")
			return
		}
		s.respond(c, err, nil)
		return
	}

	s.issueSession(c, account)
}

func (s *Server) login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusUnprocessableEntity, "validation_failed", err.Error())
		return
	}

	email := strings.TrimSpace(strings.ToLower(req.Email))
	account, err := s.store.AccountByEmail(c, email)

	switch {
	case errors.Is(err, db.ErrNotFound), account != nil && account.PasswordHash == "":
		// Hash anyway. Returning immediately here makes a login attempt for an
		// unregistered address measurably faster than one for a real account,
		// which turns this endpoint into an oracle for which addresses exist.
		auth.EqualiseTiming(req.Password)
		failInvalidCredentials(c)
		return
	case err != nil:
		s.respond(c, err, nil)
		return
	}

	ok, err := auth.VerifyPassword(req.Password, account.PasswordHash)
	if err != nil {
		s.log.Error("verify password", "user", account.ID, "error", err)
		failInvalidCredentials(c)
		return
	}
	if !ok {
		failInvalidCredentials(c)
		return
	}

	s.issueSession(c, account)
}

// refresh exchanges a refresh token for a new access token, rotating the
// refresh token in the process.
func (s *Server) refresh(c *gin.Context) {
	presented, err := c.Cookie(refreshCookie)
	if err != nil || presented == "" {
		fail(c, http.StatusUnauthorized, "no_session", "No session to refresh.")
		return
	}

	newToken, newHash, err := auth.NewRefreshToken()
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	expires := time.Now().Add(auth.RefreshTokenTTL)

	userID, err := s.store.RotateRefreshToken(
		c, auth.HashRefreshToken(presented), newHash, expires, c.Request.UserAgent())
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			// Expired, revoked and never-existed are one answer. Telling a
			// caller that a token was real but already used tells an attacker
			// holding a stolen copy that the theft was worth something.
			s.clearRefreshCookie(c)
			fail(c, http.StatusUnauthorized, "invalid_session", "Session is no longer valid. Sign in again.")
			return
		}
		s.respond(c, err, nil)
		return
	}

	account, err := s.store.AccountByID(c, userID)
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	access, accessExpires, err := s.issuer.Issue(userID)
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	s.setRefreshCookie(c, newToken, expires)
	c.JSON(http.StatusOK, tokenResponse{
		AccessToken: access,
		TokenType:   "Bearer",
		ExpiresIn:   int(time.Until(accessExpires).Seconds()),
		User:        domain.User{ID: account.ID, DisplayName: account.DisplayName, Email: account.Email},
	})
}

// logout revokes the presented session.
//
// Always answers 204. Whether the token was valid is not the caller's
// business, and a client that has decided to log out should end up logged out
// either way.
func (s *Server) logout(c *gin.Context) {
	if presented, err := c.Cookie(refreshCookie); err == nil && presented != "" {
		if err := s.store.RevokeRefreshToken(c, auth.HashRefreshToken(presented)); err != nil {
			s.log.Error("revoke refresh token", "error", err)
		}
	}
	s.clearRefreshCookie(c)
	c.Status(http.StatusNoContent)
}

// issueSession mints both tokens and answers with the access token in the body
// and the refresh token in an HttpOnly cookie.
func (s *Server) issueSession(c *gin.Context, account *db.Account) {
	access, accessExpires, err := s.issuer.Issue(account.ID)
	if err != nil {
		s.respond(c, err, nil)
		return
	}

	refreshToken, refreshHash, err := auth.NewRefreshToken()
	if err != nil {
		s.respond(c, err, nil)
		return
	}
	refreshExpires := time.Now().Add(auth.RefreshTokenTTL)

	if err := s.store.StoreRefreshToken(
		c, account.ID, refreshHash, refreshExpires, c.Request.UserAgent()); err != nil {
		s.respond(c, err, nil)
		return
	}

	s.setRefreshCookie(c, refreshToken, refreshExpires)
	c.JSON(http.StatusOK, tokenResponse{
		AccessToken: access,
		TokenType:   "Bearer",
		ExpiresIn:   int(time.Until(accessExpires).Seconds()),
		User:        domain.User{ID: account.ID, DisplayName: account.DisplayName, Email: account.Email},
	})
}

func (s *Server) setRefreshCookie(c *gin.Context, token string, expires time.Time) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     refreshCookie,
		Value:    token,
		Path:     "/api/v1/auth",
		Expires:  expires,
		MaxAge:   int(time.Until(expires).Seconds()),
		HttpOnly: true,
		// Secure is off over plain HTTP or the cookie is silently dropped and
		// local development breaks in a way that looks like a server bug. It is
		// on wherever the deployment is served over TLS.
		Secure:   s.cfg.CookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
}

func (s *Server) clearRefreshCookie(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     refreshCookie,
		Value:    "",
		Path:     "/api/v1/auth",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   s.cfg.CookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
}

// requireAuth rejects any request without a valid access token.
//
// Applied to the whole v1 group rather than listed per route. An allow-list
// that has to be extended for every new endpoint eventually misses one, and a
// missed one is an open endpoint nobody notices.
func (s *Server) requireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token, ok := strings.CutPrefix(header, "Bearer ")
		if !ok || strings.TrimSpace(token) == "" {
			fail(c, http.StatusUnauthorized, "unauthorized", "Sign in to continue.")
			return
		}

		userID, err := s.issuer.Verify(strings.TrimSpace(token))
		if err != nil {
			// The expiry case is separated because it is the one a client can
			// act on: refresh and retry. Anything else means sign in again.
			if errors.Is(err, auth.ErrExpiredToken) {
				fail(c, http.StatusUnauthorized, "token_expired", "Session expired.")
				return
			}
			fail(c, http.StatusUnauthorized, "unauthorized", "Sign in to continue.")
			return
		}

		c.Set(contextUserKey, userID)
		c.Next()
	}
}

func failInvalidCredentials(c *gin.Context) {
	// One message for a wrong password and for an address that was never
	// registered. Two messages would let anyone enumerate the user base.
	fail(c, http.StatusUnauthorized, "invalid_credentials", "Email or password is incorrect.")
}

// looksLikeEmail is a shape check, not a validity check.
//
// Whether an address can receive mail is only ever answered by sending to it.
// A regex that tries to implement RFC 5322 rejects real addresses, so this
// checks the two things a typo actually breaks.
func looksLikeEmail(s string) bool {
	at := strings.Index(s, "@")
	if at <= 0 || at != strings.LastIndex(s, "@") || len(s) > 254 {
		return false
	}
	domain := s[at+1:]
	return strings.Contains(domain, ".") &&
		!strings.HasPrefix(domain, ".") && !strings.HasSuffix(domain, ".") &&
		!strings.ContainsAny(s, " \t\r\n")
}

// currentUser resolves the caller from the access token the middleware
// verified. The single seam through which identity enters the handlers.
func currentUser(c *gin.Context) uuid.UUID {
	if v, ok := c.Get(contextUserKey); ok {
		if id, ok := v.(uuid.UUID); ok {
			return id
		}
	}
	// Unreachable behind requireAuth. Returning nil rather than panicking means
	// a route accidentally mounted outside the guarded group fails closed: the
	// nil uuid owns no rows, so every user-scoped query returns nothing.
	return uuid.Nil
}
