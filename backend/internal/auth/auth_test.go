package auth

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestHashPasswordRoundTrip(t *testing.T) {
	hash, err := HashPassword("correct horse battery staple")
	if err != nil {
		t.Fatal(err)
	}

	ok, err := VerifyPassword("correct horse battery staple", hash)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("the correct password did not verify")
	}

	ok, err = VerifyPassword("correct horse battery stapl", hash)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("a wrong password verified")
	}
}

// Two accounts with the same password must not produce the same hash, or a
// leaked table shows at a glance who shares a password with whom.
func TestHashPasswordIsSalted(t *testing.T) {
	a, err := HashPassword("same password")
	if err != nil {
		t.Fatal(err)
	}
	b, err := HashPassword("same password")
	if err != nil {
		t.Fatal(err)
	}
	if a == b {
		t.Fatal("identical passwords produced identical hashes, the salt is not random")
	}
}

// The parameters travel with the hash so they can be raised later without
// invalidating every stored password.
func TestHashPasswordIsSelfDescribing(t *testing.T) {
	hash, err := HashPassword("whatever")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(hash, "$argon2id$v=19$m=19456,t=2,p=1$") {
		t.Fatalf("unexpected hash format: %q", hash)
	}
}

func TestVerifyPasswordRejectsMalformedHashes(t *testing.T) {
	cases := map[string]string{
		"empty":             "",
		"not phc":           "plaintext",
		"wrong field count": "$argon2id$v=19$m=19456,t=2,p=1$salt",
		"unknown algorithm": "$bcrypt$v=19$m=19456,t=2,p=1$c2FsdA$ZGlnZXN0",
		"bad base64":        "$argon2id$v=19$m=19456,t=2,p=1$!!!!$!!!!",
	}
	for name, hash := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := VerifyPassword("x", hash); err == nil {
				t.Fatal("expected an error, got nil")
			}
		})
	}
}

func TestNewIssuerRejectsWeakSecrets(t *testing.T) {
	if _, err := NewIssuer("too short"); err == nil {
		t.Fatal("a 9-byte signing secret was accepted")
	}
	if _, err := NewIssuer(strings.Repeat("k", 32)); err != nil {
		t.Fatalf("a 32-byte secret was rejected: %v", err)
	}
}

func TestIssueAndVerify(t *testing.T) {
	issuer, err := NewIssuer(strings.Repeat("s", 32))
	if err != nil {
		t.Fatal(err)
	}
	want := uuid.New()

	token, expires, err := issuer.Issue(want)
	if err != nil {
		t.Fatal(err)
	}
	if time.Until(expires) > AccessTokenTTL+time.Second {
		t.Fatalf("expiry is further out than the TTL: %v", expires)
	}

	got, err := issuer.Verify(token)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("got subject %s, want %s", got, want)
	}
}

// A token signed with a different key must not verify, which is the entire
// point of signing it.
func TestVerifyRejectsForeignSignature(t *testing.T) {
	mine, _ := NewIssuer(strings.Repeat("a", 32))
	theirs, _ := NewIssuer(strings.Repeat("b", 32))

	token, _, err := theirs.Issue(uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := mine.Verify(token); err == nil {
		t.Fatal("a token signed with another key verified")
	}
}

// Algorithm confusion: a token whose header says the algorithm is "none" must
// be rejected outright, not treated as unsigned-but-fine.
func TestVerifyRejectsAlgNone(t *testing.T) {
	issuer, _ := NewIssuer(strings.Repeat("s", 32))

	// {"alg":"none","typ":"JWT"} . {"sub":...} . (empty signature)
	forged := "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0." +
		"eyJzdWIiOiIxMTExMTExMS0xMTExLTExMTEtMTExMS0xMTExMTExMTExMTEiLCJpc3MiOiJza3JpcHRyYSJ9."

	if _, err := issuer.Verify(forged); err == nil {
		t.Fatal("an alg=none token was accepted")
	}
}

func TestVerifyRejectsExpiredToken(t *testing.T) {
	issuer, _ := NewIssuer(strings.Repeat("s", 32))

	// Issue in the past, then verify at real "now".
	past := time.Now().Add(-2 * AccessTokenTTL)
	issuer.now = func() time.Time { return past }
	token, _, err := issuer.Issue(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	issuer.now = time.Now
	_, err = issuer.Verify(token)
	if err != ErrExpiredToken {
		t.Fatalf("got %v, want ErrExpiredToken", err)
	}
}

func TestVerifyRejectsGarbage(t *testing.T) {
	issuer, _ := NewIssuer(strings.Repeat("s", 32))
	for _, s := range []string{"", "not.a.token", "a.b.c", strings.Repeat("x", 500)} {
		if _, err := issuer.Verify(s); err == nil {
			t.Fatalf("garbage token %q verified", s)
		}
	}
}

// The stored value must not be the token itself, or a database dump hands an
// attacker a working session.
func TestRefreshTokenIsStoredHashed(t *testing.T) {
	token, hash, err := NewRefreshToken()
	if err != nil {
		t.Fatal(err)
	}
	if token == "" || hash == "" {
		t.Fatal("empty token or hash")
	}
	if token == hash {
		t.Fatal("the stored hash is the token itself")
	}
	if HashRefreshToken(token) != hash {
		t.Fatal("hashing the token does not reproduce the stored hash")
	}

	other, _, err := NewRefreshToken()
	if err != nil {
		t.Fatal(err)
	}
	if other == token {
		t.Fatal("two refresh tokens came out identical")
	}
}
