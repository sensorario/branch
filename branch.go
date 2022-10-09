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

func (b Branch) isRefactoring() bool {
	return strings.HasPrefix(b.name, "refactor/")
}

func (b Branch) isPatch() bool {
	return strings.HasPrefix(b.name, "patch/")
}

func (b Branch) isFeature() bool {
	featureBranches := []string{"feature", "feat"}
	for _, v := range featureBranches {
		if strings.HasPrefix(b.name, v) {
			return true
		}
	}
	return false
}

func (b Branch) isHotfix() bool {
	return strings.HasPrefix(b.name, "hotfix/")
}

func (b Branch) isBugfix() bool {
	return strings.HasPrefix(b.name, "bugfix/")
}

func (b Branch) isDevelopment(DevBranchName string) bool {
	return strings.HasPrefix(b.name, DevBranchName)
}

func (b Branch) isRelease() bool {
	return strings.HasPrefix(b.name, "release/")
}

func (b Branch) commitPrefix() string {
	if b.isBugfix() == true || b.isHotfix() == true {
		return "fix: "
	}
	return "feat: "
}
