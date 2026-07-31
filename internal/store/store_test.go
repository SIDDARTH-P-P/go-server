package store

import "testing"

func TestStoreCRUD(t *testing.T) {
	s := NewStore()

	created, err := s.Create("First item")
	if err != nil {
		t.Fatalf("Create() unexpected error: %v", err)
	}

	if created.ID != 1 {
		t.Fatalf("expected id 1, got %d", created.ID)
	}

	got, err := s.Get(1)
	if err != nil {
		t.Fatalf("Get() unexpected error: %v", err)
	}
	if got.Title != "First item" {
		t.Fatalf("expected title %q, got %q", "First item", got.Title)
	}

	updated, err := s.Update(1, "Updated item")
	if err != nil {
		t.Fatalf("Update() unexpected error: %v", err)
	}
	if updated.Title != "Updated item" {
		t.Fatalf("expected updated title %q, got %q", "Updated item", updated.Title)
	}

	if err := s.Delete(1); err != nil {
		t.Fatalf("Delete() unexpected error: %v", err)
	}

	if _, err := s.Get(1); err == nil {
		t.Fatal("expected Get() to fail after Delete()")
	}
}
