package main

type PullRequestAction int

const (
	START_PULL_REQUEST PullRequestAction = iota
	APROVED_PULL_REQUEST
	MERGED_PULL_REQUEST
	CLOSED_PULL_REQUEST
)
