package azfile

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// fileClient is the client for the File methods of the Azfile service.
type fileClient struct {
	managementClient
}

// newFileClient creates an instance of the fileClient client.
func newFileClient(url url.URL, p pipeline.Pipeline) fileClient {
	return fileClient{newManagementClient(url, p)}
}

// AbortCopy aborts a pending Copy File operation, and leaves a destination file with zero length and full metadata.
//
// copyID is the copy identifier provided in the x-ms-copy-id header of the original Copy File operation. timeout is
// the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client fileClient) AbortCopy(ctx context.Context, copyID string, copyActionAbortConstant string, timeout *int32) (*FileAbortCopyResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.abortCopyPreparer(copyID, copyActionAbortConstant, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.abortCopyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileAbortCopyResponse), err
}

// abortCopyPreparer prepares the AbortCopy request.
func (client fileClient) abortCopyPreparer(copyID string, copyActionAbortConstant string, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("copyid", copyID)
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "copy")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-copy-action", copyActionAbortConstant)
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// abortCopyResponder handles the response to the AbortCopy request.
func (client fileClient) abortCopyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	return &FileAbortCopyResponse{rawResponse: resp.Response()}, err
}

// Copy copies a blob or file to a destination file within the storage account.
//
// copySource is specifies the URL of the source file or blob, up to 2 KB in length. To copy a file to another file
// within the same storage account, you may use Shared Key to authenticate the source file. If you are copying a file
// from another storage account, or if you are copying a blob from the same storage account or another storage account,
// then you must authenticate the source file or blob using a shared access signature. If the source is a public blob,
// no authentication is required to perform the copy operation. A file in a share snapshot can also be specified as a
// copy source. timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
// Metadata names must adhere to the naming rules for C# identifiers.
func (client fileClient) Copy(ctx context.Context, copySource string, timeout *int32, metadata map[string]string) (*FileCopyResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.copyPreparer(copySource, timeout, metadata)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.copyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileCopyResponse), err
}

// copyPreparer prepares the Copy request.
func (client fileClient) copyPreparer(copySource string, timeout *int32, metadata map[string]string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	req.Header.Set("x-ms-copy-source", copySource)
	return req, nil
}

// copyResponder handles the response to the Copy request.
func (client fileClient) copyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	return &FileCopyResponse{rawResponse: resp.Response()}, err
}

// Create creates a new file or replaces a file. Note it only initializes the file with no content.
//
// fileContentLength is specifies the maximum size for the file, up to 1 TB. fileTypeConstant is dummy constant
// parameter, file type can only be file. timeout is the timeout parameter is expressed in seconds. For more
// information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> fileContentType is sets the MIME content type of the file. The default
// type is 'application/octet-stream'. fileContentEncoding is specifies which content encodings have been applied to
// the file. fileContentLanguage is specifies the natural languages used by this resource. fileCacheControl is sets the
// file's cache control. The File service stores this value but does not use or modify it. fileContentMD5 is sets the
// file's MD5 hash. fileContentDisposition is sets the file's Content-Disposition header. metadata is a name-value pair
// to associate with a file storage object. Metadata names must adhere to the naming rules for C# identifiers.
func (client fileClient) Create(ctx context.Context, fileContentLength int64, fileTypeConstant string, timeout *int32, fileContentType *string, fileContentEncoding *string, fileContentLanguage *string, fileCacheControl *string, fileContentMD5 *string, fileContentDisposition *string, metadata map[string]string) (*FileCreateResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createPreparer(fileContentLength, fileTypeConstant, timeout, fileContentType, fileContentEncoding, fileContentLanguage, fileCacheControl, fileContentMD5, fileContentDisposition, metadata)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileCreateResponse), err
}

// createPreparer prepares the Create request.
func (client fileClient) createPreparer(fileContentLength int64, fileTypeConstant string, timeout *int32, fileContentType *string, fileContentEncoding *string, fileContentLanguage *string, fileCacheControl *string, fileContentMD5 *string, fileContentDisposition *string, metadata map[string]string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	req.Header.Set("x-ms-content-length", fmt.Sprintf("%v", fileContentLength))
	req.Header.Set("x-ms-type", fileTypeConstant)
	if fileContentType != nil {
		req.Header.Set("x-ms-content-type", *fileContentType)
	}
	if fileContentEncoding != nil {
		req.Header.Set("x-ms-content-encoding", *fileContentEncoding)
	}
	if fileContentLanguage != nil {
		req.Header.Set("x-ms-content-language", *fileContentLanguage)
	}
	if fileCacheControl != nil {
		req.Header.Set("x-ms-cache-control", *fileCacheControl)
	}
	if fileContentMD5 != nil {
		req.Header.Set("x-ms-content-md5", *fileContentMD5)
	}
	if fileContentDisposition != nil {
		req.Header.Set("x-ms-content-disposition", *fileContentDisposition)
	}
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	return req, nil
}

// createResponder handles the response to the Create request.
func (client fileClient) createResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	return &FileCreateResponse{rawResponse: resp.Response()}, err
}

