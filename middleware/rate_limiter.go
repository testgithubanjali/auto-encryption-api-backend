package middleware

import (
	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
)

func EncryptLimiter() *limiter.Limiter {
	lmt := tollbooth.NewLimiter(5, nil)
	lmt.SetBurst(10)
	return lmt
}

func DecryptLimiter() *limiter.Limiter {
	lmt := tollbooth.NewLimiter(3, nil)
	lmt.SetBurst(5)
	return lmt
}

func EncodeLimiter() *limiter.Limiter {
	return EncryptLimiter()
}

func DecodeLimiter() *limiter.Limiter {
	return DecryptLimiter()
}

func EncryptFileLimiter() *limiter.Limiter {
	lmt := tollbooth.NewLimiter(2, nil)
	lmt.SetBurst(3)
	return lmt
}

// File Decrypt
func DecryptFileLimiter() *limiter.Limiter {
	lmt := tollbooth.NewLimiter(1, nil)
	lmt.SetBurst(2)
	return lmt
}
