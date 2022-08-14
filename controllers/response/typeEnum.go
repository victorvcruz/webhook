package response

type PullRequestType int

const (
	OPEN_PULL_REQUEST PullRequestType = iota
	APPROVED_PULL_REQUEST
	MERGED_PULL_REQUEST
	CLOSED_PULL_REQUEST
	REOPEN_PULL_REQUEST
)
