package entities

import (
	"errors"
	"time"

	"encore.app/apicore/settings/domain/valueobjects"
	"github.com/google/uuid"
)

// Setting represents a setting flag in the domain layer.
type Setting struct {
	ID             uuid.UUID
	Name           string
	Slug           string
	Hint           string
	IsActive       bool
	CreatedAt      time.Time
	CreatedBy      uuid.UUID
	DeletedAt      *time.Time
	DeletedBy      *uuid.UUID
	UpdatedAt      *time.Time
	UpdatedBy      *uuid.UUID
	Version        int32
	tags           []Tag
	targetingRules []valueobjects.TargetingRuleVO
}

// NewSetting creates a new setting instance without tags and targeting rules.
func NewSetting(id uuid.UUID, name, slug, hint string, isActive bool, createdAt time.Time, createdBy uuid.UUID, DeletedAt *time.Time, DeletedBy *uuid.UUID, updatedAt *time.Time, updatedBy *uuid.UUID, version int32) (Setting, error) {
	if name == "" {
		return Setting{}, errors.New("setting name cannot be empty")
	}

	if slug == "" {
		return Setting{}, errors.New("setting slug cannot be empty")
	}

	domainEntity := Setting{
		ID:             id,
		Name:           name,
		Slug:           slug,
		Hint:           hint,
		IsActive:       isActive,
		CreatedAt:      createdAt,
		CreatedBy:      createdBy,
		DeletedAt:      updatedAt,
		DeletedBy:      updatedBy,
		UpdatedAt:      updatedAt,
		UpdatedBy:      updatedBy,
		Version:        version,
		tags:           []Tag{},
		targetingRules: []valueobjects.TargetingRuleVO{},
	}

	return domainEntity, nil
}

// GetTags returns a copy of the setting's tags to prevent direct slice manipulation
func (f *Setting) GetTags() []Tag {
	tagsCopy := make([]Tag, len(f.tags))
	copy(tagsCopy, f.tags)
	return tagsCopy
}

// SetTags replaces all existing tags with the provided tags
func (f *Setting) SetTags(tags []Tag) {
	f.tags = make([]Tag, len(tags))
	copy(f.tags, tags)
}

// AddTag adds a single tag to the setting
func (f *Setting) AddTag(tag Tag) {
	f.tags = append(f.tags, tag)
}

// RemoveTag removes a tag from the setting by its ID
func (f *Setting) RemoveTag(tagID string) {
	for i, tag := range f.tags {
		if tag.ID == tagID {
			f.tags = append(f.tags[:i], f.tags[i+1:]...)
			return
		}
	}
}

// GetTargetingRules returns a copy of the setting's targeting rules
func (f *Setting) GetTargetingRules() []valueobjects.TargetingRuleVO {
	rulesCopy := make([]valueobjects.TargetingRuleVO, len(f.targetingRules))
	copy(rulesCopy, f.targetingRules)
	return rulesCopy
}

// SetTargetingRules replaces all existing targeting rules
func (f *Setting) SetTargetingRules(rules []valueobjects.TargetingRuleVO) {
	f.targetingRules = make([]valueobjects.TargetingRuleVO, len(rules))
	copy(f.targetingRules, rules)
}

// WithTags adds tags to the setting.
func (f *Setting) WithTags(tags []Tag) *Setting {
	f.SetTags(tags)
	return f
}

// WithTargetingRules adds targeting rules to the setting.
func (f *Setting) WithTargetingRules(rules []valueobjects.TargetingRuleVO) *Setting {
	f.SetTargetingRules(rules)
	return f
}
