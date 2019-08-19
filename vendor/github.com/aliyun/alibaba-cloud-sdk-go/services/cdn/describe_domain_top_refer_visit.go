package cdn

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

// DescribeDomainTopReferVisit invokes the cdn.DescribeDomainTopReferVisit API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomaintoprefervisit.html
func (client *Client) DescribeDomainTopReferVisit(request *DescribeDomainTopReferVisitRequest) (response *DescribeDomainTopReferVisitResponse, err error) {
	response = CreateDescribeDomainTopReferVisitResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainTopReferVisitWithChan invokes the cdn.DescribeDomainTopReferVisit API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomaintoprefervisit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainTopReferVisitWithChan(request *DescribeDomainTopReferVisitRequest) (<-chan *DescribeDomainTopReferVisitResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainTopReferVisitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainTopReferVisit(request)
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

// DescribeDomainTopReferVisitWithCallback invokes the cdn.DescribeDomainTopReferVisit API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomaintoprefervisit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainTopReferVisitWithCallback(request *DescribeDomainTopReferVisitRequest, callback func(response *DescribeDomainTopReferVisitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainTopReferVisitResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainTopReferVisit(request)
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

// DescribeDomainTopReferVisitRequest is the request struct for api DescribeDomainTopReferVisit
type DescribeDomainTopReferVisitRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	Percent    string           `position:"Query" name:"Percent"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	SortBy     string           `position:"Query" name:"SortBy"`
}

// DescribeDomainTopReferVisitResponse is the response struct for api DescribeDomainTopReferVisit
type DescribeDomainTopReferVisitResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	DomainName   string       `json:"DomainName" xml:"DomainName"`
	StartTime    string       `json:"StartTime" xml:"StartTime"`
	TopReferList TopReferList `json:"TopReferList" xml:"TopReferList"`
}

// CreateDescribeDomainTopReferVisitRequest creates a request to invoke DescribeDomainTopReferVisit API
func CreateDescribeDomainTopReferVisitRequest() (request *DescribeDomainTopReferVisitRequest) {
	request = &DescribeDomainTopReferVisitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainTopReferVisit", "cdn", "openAPI")
	return
}

// CreateDescribeDomainTopReferVisitResponse creates a response to parse from DescribeDomainTopReferVisit response
func CreateDescribeDomainTopReferVisitResponse() (response *DescribeDomainTopReferVisitResponse) {
	response = &DescribeDomainTopReferVisitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
