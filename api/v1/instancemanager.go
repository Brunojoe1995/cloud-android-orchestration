package v1

type CreateHostRequest struct {
	// [REQUIRED]
	CreateHostInstanceRequest *CreateHostInstanceRequest `json:"create_host_instance_request"`
}

type CreateHostInstanceRequest struct {
	// Required if using GCP.
	GCP *GCPInstance `json:"gcp"`
}

type GCPInstance struct {
	// [REQUIRED]
	DiskSizeGB int64 `json:"disk_size_gb"`
	// [REQUIRED] More info about this field in https://cloud.google.com/compute/docs/reference/rest/v1/instances/insert#request-body
	MachineType string `json:"machine_type"`
	// More info about this field in https://cloud.google.com/compute/docs/reference/rest/v1/instances/insert#request-body
	MinCPUPlatform string `json:"min_cpu_platform"`
}

type Operation struct {
	Name string `json:"name"`
	// Service-specific metadata associated with the operation.  It typically
	// contains progress information and common metadata such as create time.
	Metadata interface{} `json:"metadata,omitempty"`
	// If the value is `false`, it means the operation is still in progress.
	// If `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `json:"done"`
	// Result will contain either an error or a result object but never both.
	Result *OperationResult `json:"result,omitempty"`
}

type OperationResult struct {
	// The error result of the operation in case of failure or cancellation.
	Error *Error `json:"error,omitempty"`
	// The expected response of the operation in case of success.  If the original
	// method returns no data on success, such as `Delete`, this field will be
	// empty, hence omitted. If the original method is standard:
	// `Get`/`Create`/`Update`, the response should be the relevant resource.
	Response interface{} `json:"response,omitempty"`
}

type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type HostInstance struct {
	Name string       `json:"name"`
	GCP  *GCPInstance `json:"gcp,omitempty"`
}

type ListHostsResponse struct {
	Items []*HostInstance `json:"items"`
	// This token allows you to get the next page of results for list requests.
	// If the number of results is larger than maxResults, use the `nextPageToken`
	// as a value for the query parameter `pageToken` in the next list request.
	// Subsequent list requests will have their own `nextPageToken` to continue
	// paging through out all the results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}