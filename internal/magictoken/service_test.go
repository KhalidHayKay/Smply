package magictoken

import (
	"context"
	"errors"
	"smply/utils"
	"testing"
)

type mockRepository struct {
	createFn      func(ctx context.Context, email, tokenHash string) error
	getValidFn    func(ctx context.Context, tokenHash string) (*MagicToken, error)
	markUsedFn    func(ctx context.Context, tokenHash string) error
	markAllUsedFn func(ctx context.Context, email string) error

	calls []string
}

func (m *mockRepository) Create(ctx context.Context, email, tokenHash string) error {
	m.calls = append(m.calls, "Create")
	return m.createFn(ctx, email, tokenHash)
}

func (m *mockRepository) GetValid(ctx context.Context, tokenHash string) (*MagicToken, error) {
	m.calls = append(m.calls, "GetValid")
	return m.getValidFn(ctx, tokenHash)
}

func (m *mockRepository) MarkUsed(ctx context.Context, tokenHash string) error {
	m.calls = append(m.calls, "MarkUsed")
	return m.markUsedFn(ctx, tokenHash)
}

func (m *mockRepository) MarkAllUsed(ctx context.Context, email string) error {
	m.calls = append(m.calls, "MarkAllUsed")
	return m.markAllUsedFn(ctx, email)
}

func TestCreateStoresHashedToken(t *testing.T) {
	ctx := context.Background()
	var capturedHash string
	repo := &mockRepository{
		createFn: func(ctx context.Context, email, tokenHash string) error {
			capturedHash = tokenHash
			return nil
		},
	}

	svc := NewService(repo)
	token, err := svc.Create(ctx, "alice@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if token == "" {
		t.Fatal("expected token to be generated, got empty string")
	}

	if capturedHash != utils.Hash(token) {
		t.Fatalf("expected token hash %q, got %q", utils.Hash(token), capturedHash)
	}

	if len(repo.calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(repo.calls))
	}

	if repo.calls[0] != "Create" {
		t.Fatalf("expected Create call, got %q", repo.calls[0])
	}
}

func TestCreateReturnsErrorWhenRepoFails(t *testing.T) {
	ctx := context.Background()
	repo := &mockRepository{
		createFn: func(ctx context.Context, email, tokenHash string) error {
			return errors.New("create failed")
		},
	}

	svc := NewService(repo)
	_, err := svc.Create(ctx, "alice@example.com")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestValidateMarksTokenUsedAfterSuccess(t *testing.T) {
	ctx := context.Background()
	gotHash := ""
	repo := &mockRepository{
		getValidFn: func(ctx context.Context, tokenHash string) (*MagicToken, error) {
			return &MagicToken{TokenHash: tokenHash, Email: "alice@example.com"}, nil
		},
		markUsedFn: func(ctx context.Context, tokenHash string) error {
			gotHash = tokenHash
			return nil
		},
	}

	svc := NewService(repo)
	magicToken, err := svc.Validate(ctx, "token")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if magicToken == nil {
		t.Fatal("expected magic token, got nil")
	}

	if gotHash != utils.Hash("token") {
		t.Fatalf("expected markUsed to be called with %q, got %q", utils.Hash("token"), gotHash)
	}

	if len(repo.calls) != 2 {
		t.Fatalf("expected 2 call, got %d", len(repo.calls))
	}

	if repo.calls[0] != "GetValid" || repo.calls[1] != "MarkUsed" {
		t.Fatalf("expected GetValid then MarkUsed, got %v", repo.calls)
	}
}

func TestValidateReturnsErrorWhenGetValidFails(t *testing.T) {
	ctx := context.Background()
	repo := &mockRepository{
		getValidFn: func(ctx context.Context, tokenHash string) (*MagicToken, error) {
			return nil, errors.New("invalid token")
		},
		markUsedFn: func(ctx context.Context, tokenHash string) error {
			t.Fatal("MarkUsed should not be called when GetValid fails")
			return nil
		},
	}

	svc := NewService(repo)
	_, err := svc.Validate(ctx, "token")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestValidateReturnsErrorWhenMarkUsedFails(t *testing.T) {
	ctx := context.Background()
	repo := &mockRepository{
		getValidFn: func(ctx context.Context, tokenHash string) (*MagicToken, error) {
			return &MagicToken{TokenHash: tokenHash}, nil
		},
		markUsedFn: func(ctx context.Context, tokenHash string) error {
			return errors.New("mark used failed")
		},
	}

	svc := NewService(repo)
	_, err := svc.Validate(ctx, "token")
	if err == nil {
		t.Fatal("expected error")
	}
}
