package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	common "github.com/apiheat/akamai-cli-common/v4"
// 	edgegrid "github.com/apiheat/go-edgegrid"
// 	"github.com/stretchr/testify/assert"
// )

// //setupEdgeClient prepares and inits client for making all calls towards Akamai's APIs
// func setupEdgeClient(aka, section, edgercFile, baseURL string) *edgegrid.Client {

// 	// Provide struct details needed for apiClient init
// 	apiClientOpts := &edgegrid.ClientOptions{}
// 	apiClientOpts.ConfigPath = edgercFile
// 	apiClientOpts.ConfigSection = section
// 	// apiClientOpts.DebugLevel = "debug"
// 	apiClientOpts.AccountSwitchKey = aka

// 	// Initialize client for tests with init options
// 	apiClient, _ := common.EdgeClientInit(apiClientOpts)
// 	apiClient.SetBaseURL(baseURL, true)

// 	return apiClient
// }

// //TestListNetworkLists checks if listing all network lists works
// func TestListNetworkLists(t *testing.T) {

// 	response := `{"networkLists":[{"networkListType":"networkListResponse","accessControlGroup":"KSD\nwith ION 3-13H1234","name":"General List","elementCount":3011,"syncPoint":22,"type":"IP","uniqueId":"25614_GENERALLIST","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/25614_GENERALLIST"},"statusInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"PUT"}}},{"networkListType":"networkListResponse","account":"Kona\nSecurity Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","name":"Ec2 Akamai Network List","elementCount":235,"readOnly":true,"syncPoint":65,"type":"IP","uniqueId":"1024_AMAZONELASTICCOMPUTECLOU","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}},{"networkListType":"networkListResponse","accessControlGroup":"KSD\nTest - 3-13H5523","name":"GeoList_1913New","elementCount":16,"syncPoint":2,"type":"GEO","uniqueId":"26732_GEOLIST1913","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913"},"statusInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913","method":"PUT"}}}],"links":{"create":{"href":"/network-list/v2/network-lists/","method":"POST"}}}`
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		body, _ := ioutil.ReadAll(r.Body)
// 		assert.Empty(t, string(body), "Request body should be empty")
// 		assert.Equal(t, "GET", r.Method, "Method should be GET")

// 		fmt.Fprintln(w, response)
// 	}))

// 	// Init API client
// 	apiClient := setupEdgeClient("", "dummy", ".edgerc-test", server.URL)

// 	listNetListOptsv2 := edgegrid.ListNetworkListsOptionsv2{}
// 	listNetListOptsv2.Search = "" // Since we are listing all we do not filter results

// 	apiResp, clientResp, reqErr := apiClient.NetworkListsv2.ListNetworkLists(listNetListOptsv2)

// 	if reqErr != nil {
// 		fmt.Println(reqErr)
// 	}

// 	var expectedType *[]edgegrid.NetworkListv2

// 	assert.IsType(t, expectedType, apiResp)
// 	assert.Equal(t, 200, clientResp.Response.StatusCode, "Response should be 200 OK")

// 	defer server.Close()
// }

// //TestGetNetworkListById checks if listing specific network list by-id works
// func TestGetNetworkListById(t *testing.T) {

// 	response := `{"networkLists":[{"networkListType":"networkListResponse","accessControlGroup":"KSD\nwith ION 3-13H1234","name":"General List","elementCount":3011,"syncPoint":22,"type":"IP","uniqueId":"25614_GENERALLIST","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/25614_GENERALLIST"},"statusInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"PUT"}}},{"networkListType":"networkListResponse","account":"Kona\nSecurity Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","name":"Ec2 Akamai Network List","elementCount":235,"readOnly":true,"syncPoint":65,"type":"IP","uniqueId":"1024_AMAZONELASTICCOMPUTECLOU","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}},{"networkListType":"networkListResponse","accessControlGroup":"KSD\nTest - 3-13H5523","name":"GeoList_1913New","elementCount":16,"syncPoint":2,"type":"GEO","uniqueId":"26732_GEOLIST1913","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913"},"statusInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913","method":"PUT"}}}],"links":{"create":{"href":"/network-list/v2/network-lists/","method":"POST"}}}`
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		body, _ := ioutil.ReadAll(r.Body)
// 		assert.Empty(t, string(body), "Request body should be empty")
// 		assert.Equal(t, "GET", r.Method, "Request method should be GET")
// 		assert.Contains(t, r.URL.String(), "123_TEST", "Request URL should contain list ID")
// 		fmt.Fprintln(w, response)
// 	}))

