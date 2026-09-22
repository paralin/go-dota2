package dota2

import "errors"

// ErrNotReady indicates that the GC session is unavailable or was interrupted.
var ErrNotReady = errors.New("the Dota2 game coordinator session is not ready")
