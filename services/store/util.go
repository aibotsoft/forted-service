package store

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *Store) CheckInCache(ctx context.Context, key string) (int64, bool) {
	if get, b := s.cache.Get(key); b {
		if value, ok := get.(int64); ok {
			//s.log.Debugf("Got %q from cache ", key)
			return value, true
		}
	}
	return 0, false
}

func (s *Store) FormKey(name string, keys ...string) string {
	return name + ":" + strings.Join(keys[:], ":")
}

func (s *Store) LogAndReturnErr(err error, code codes.Code, msg string) error {
	s.log.Error(msg, err)
	return status.Errorf(code, "%v error: %v ", msg, err)
}