// Delete removes the file from the storage account.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client fileClient) Delete(ctx context.Context, timeout *int32) (*FileDeleteResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.deletePreparer(timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileDeleteResponse), err
}

// deletePreparer prepares the Delete request.
func (client fileClient) deletePreparer(timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("DELETE", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client fileClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	return &FileDeleteResponse{rawResponse: resp.Response()}, err
}

// Get reads or downloads a file from the system, including its metadata and properties.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> rangeParameter is return file data only from the specified byte range.
// rangeGetContentMD5 is when this header is set to true and specified together with the Range header, the service
// returns the MD5 hash for the range, as long as the range is less than or equal to 4 MB in size.
func (client fileClient) Get(ctx context.Context, timeout *int32, rangeParameter *string, rangeGetContentMD5 *bool) (*GetResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPreparer(timeout, rangeParameter, rangeGetContentMD5)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetResponse), err
}

// getPreparer prepares the Get request.
func (client fileClient) getPreparer(timeout *int32, rangeParameter *string, rangeGetContentMD5 *bool) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if rangeParameter != nil {
		req.Header.Set("x-ms-range", *rangeParameter)
	}
	if rangeGetContentMD5 != nil {
		req.Header.Set("x-ms-range-get-content-md5", fmt.Sprintf("%v", *rangeGetContentMD5))
	}
	return req, nil
}

// getResponder handles the response to the Get request.
func (client fileClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusPartialContent)
	if resp == nil {
		return nil, err
	}
	return &GetResponse{rawResponse: resp.Response()}, err
}

// GetMetadata returns all user-defined metadata for the specified file
//
// sharesnapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot
// to query to retrieve the properties. timeout is the timeout parameter is expressed in seconds. For more information,
// see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client fileClient) GetMetadata(ctx context.Context, sharesnapshot *string, timeout *int32) (*FileGetMetadataResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getMetadataPreparer(sharesnapshot, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getMetadataResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileGetMetadataResponse), err
}

// getMetadataPreparer prepares the GetMetadata request.
func (client fileClient) getMetadataPreparer(sharesnapshot *string, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if sharesnapshot != nil {
		params.Set("sharesnapshot", *sharesnapshot)
	}
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "metadata")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getMetadataResponder handles the response to the GetMetadata request.
func (client fileClient) getMetadataResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	return &FileGetMetadataResponse{rawResponse: resp.Response()}, err
}

// GetProperties returns all user-defined metadata, standard HTTP properties, and system properties for the file. It
// does not return the content of the file.
//
// sharesnapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot
// to query to retrieve the properties. timeout is the timeout parameter is expressed in seconds. For more information,
// see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client fileClient) GetProperties(ctx context.Context, sharesnapshot *string, timeout *int32) (*FileGetPropertiesResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPropertiesPreparer(sharesnapshot, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileGetPropertiesResponse), err
}

// getPropertiesPreparer prepares the GetProperties request.
func (client fileClient) getPropertiesPreparer(sharesnapshot *string, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("HEAD", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if sharesnapshot != nil {
		params.Set("sharesnapshot", *sharesnapshot)
	}
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getPropertiesResponder handles the response to the GetProperties request.
func (client fileClient) getPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	return &FileGetPropertiesResponse{rawResponse: resp.Response()}, err
}

// ListRanges returns the list of valid ranges for a file.
//
// sharesnapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot
// to query to retrieve the properties. timeout is the timeout parameter is expressed in seconds. For more information,
// see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> rangeParameter is specifies the range of bytes over which to list ranges,
// inclusively.
func (client fileClient) ListRanges(ctx context.Context, sharesnapshot *string, timeout *int32, rangeParameter *string) (*Ranges, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.listRangesPreparer(sharesnapshot, timeout, rangeParameter)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listRangesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Ranges), err
}

// listRangesPreparer prepares the ListRanges request.
func (client fileClient) listRangesPreparer(sharesnapshot *string, timeout *int32, rangeParameter *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if sharesnapshot != nil {
		params.Set("sharesnapshot", *sharesnapshot)
	}
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "rangelist")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if rangeParameter != nil {
		req.Header.Set("x-ms-range", *rangeParameter)
	}
	return req, nil
}

