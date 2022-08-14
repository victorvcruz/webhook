package response

func GetType(data map[string]interface{}) PullRequestType {
	switch data["action"].(string) {
	case "closed":
		if data["pull_request"].(map[string]interface{})["merged"].(bool) == true {
			return MERGED_PULL_REQUEST
		}
		return CLOSED_PULL_REQUEST
	case "reopened":
		return REOPEN_PULL_REQUEST
	case "opened":
		return OPEN_PULL_REQUEST

	}
	return 0
}
