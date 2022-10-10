package branch

import (
	"strings"
)

type Branch struct {
	name string
}

func (b Branch) destination() string {
	tokens := strings.Split(b.name, "/")
	return tokens[len(tokens)-1]
}

func (b Branch) IsRefactoring() bool {
	return strings.HasPrefix(b.name, "refactor/")
}

func (b Branch) IsPatch() bool {
	return strings.HasPrefix(b.name, "patch/")
}

func (b Branch) IsFeature() bool {
	featureBranches := []string{"feature", "feat"}
	for _, v := range featureBranches {
		if strings.HasPrefix(b.name, v) {
			return true
		}
	}
	return false
}

func (b Branch) IsHotfix() bool {
	return strings.HasPrefix(b.name, "hotfix/")
}

func (b Branch) IsBugfix() bool {
	return strings.HasPrefix(b.name, "bugfix/")
}

func (b Branch) IsDevelopment(DevBranchName string) bool {
	return strings.HasPrefix(b.name, DevBranchName)
}

func (b Branch) IsRelease() bool {
	return strings.HasPrefix(b.name, "release/")
}

func (b Branch) commitPrefix() string {
	if b.isBugfix() == true || b.isHotfix() == true {
		return "fix: "
	}
	return "feat: "
}
