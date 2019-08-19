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

// DeleteCdnDomain invokes the cdn.DeleteCdnDomain API synchronously
// api document: https://help.aliyun.com/api/cdn/deletecdndomain.html
func (client *Client) DeleteCdnDomain(request *DeleteCdnDomainRequest) (response *DeleteCdnDomainResponse, err error) {
	response = CreateDeleteCdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCdnDomainWithChan invokes the cdn.DeleteCdnDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/deletecdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCdnDomainWithChan(request *DeleteCdnDomainRequest) (<-chan *DeleteCdnDomainResponse, <-chan error) {
	responseChan := make(chan *DeleteCdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCdnDomain(request)
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

// DeleteCdnDomainWithCallback invokes the cdn.DeleteCdnDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/deletecdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCdnDomainWithCallback(request *DeleteCdnDomainRequest, callback func(response *DeleteCdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCdnDomainResponse
		var err error
		defer close(result)
		response, err = client.DeleteCdnDomain(request)
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

// DeleteCdnDomainRequest is the request struct for api DeleteCdnDomain
type DeleteCdnDomainRequest struct {
	*requests.RpcRequest
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	OwnerAccount    string           `position:"Query" name:"OwnerAccount"`
	DomainName      string           `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteCdnDomainResponse is the response struct for api DeleteCdnDomain
type DeleteCdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCdnDomainRequest creates a request to invoke DeleteCdnDomain API
func CreateDeleteCdnDomainRequest() (request *DeleteCdnDomainRequest) {
	request = &DeleteCdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DeleteCdnDomain", "cdn", "openAPI")
	return
}

// CreateDeleteCdnDomainResponse creates a response to parse from DeleteCdnDomain response
func CreateDeleteCdnDomainResponse() (response *DeleteCdnDomainResponse) {
	response = &DeleteCdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
