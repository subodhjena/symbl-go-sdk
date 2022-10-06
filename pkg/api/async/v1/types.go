// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package async

/*
	Input parameters for API calls
*/
// WaitForJobStatusOpts parameter needed for Wait call
type WaitForJobStatusOpts struct {
	JobId         string `validate:"required"`
	WaitInSeconds int
}

/*
	Output structs for API calls
*/
// JobStatus captures the API for getting status
type JobStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// JobConversation represents processing an Async API request
type JobConversation struct {
	JobID          string `json:"jobId"`
	ConversationID string `json:"conversationId"`
}
