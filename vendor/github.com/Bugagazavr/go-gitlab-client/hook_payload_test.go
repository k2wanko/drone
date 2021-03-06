package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestParsePushHook(t *testing.T) {
	stub, _ := ioutil.ReadFile("stubs/hook_payloads/push.json")
	p, err := ParseHook([]byte(stub))

	assert.Equal(t, err, nil)
	assert.Equal(t, p.ObjectKind, "push")
	assert.IsType(t, new(HookPayload), p)
	assert.Equal(t, p.After, "da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	assert.Equal(t, p.Repository.URL, "git@example.com:mike/diasporadiaspora.git")
	assert.Equal(t, p.Repository.GitHttpUrl, "http://example.com/mike/diaspora.git")
	assert.Equal(t, p.Repository.GitSshUrl, "git@example.com:mike/diaspora.git")
	assert.Equal(t, p.Repository.VisibilityLevel, 0)
	assert.Equal(t, len(p.Commits), 2)
	assert.Equal(t, p.Commits[0].Author.Email, "jordi@softcatala.org")
	assert.Equal(t, p.Commits[1].Id, "da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	assert.Equal(t, p.Branch(), "master")
	assert.Equal(t, p.Head().Message, "fixed readme")
}

func TestParseIssueHook(t *testing.T) {
	stub, _ := ioutil.ReadFile("stubs/hook_payloads/issue.json")
	p, err := ParseHook([]byte(stub))

	assert.Equal(t, err, nil)
	assert.Equal(t, p.ObjectKind, "issue")
	assert.Equal(t, p.ObjectAttributes.Id, 301)
}

func TestParseMergeRequestHook(t *testing.T) {
	stub, _ := ioutil.ReadFile("stubs/hook_payloads/merge_request.json")
	p, err := ParseHook([]byte(stub))

	assert.Equal(t, err, nil)
	assert.Equal(t, p.ObjectKind, "merge_request")
	assert.Equal(t, p.ObjectAttributes.TargetBranch, "master")
	assert.Equal(t, p.ObjectAttributes.SourceProjectId, p.ObjectAttributes.TargetProjectId)
}
