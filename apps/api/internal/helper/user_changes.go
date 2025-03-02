package helper

import "github.com/abgeo/follytics/internal/model"

// UserChanges represents the differences between two collections of users,
// tracking which users were added or removed based on GitHub IDs.
type UserChanges struct {
	additions []*model.User
	deletions []*model.User

	additionsCount int
	deletionsCount int
}

// NewUserChanges creates a UserChanges by comparing original and updated user lists.
// It identifies additions (users in updated but not in original) and
// deletions (users in original but not in updated).
func NewUserChanges(original, updated []*model.User) *UserChanges {
	uc := &UserChanges{}
	originalMap := make(map[int64]struct{})
	updatedMap := make(map[int64]struct{})

	for _, user := range original {
		originalMap[user.GHID] = struct{}{}
	}

	for _, user := range updated {
		updatedMap[user.GHID] = struct{}{}

		if _, exists := originalMap[user.GHID]; !exists {
			uc.additions = append(uc.additions, user)
		}
	}

	for _, user := range original {
		if _, exists := updatedMap[user.GHID]; !exists {
			uc.deletions = append(uc.deletions, user)
		}
	}

	uc.additionsCount = len(uc.additions)
	uc.deletionsCount = len(uc.deletions)

	return uc
}

// HasAdditions returns true if there are any added users.
func (uc *UserChanges) HasAdditions() bool {
	return uc.AdditionsCount() > 0
}

// AdditionsCount returns the number of added users.
func (uc *UserChanges) AdditionsCount() int {
	return uc.additionsCount
}

// HasDeletions returns true if there are any deleted users.
func (uc *UserChanges) HasDeletions() bool {
	return uc.DeletionsCount() > 0
}

// DeletionsCount returns the number of deleted users.
func (uc *UserChanges) DeletionsCount() int {
	return uc.deletionsCount
}

// Additions returns the slice of added users.
func (uc *UserChanges) Additions() []*model.User {
	return uc.additions
}

// Deletions returns the slice of deleted users.
func (uc *UserChanges) Deletions() []*model.User {
	return uc.deletions
}

// HasChanges returns true if there are any additions or deletions.
func (uc *UserChanges) HasChanges() bool {
	return uc.HasAdditions() || uc.HasDeletions()
}
