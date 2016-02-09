package angelco

import (
   "fmt"
    "net/url"
)

func (api *AngelApi) JobsList() (res *JobsListResponse, err error) {
    res = new(JobsListResponse)
    err = api.get("/jobs/", url.Values{}, res)
    return
}

func (api *AngelApi) Job(jobId int64) (res *JobResponse, err error) {
    res = new(JobResponse)
    err = api.get(fmt.Sprintf("/jobs/%d", jobId), url.Values{}, res)
    return
}

func (api *AngelApi) JobsOfStartup(startupId int64) (res []Job,  err error) {
    res = []Job{} 
    err = api.get(fmt.Sprintf("/startups/%d/jobs/", startupId), url.Values{}, &res)
    return 
}

func (api *AngelApi) JobsOfTag(tagId int64) (res *JobsListResponse, err error) {
    res = new(JobsListResponse)
    err = api.get(fmt.Sprintf("/tags/%d/jobs/", tagId), url.Values{}, res)
    return 
}
