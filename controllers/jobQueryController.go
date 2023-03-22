package controllers

import (
	"encoding/json"
	"net/url"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/util"
)

const JobListURL string = "https://www.phoenix.global/sdk/computation/panel/jobListByUser"
const JobDetailURL string = "https://www.phoenix.global/sdk/computation/panel/jobDetailByUser"
const NoticeListURL string = "https://www.phoenix.global/sdk/computation/panel/noticeList"
const ResultListURL string = "https://www.phoenix.global/sdk/computation/panel/subListByUser"
const CcdUsageURL string = "https://www.phoenix.global/sdk/computation/panel/ccdUsageList"
const DownloadDatasetURL string = "https://www.phoenix.global/sdk/computation/panel/downloadDataset"

func JobListByUser(reqBody common.ReqPage, token string) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(JobListURL, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func JobDetailByUser(reqBody common.ReqJobDetailByUser, token string) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(JobDetailURL, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NoticeList(reqBody common.ReqPage, token string) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(NoticeListURL, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SubmissionList(reqBody common.ReqSubListByUser, token string) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(ResultListURL, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CcdUsageList(reqBody common.ReqPage, token string) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(CcdUsageURL, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DownloadDataset(reqParam common.ReqDownloadDataset, token string) (json.RawMessage, error) {
	params := url.Values{}
	params.Add("datasetType", reqParam.DatasetType)
	params.Add("jobID", reqParam.JobID)
	result, err := util.SendHttpGet(DownloadDatasetURL, params, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}
