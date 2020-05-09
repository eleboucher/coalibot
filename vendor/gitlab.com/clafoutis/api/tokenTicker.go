package api

import (
	"time"
)

func (t *Token) startTicker() {
	ticker := time.NewTicker(t.expiresIn)
	defer ticker.Stop()
	for {
		select {
		case <-t.c:
			return
		case <-ticker.C:
			t.c <- true
			return
		}
	}
}

func (t *Token) Expired() error {
	select {
	case <-t.c:
		if _, err := t.Reinit(); err == nil {
			return err
		}
	default:
	}
	return nil
}

func (t *Token) Kill() {
	select {
	case t.c <- true:
	default:
	}
	t.Client.Token = nil
}

func (t *Token) GetExpiresIn() time.Duration {
	return t.expiresIn
}

func (t *Token) GetCreatedAt() time.Time {
	return t.createdAt
}
