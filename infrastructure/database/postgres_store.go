package database

import (
	"fmt"
	"golang-course-registration/common/exception"

	supabase "github.com/supabase-community/supabase-go"
)

type SupabaseStore struct {
	Client *supabase.Client
}

func NewSupabase(url, key string) (*SupabaseStore, error) {
	if url == "" || key == "" {
		return nil, fmt.Errorf(exception.ErrDatabaseConfigInvalid)
	}

	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		return nil, fmt.Errorf(exception.ErrFailedCreateClient)
	}

	return &SupabaseStore{Client: client}, nil
}