// 	// Init API client
// 	apiClient := setupEdgeClient("", "dummy", ".edgerc-test", server.URL)

// 	listNetListOptsv2 := edgegrid.ListNetworkListsOptionsv2{}
// 	listNetListOptsv2.Search = "" // Since we are listing all we do not filter results

// 	apiResp, clientResp, reqErr := apiClient.NetworkListsv2.GetNetworkList("123_TEST", listNetListOptsv2)

// 	if reqErr != nil {
// 		fmt.Println(reqErr)
// 	}

// 	var expectedType *edgegrid.NetworkListv2

// 	assert.IsType(t, expectedType, apiResp)
// 	assert.Equal(t, 200, clientResp.Response.StatusCode, "Response should be 200 OK")

// 	defer server.Close()
// }

// //TestCreateNetworkList creates new network list
// func TestCreateNetworkList(t *testing.T) {

// 	response := `{"name":"name-of-netlist","uniqueId":"1024_AMAZONELASTICCOMPUTECLOU","syncPoint":65,"type":"IP","networkListType":"networkListResponse","account":"Kona Security Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","elementCount":13,"readOnly":true,"list":["13.125.0.0/16","13.126.0.0/15","13.210.0.0/15","13.228.0.0/15","13.230.0.0/15","13.232.0.0/14","13.236.0.0/14","13.250.0.0/15","13.54.0.0/15","13.56.0.0/16","13.57.0.0/16","13.58.0.0/15","174.129.0.0/16"],"links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}}
// 	`
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(201)

// 		body, _ := ioutil.ReadAll(r.Body)
// 		assert.Equal(t, "{\"name\":\"name-of-netlist\",\"type\":\"IP\",\"description\":\"desc-by-test\"}", string(body), "Request body should be empty")
// 		assert.Equal(t, "POST", r.Method, "Request method should be GET")
// 		fmt.Fprintln(w, response)
// 	}))

// 	//--Init API client
// 	apiClient := setupEdgeClient("", "dummy", ".edgerc-test", server.URL)

// 	//--Create options
// 	newNetworkListOpst := edgegrid.NetworkListsOptionsv2{}
// 	newNetworkListOpst.Description = "desc-by-test"
// 	newNetworkListOpst.Name = "name-of-netlist"
// 	newNetworkListOpst.Type = "IP"

// 	apiResp, clientResp, reqErr := apiClient.NetworkListsv2.CreateNetworkList(newNetworkListOpst)

// 	if reqErr != nil {
// 		fmt.Println(reqErr)
// 	}

// 	var expectedType *edgegrid.NetworkListv2

// 	assert.IsType(t, expectedType, apiResp)
// 	assert.Equal(t, 201, clientResp.Response.StatusCode, "Response should be 201 OK")

// 	defer server.Close()
// }

// //TestAddNetworNetworkListElement adds network elements to list
// func TestAddNetworNetworkListElement(t *testing.T) {
// 	response := `{"name":"Ec2 Akamai Network List","uniqueId":"345_BOTLIST","syncPoint":65,"type":"IP","networkListType":"networkListResponse","account":"Kona Security Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","elementCount":13,"readOnly":true,"list":["1.2.3.4/32","13.126.0.0/15","13.210.0.0/15","13.228.0.0/15","13.230.0.0/15","13.232.0.0/14","13.236.0.0/14","13.250.0.0/15","13.54.0.0/15","13.56.0.0/16","13.57.0.0/16","13.58.0.0/15","174.129.0.0/16"],"links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}}
// 	`
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(201)

// 		body, _ := ioutil.ReadAll(r.Body)
// 		assert.Equal(t, "{\"list\":[\"1.2.3.4/32\"]}", string(body), "Request body should contain list of addresses to activate")
// 		assert.NotEmpty(t, string(body), "Request body should not be empty")
// 		assert.Equal(t, "POST", r.Method, "Request method should be GET")
// 		assert.Contains(t, r.URL.String(), "345_BOTLIST/append", "Request URL should contain list ID")
// 		fmt.Fprintln(w, response)
// 	}))

// 	//--Init API client
// 	apiClient := setupEdgeClient("", "dummy", ".edgerc-test", server.URL)

// 	//--Modify existing network list
// 	itemsToAdd := []string{"1.2.3.4/32"}
// 	editListOpts := edgegrid.NetworkListsOptionsv2{
// 		List: itemsToAdd,
// 	}

// 	apiResp, clientResp, reqErr := apiClient.NetworkListsv2.AppendListNetworkList("345_BOTLIST", editListOpts)

// 	if reqErr != nil {
// 		fmt.Println(reqErr)
// 	}