// listRangesResponder handles the response to the ListRanges request.
func (client fileClient) listRangesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Ranges{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// PutRange writes a range of bytes to a file.
//
// rangeParameter is specifies the range of bytes to be written. Both the start and end of the range must be specified.
// For an update operation, the range can be up to 4 MB in size. For a clear operation, the range can be up to the
// value of the file's full size. The File service accepts only a single byte range for the Range and 'x-ms-range'
// headers, and the byte range must be specified in the following format: bytes=startByte-endByte. fileRangeWrite is
// specify one of the following options: - Update: Writes the bytes specified by the request body into the specified
// range. The Range and Content-Length headers must match to perform the update. - Clear: Clears the specified range
// and releases the space used in storage for that range. To clear a range, set the Content-Length header to zero, and
// set the Range header to a value that indicates the range to clear, up to maximum file size. contentLength is
// specifies the number of bytes being transmitted in the request body. When the x-ms-write header is set to clear, the
// value of this header must be set to zero. optionalbody is initial data. optionalbody will be closed upon successful
// return. Callers should ensure closure when receiving an error.timeout is the timeout parameter is expressed in
// seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> contentMD5 is an MD5 hash of the content. This hash is used to verify the
// integrity of the data during transport. When the Content-MD5 header is specified, the File service compares the hash
// of the content that has arrived with the header value that was sent. If the two hashes do not match, the operation
// will fail with error code 400 (Bad Request).
func (client fileClient) PutRange(ctx context.Context, rangeParameter string, fileRangeWrite FileRangeWriteType, contentLength int64, body io.ReadSeeker, timeout *int32, contentMD5 *string) (*FilePutRangeResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.putRangePreparer(rangeParameter, fileRangeWrite, contentLength, body, timeout, contentMD5)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.putRangeResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FilePutRangeResponse), err
}

// putRangePreparer prepares the PutRange request.
func (client fileClient) putRangePreparer(rangeParameter string, fileRangeWrite FileRangeWriteType, contentLength int64, body io.ReadSeeker, timeout *int32, contentMD5 *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, body)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "range")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-range", rangeParameter)
	req.Header.Set("x-ms-write", fmt.Sprintf("%v", fileRangeWrite))
	req.Header.Set("Content-Length", fmt.Sprintf("%v", contentLength))
	if contentMD5 != nil {
		req.Header.Set("Content-MD5", *contentMD5)
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// putRangeResponder handles the response to the PutRange request.
func (client fileClient) putRangeResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	return &FilePutRangeResponse{rawResponse: resp.Response()}, err
}

// SetMetadata updates user-defined metadata for the specified file.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
// Metadata names must adhere to the naming rules for C# identifiers.
func (client fileClient) SetMetadata(ctx context.Context, timeout *int32, metadata map[string]string) (*FileSetMetadataResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setMetadataPreparer(timeout, metadata)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setMetadataResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileSetMetadataResponse), err
}

// setMetadataPreparer prepares the SetMetadata request.
func (client fileClient) setMetadataPreparer(timeout *int32, metadata map[string]string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "metadata")
	req.URL.RawQuery = params.Encode()
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// setMetadataResponder handles the response to the SetMetadata request.
func (client fileClient) setMetadataResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	return &FileSetMetadataResponse{rawResponse: resp.Response()}, err
}

// SetProperties sets system properties on the file.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> fileContentLength is resizes a file to the specified size. If the
// specified byte value is less than the current size of the file, then all ranges above the specified byte value are
// cleared. fileContentType is sets the MIME content type of the file. The default type is 'application/octet-stream'.
// fileContentEncoding is specifies which content encodings have been applied to the file. fileContentLanguage is
// specifies the natural languages used by this resource. fileCacheControl is sets the file's cache control. The File
// service stores this value but does not use or modify it. fileContentMD5 is sets the file's MD5 hash.
// fileContentDisposition is sets the file's Content-Disposition header.
func (client fileClient) SetProperties(ctx context.Context, timeout *int32, fileContentLength *int64, fileContentType *string, fileContentEncoding *string, fileContentLanguage *string, fileCacheControl *string, fileContentMD5 *string, fileContentDisposition *string) (*FileSetPropertiesResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setPropertiesPreparer(timeout, fileContentLength, fileContentType, fileContentEncoding, fileContentLanguage, fileCacheControl, fileContentMD5, fileContentDisposition)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FileSetPropertiesResponse), err
}

// setPropertiesPreparer prepares the SetProperties request.
func (client fileClient) setPropertiesPreparer(timeout *int32, fileContentLength *int64, fileContentType *string, fileContentEncoding *string, fileContentLanguage *string, fileCacheControl *string, fileContentMD5 *string, fileContentDisposition *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "properties")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if fileContentLength != nil {
		req.Header.Set("x-ms-content-length", fmt.Sprintf("%v", *fileContentLength))
	}
	if fileContentType != nil {
		req.Header.Set("x-ms-content-type", *fileContentType)
	}
	if fileContentEncoding != nil {
		req.Header.Set("x-ms-content-encoding", *fileContentEncoding)
	}
	if fileContentLanguage != nil {
		req.Header.Set("x-ms-content-language", *fileContentLanguage)
	}
	if fileCacheControl != nil {
		req.Header.Set("x-ms-cache-control", *fileCacheControl)
	}
	if fileContentMD5 != nil {
		req.Header.Set("x-ms-content-md5", *fileContentMD5)
	}
	if fileContentDisposition != nil {
		req.Header.Set("x-ms-content-disposition", *fileContentDisposition)
	}
	return req, nil
}

// setPropertiesResponder handles the response to the SetProperties request.
func (client fileClient) setPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	return &FileSetPropertiesResponse{rawResponse: resp.Response()}, err
}
