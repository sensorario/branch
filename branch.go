package branch

import (
	"strings"
)

type Branch struct {
	Name string
}

func (b Branch) destination() string {
	tokens := strings.Split(b.Name, "/")
	return tokens[len(tokens)-1]
}

func (b Branch) IsRefactoring() bool {
	return strings.HasPrefix(b.Name, "refactor/")
}

func (b Branch) IsPatch() bool {
	return strings.HasPrefix(b.Name, "patch/")
}

func (b Branch) IsFeature() bool {
	featureBranches := []string{"feature", "feat"}
	for _, v := range featureBranches {
		if strings.HasPrefix(b.Name, v) {
			return true
		}
	}
	return false
}

func (b Branch) IsHotfix() bool {
	return strings.HasPrefix(b.Name, "hotfix/")
}

func (b Branch) IsBugfix() bool {
	return strings.HasPrefix(b.Name, "bugfix/")
}

func (b Branch) IsDevelopment(DevBranchName string) bool {
	return strings.HasPrefix(b.Name, DevBranchName)
}

func (b Branch) IsRelease() bool {
	return strings.HasPrefix(b.Name, "release/")
}

func (b Branch) commitPrefix() string {
	if b.isBugfix() == true || b.isHotfix() == true {
		return "fix: "
	}
	return "feat: "
}