// 	var expectedType *edgegrid.NetworkListv2

// 	assert.IsType(t, expectedType, apiResp)
// 	assert.Equal(t, 201, clientResp.Response.StatusCode, "Response should be 201 OK")
// 	assert.Equal(t, "345_BOTLIST", apiResp.UniqueID, "UniqueId should match 345_BOTLIST")

// 	defer server.Close()
// }

// //TestRemoveNetworkListItem removes item from network list elements
// func TestRemoveNetworkListItem(t *testing.T) {
// 	response := `{"name":"Ec2 Akamai Network List","uniqueId":"345_BOTLIST","syncPoint":65,"type":"IP","networkListType":"networkListResponse","account":"Kona Security Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","elementCount":13,"readOnly":true,"list":["13.126.0.0/15","13.210.0.0/15","13.228.0.0/15","13.230.0.0/15","13.232.0.0/14","13.236.0.0/14","13.250.0.0/15","13.54.0.0/15","13.56.0.0/16","13.57.0.0/16","13.58.0.0/15","174.129.0.0/16"],"links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}}
// 	`
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(200)

// 		body, _ := ioutil.ReadAll(r.Body)
// 		assert.Empty(t, string(body), "Request body should be empty")
// 		assert.Equal(t, "DELETE", r.Method, "Request method should be DELETE")
// 		assert.Contains(t, r.URL.String(), "345_BOTLIST/elements", "Request URL should contain list ID")
// 		fmt.Fprintln(w, response)
// 	}))

// 	//--Init API client
// 	apiClient := setupEdgeClient("", "dummy", ".edgerc-test", server.URL)
// 	apiResp, clientResp, reqErr := apiClient.NetworkListsv2.RemoveNetworkListElement("345_BOTLIST", "1.2.3.4/32")

// 	if reqErr != nil {
// 		fmt.Println(reqErr)
// 	}

// 	var expectedType *edgegrid.NetworkListv2

// 	assert.IsType(t, expectedType, apiResp)
// 	assert.Equal(t, 200, clientResp.Response.StatusCode, "Response should be 200 OK")
// 	assert.Equal(t, "345_BOTLIST", apiResp.UniqueID, "UniqueId should match 345_BOTLIST")
// 	assert.NotContains(t, apiResp.List, "1.2.3.4/32")

// 	defer server.Close()
// }

// func TestActivateNetworkList(t *testing.T) {
// 	response := `{"activationId":12345,"activationComments":"test-activation","activationStatus":"PENDING_ACTIVATION","syncPoint":5,"uniqueId":"345_BOTLIST","fast":false,"dispatchCount":1,"links":{"appendItems":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST/append","method":"POST"},"retrieve":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST"},"statusInProduction":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/status"},"statusInStaging":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST/environments/STAGING/status"},"syncPointHistory":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST/sync-points/5/history"},"update":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST","method":"PUT"},"activationDetails":{"href":"/network-list/v2/network-lists/activations/12345/"}}}`
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(200)

// 		body, _ := ioutil.ReadAll(r.Body)
// 		assert.NotEmpty(t, string(body), "Request body should not be empty")
// 		assert.Equal(t, `{"comments":"test-activation","notificationRecipients":["dummy@mailinator.com"],"fast":true}`, string(body), "Request body should contain list of addresses to activate")
// 		assert.Equal(t, "POST", r.Method, "Request method should be POST")
// 		assert.Equal(t, "/network-list/v2/network-lists/345_BOTLIST/environments/production/activate", r.URL.String(), "Request URL should target activation env with list ID")
// 		fmt.Fprintln(w, response)
// 	}))

// 	//--Init API client
// 	apiClient := setupEdgeClient("", "dummy", ".edgerc-test", server.URL)

// 	actNetworkListOpts := edgegrid.NetworkListActivationOptsv2{
// 		Comments: "test-activation",
// 		Fast:     true,
// 		NotificationRecipients: []string{"dummy@mailinator.com"},
// 	}
// 	apiResp, clientResp, reqErr := apiClient.NetworkListsv2.ActivateNetworkList("345_BOTLIST", edgegrid.Production, actNetworkListOpts)

// 	if reqErr != nil {
// 		fmt.Println(reqErr)
// 	}

// 	var expectedType *edgegrid.NetworkListActivationStatusv2

// 	assert.IsType(t, expectedType, apiResp)
// 	assert.Equal(t, 200, clientResp.Response.StatusCode, "Response should be 200 OK")
// 	assert.Equal(t, "345_BOTLIST", apiResp.UniqueID, "UniqueId should match 345_BOTLIST")

// 	defer server.Close()
// }
