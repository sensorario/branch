package branch

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	br := Branch{"feature/Branch-semantico/master"}

	if br.IsFeature() == false {
		t.Errorf("Branch should be recognyzed as feature")
	}

	if br.destination() != "master" {
		t.Errorf("destination Branch should be master")
	}
}

func TestDetectIfBranchIsPatchOrNot(t *testing.T) {
	br := Branch{"patch/Branch-semantico/master"}
	if br.IsPatch() == false {
		t.Errorf("Branch should be recognyzed as patch")
	}
}

func TestDetectIfBranchIsNotAFeature(t *testing.T) {
	br := Branch{"patch/Branch-semantico/master"}
	if br.IsFeature() == true {
		t.Errorf("Branch should not be recognyzed as feature")
	}
}

func TestExtractCurrentBranch(t *testing.T) {
	br := Branch{"1.0"}
	if br.Name != "1.0" {
		t.Errorf("Oops! Branch detection fails!")
	}
}

func TestRefactoringBranchStartsWithRefactor(t *testing.T) {
	br := Branch{"refactor/description/develop"}
	if br.IsRefactoring() == false {
		t.Errorf("Branch should be rafactorng but is not")
	}
}

func TestHotfixBranchStartsWithHotfix(t *testing.T) {
	br := Branch{"hotfix/description/develop"}
	if br.IsHotfix() == false {
		t.Errorf("Branch should be hotfix but is not")
	}
}

func TestBugfixBranchStartsWithBugfix(t *testing.T) {
	br := Branch{"bugfix/description/develop"}
	if br.IsBugfix() == false {
		t.Errorf("Branch should be bugfix but is not")
	}
}

func TestDevelopmentBranchContains(t *testing.T) {
	br := Branch{"master"}
	if br.IsDevelopment("master") == false {
		t.Errorf("Branch should be development but is not")
	}
}

func TestReleaseBranchStartsWithRelease(t *testing.T) {
	br := Branch{"release/foo/bar"}
	if br.IsRelease() == false {
		t.Errorf("Branch should be release but is not")
	}
}

func TestExtractCommitPrefix(t *testing.T) {
	br := Branch{"feature/foo/bar"}
	if br.commitPrefix() != "feat: " {
		t.Errorf("commit prefix should start with 'feat: '")
	}
}

func TestPrefixMustBeDifferentForFixes(t *testing.T) {
	br := Branch{"bugfix/foo/bar"}
	if br.commitPrefix() != "fix: " {
		t.Errorf("commit prefix should start with 'fix: '")
	}
}

func TestPrefixMustBeDifferentForHotFixes(t *testing.T) {
	br := Branch{"hotfix/foo/bar"}
	if br.commitPrefix() != "fix: " {
		t.Errorf("commit prefix should start with 'fix: '")
	}
}
