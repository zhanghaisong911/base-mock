package ft

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

func (client *Client) MultiInputParamApi(request *MultiInputParamApiRequest) (response *MultiInputParamApiResponse, err error) {
	response = CreateMultiInputParamApiResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) MultiInputParamApiWithChan(request *MultiInputParamApiRequest) (<-chan *MultiInputParamApiResponse, <-chan error) {
	responseChan := make(chan *MultiInputParamApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MultiInputParamApi(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) MultiInputParamApiWithCallback(request *MultiInputParamApiRequest, callback func(response *MultiInputParamApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MultiInputParamApiResponse
		var err error
		defer close(result)
		response, err = client.MultiInputParamApi(request)
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

type MultiInputParamApiRequest struct {
	*requests.RpcRequest
	RequiredValue  string                              `position:"Query" name:"RequiredValue"`
	JsonRepeatList *[]MultiInputParamApiJsonRepeatList `position:"Query" name:"JsonRepeatList"  type:"Repeated"`
	RequestId      string                              `position:"Query" name:"RequestId"`
	SwitchValue    string                              `position:"Query" name:"SwitchValue"`
}

type MultiInputParamApiJsonRepeatList struct {
	Size string `name:"Size"`
	Type string `name:"Type"`
}

type MultiInputParamApiResponse struct {
	*responses.BaseResponse
	L1RequestId    string `json:"l1RequestId"`
	L1Map          string `json:"l1Map"`
	L1FtTestParam  string `json:"l1FtTestParam"`
	L1FtTestParams string `json:"l1FtTestParams"`
}

func CreateMultiInputParamApiRequest() (request *MultiInputParamApiRequest) {
	request = &MultiInputParamApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ft", "2015-01-01", "MultiInputParamApi", "", "")
	return
}

func CreateMultiInputParamApiResponse() (response *MultiInputParamApiResponse) {
	response = &MultiInputParamApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
