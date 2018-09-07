package dds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeAvaliableTimeRange invokes the dds.DescribeAvaliableTimeRange API synchronously
// api document: https://help.aliyun.com/api/dds/describeavaliabletimerange.html
func (client *Client) DescribeAvaliableTimeRange(request *DescribeAvaliableTimeRangeRequest) (response *DescribeAvaliableTimeRangeResponse, err error) {
	response = CreateDescribeAvaliableTimeRangeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAvaliableTimeRangeWithChan invokes the dds.DescribeAvaliableTimeRange API asynchronously
// api document: https://help.aliyun.com/api/dds/describeavaliabletimerange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvaliableTimeRangeWithChan(request *DescribeAvaliableTimeRangeRequest) (<-chan *DescribeAvaliableTimeRangeResponse, <-chan error) {
	responseChan := make(chan *DescribeAvaliableTimeRangeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAvaliableTimeRange(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeAvaliableTimeRangeWithCallback invokes the dds.DescribeAvaliableTimeRange API asynchronously
// api document: https://help.aliyun.com/api/dds/describeavaliabletimerange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvaliableTimeRangeWithCallback(request *DescribeAvaliableTimeRangeRequest, callback func(response *DescribeAvaliableTimeRangeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAvaliableTimeRangeResponse
		var err error
		defer close(result)
		response, err = client.DescribeAvaliableTimeRange(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeAvaliableTimeRangeRequest is the request struct for api DescribeAvaliableTimeRange
type DescribeAvaliableTimeRangeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NodeId               string           `position:"Query" name:"NodeId"`
}

// DescribeAvaliableTimeRangeResponse is the response struct for api DescribeAvaliableTimeRange
type DescribeAvaliableTimeRangeResponse struct {
	*responses.BaseResponse
	RequestId string                                `json:"RequestId" xml:"RequestId"`
	TimeRange TimeRangeInDescribeAvaliableTimeRange `json:"TimeRange" xml:"TimeRange"`
}

// CreateDescribeAvaliableTimeRangeRequest creates a request to invoke DescribeAvaliableTimeRange API
func CreateDescribeAvaliableTimeRangeRequest() (request *DescribeAvaliableTimeRangeRequest) {
	request = &DescribeAvaliableTimeRangeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeAvaliableTimeRange", "dds", "openAPI")
	return
}

// CreateDescribeAvaliableTimeRangeResponse creates a response to parse from DescribeAvaliableTimeRange response
func CreateDescribeAvaliableTimeRangeResponse() (response *DescribeAvaliableTimeRangeResponse) {
	response = &DescribeAvaliableTimeRangeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}