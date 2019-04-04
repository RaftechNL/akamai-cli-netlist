package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	"github.com/stretchr/testify/assert"
)

//setupEdgeClient prepares and inits client for making all calls towards Akamai's APIs
func setupEdgeClient(aka, section, edgercFile, baseURL string) *edgegrid.Client {

	// Provide struct details needed for apiClient init
	apiClientOpts := &edgegrid.ClientOptions{}
	apiClientOpts.ConfigPath = edgercFile
	apiClientOpts.ConfigSection = section
	// apiClientOpts.DebugLevel = "debug"
	apiClientOpts.AccountSwitchKey = aka

	// Initialize client for tests with init options
	apiClient, _ := common.EdgeClientInit(apiClientOpts)
	apiClient.SetBaseURL(baseURL, true)

	return apiClient
}

//TestListNetworkLists checks if listing all network lists works
func TestListNetworkLists(t *testing.T) {

	response := `{"networkLists":[{"networkListType":"networkListResponse","accessControlGroup":"KSD\nwith ION 3-13H1234","name":"General List","elementCount":3011,"syncPoint":22,"type":"IP","uniqueId":"25614_GENERALLIST","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/25614_GENERALLIST"},"statusInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"PUT"}}},{"networkListType":"networkListResponse","account":"Kona\nSecurity Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","name":"Ec2 Akamai Network List","elementCount":235,"readOnly":true,"syncPoint":65,"type":"IP","uniqueId":"1024_AMAZONELASTICCOMPUTECLOU","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}},{"networkListType":"networkListResponse","accessControlGroup":"KSD\nTest - 3-13H5523","name":"GeoList_1913New","elementCount":16,"syncPoint":2,"type":"GEO","uniqueId":"26732_GEOLIST1913","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913"},"statusInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913","method":"PUT"}}}],"links":{"create":{"href":"/network-list/v2/network-lists/","method":"POST"}}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body, _ := ioutil.ReadAll(r.Body)
		assert.Equal(t, "", string(body), "Request body should be empty")
		assert.Equal(t, "GET", r.Method, "Method should be GET")

		fmt.Fprintln(w, response)
	}))

	// Init API client
	apiClient := setupEdgeClient("", "dummy", ".edgerc-test", server.URL)

	listNetListOptsv2 := edgegrid.ListNetworkListsOptionsv2{}
	listNetListOptsv2.Search = "" // Since we are listing all we do not filter results

	apiResp, clientResp, reqErr := apiClient.NetworkListsv2.ListNetworkLists(listNetListOptsv2)

	if reqErr != nil {
		fmt.Println(reqErr)
	}

	var expectedType *[]edgegrid.NetworkListv2

	assert.IsType(t, expectedType, apiResp)
	assert.Equal(t, 200, clientResp.Response.StatusCode, "Response should be 200 OK")

	defer server.Close()
}

//TestGetNetworkListById checks if listing specific network list by-id works
func TestGetNetworkListById(t *testing.T) {

	response := `{"networkLists":[{"networkListType":"networkListResponse","accessControlGroup":"KSD\nwith ION 3-13H1234","name":"General List","elementCount":3011,"syncPoint":22,"type":"IP","uniqueId":"25614_GENERALLIST","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/25614_GENERALLIST"},"statusInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"PUT"}}},{"networkListType":"networkListResponse","account":"Kona\nSecurity Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","name":"Ec2 Akamai Network List","elementCount":235,"readOnly":true,"syncPoint":65,"type":"IP","uniqueId":"1024_AMAZONELASTICCOMPUTECLOU","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}},{"networkListType":"networkListResponse","accessControlGroup":"KSD\nTest - 3-13H5523","name":"GeoList_1913New","elementCount":16,"syncPoint":2,"type":"GEO","uniqueId":"26732_GEOLIST1913","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913"},"statusInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913","method":"PUT"}}}],"links":{"create":{"href":"/network-list/v2/network-lists/","method":"POST"}}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body, _ := ioutil.ReadAll(r.Body)
		assert.Equal(t, "", string(body), "Request body should be empty")
		assert.Equal(t, "GET", r.Method, "Request method should be GET")
		assert.Contains(t, r.URL.String(), "123_TEST", "Request URL should contain list ID")
		fmt.Fprintln(w, response)
	}))

	// Init API client
	apiClient := setupEdgeClient("", "dummy", ".edgerc-test", server.URL)

	listNetListOptsv2 := edgegrid.ListNetworkListsOptionsv2{}
	listNetListOptsv2.Search = "" // Since we are listing all we do not filter results

	apiResp, clientResp, reqErr := apiClient.NetworkListsv2.GetNetworkList("123_TEST", listNetListOptsv2)

	if reqErr != nil {
		fmt.Println(reqErr)
	}

	var expectedType *edgegrid.NetworkListv2

	assert.IsType(t, expectedType, apiResp)
	assert.Equal(t, 200, clientResp.Response.StatusCode, "Response should be 200 OK")

	defer server.Close()
}

//
// func TestCreateNetworkList(t *testing.T) {

// 	response := `{"networkLists":[{"networkListType":"networkListResponse","accessControlGroup":"KSD\nwith ION 3-13H1234","name":"General List","elementCount":3011,"syncPoint":22,"type":"IP","uniqueId":"25614_GENERALLIST","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/25614_GENERALLIST"},"statusInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"PUT"}}},{"networkListType":"networkListResponse","account":"Kona\nSecurity Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","name":"Ec2 Akamai Network List","elementCount":235,"readOnly":true,"syncPoint":65,"type":"IP","uniqueId":"1024_AMAZONELASTICCOMPUTECLOU","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}},{"networkListType":"networkListResponse","accessControlGroup":"KSD\nTest - 3-13H5523","name":"GeoList_1913New","elementCount":16,"syncPoint":2,"type":"GEO","uniqueId":"26732_GEOLIST1913","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913"},"statusInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913","method":"PUT"}}}],"links":{"create":{"href":"/network-list/v2/network-lists/","method":"POST"}}}`
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		body, _ := ioutil.ReadAll(r.Body)
// 		assert.Equal(t, "", string(body), "Request body should be empty")
// 		assert.Equal(t, "GET", r.Method, "Request method should be GET")
// 		assert.Contains(t, r.URL.String(), "123_TEST", "Request URL should contain list ID")
// 		fmt.Fprintln(w, response)
// 	}))

// 	apiClient.SetBaseURL(server.URL, true)

// 	opt := CreateNetworkListOptions{
// 		Name:        "testnetlist",
// 		Type:        "IP",
// 		Description: "created-by-test",
// 	}

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

// 	mux.HandleFunc("/network-list/v1/network_lists", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(201)
// 		testMethod(t, r, "POST")
// 		fmt.Fprintf(w, `{
// 			"status": 201,
// 			"unique-id": "345_BOTLIST",
// 			"links": [
// 				{
// 				   "rel": "get 345_BOTLIST",
// 				   "href": "/network-list/v1/network_lists/345_BOTLIST"
// 				}
// 			],
// 			"sync-point": 0
// 		 }`)
// 	})

// 	want := &NetworkListResponse{
// 		Status:   201,
// 		UniqueID: "345_BOTLIST",
// 		Links: []AkamaiNetworkListLinks{
// 			AkamaiNetworkListLinks{
// 				Rel:  "get 345_BOTLIST",
// 				Href: "/network-list/v1/network_lists/345_BOTLIST",
// 			},
// 		},
// 		SyncPoint: 0,
// 	}

// 	received, resp, _ := client.NetworkLists.CreateNetworkList(opt)
// 	assert.Equal(t, 201, resp.Response.StatusCode, "Response should be 201")
// 	assert.IsType(t, want, received, "Should be type of NetworkListResponse")
// 	assert.Equal(t, want, received, "Response should contain details of network lists created")
// }

// func TestModifyNetworkList(t *testing.T) {
// 	mux, server, client := setup()
// 	defer teardown(server)

// 	// Modify existing network list
// 	modifyItems := []string{"4.4.3.4/32"}
// 	editListOpts := AkamaiNetworkList{
// 		Name: "simple-new",
// 		List: modifyItems,
// 	}

// 	mux.HandleFunc("/network-list/v1/network_lists/345_BOTLIST", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(200)
// 		testMethod(t, r, "PUT")
// 		fmt.Fprintf(w, `{
// 			"status": 201,
// 			"unique-id": "345_BOTLIST",
// 			"links": [
// 				{
// 				   "rel": "get 345_BOTLIST",
// 				   "href": "/network-list/v1/network_lists/345_BOTLIST"
// 				}
// 			],
// 			"sync-point": 0
// 		 }`)
// 	})

// 	want := &NetworkListResponse{
// 		Status:   201,
// 		UniqueID: "345_BOTLIST",
// 		Links: []AkamaiNetworkListLinks{
// 			AkamaiNetworkListLinks{
// 				Rel:  "get 345_BOTLIST",
// 				Href: "/network-list/v1/network_lists/345_BOTLIST",
// 			},
// 		},
// 		SyncPoint: 0,
// 	}

// 	received, resp, _ := client.NetworkLists.ModifyNetworkList("345_BOTLIST", editListOpts)
// 	assert.Equal(t, 200, resp.Response.StatusCode, "Response should be 200")
// 	assert.IsType(t, want, received, "Should be type of NetworkListResponse")
// 	assert.Equal(t, want, received, "Response should contain details of network lists created")
// }

// func TestAddNetworkListItems(t *testing.T) {
// 	mux, server, client := setup()
// 	defer teardown(server)

// 	modifyItems := []string{"4.4.3.4/32"}
// 	modifyItemsOpts := CreateNetworkListOptions{
// 		List: modifyItems,
// 	}

// 	wantResponseCode := 202

// 	mux.HandleFunc("/network-list/v1/network_lists/345_BOTLIST", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(wantResponseCode)
// 		testMethod(t, r, "POST")
// 		fmt.Fprintf(w, `{
// 			"message": "Elements successfully appended to the list",
// 			"status": 202,
// 			"links": []
// 		 }`)
// 	})

// 	want := &NetworkListResponse{
// 		Status:  202,
// 		Links:   []AkamaiNetworkListLinks{},
// 		Message: "Elements successfully appended to the list",
// 	}

// 	received, resp, _ := client.NetworkLists.AddNetworkListItems("345_BOTLIST", modifyItemsOpts)
// 	assert.Equal(t, wantResponseCode, resp.Response.StatusCode, "Response should be 202")
// 	assert.IsType(t, want, received, "Should be type of NetworkListResponse")
// 	assert.Equal(t, want, received, "Response should contain details of network lists created")
// }

// func TestAddNetworkListElement(t *testing.T) {
// 	mux, server, client := setup()
// 	defer teardown(server)

// 	wantResponseCode := 201

// 	mux.HandleFunc("/network-list/v1/network_lists/345_BOTLIST/element", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(201)
// 		testMethod(t, r, "PUT")
// 		fmt.Fprintf(w, `{
// 			"message": "Elements successfully appended to the list",
// 			"status": 201,
// 			"links": []
// 		 }`)
// 	})

// 	want := &NetworkListResponse{
// 		Status:  wantResponseCode,
// 		Links:   []AkamaiNetworkListLinks{},
// 		Message: "Elements successfully appended to the list",
// 	}

// 	received, resp, _ := client.NetworkLists.AddNetworkListElement("345_BOTLIST", "1.2.3.5/32")
// 	assert.Equal(t, wantResponseCode, resp.Response.StatusCode, "Response should be 202")
// 	assert.IsType(t, want, received, "Should be type of NetworkListResponse")
// 	assert.Equal(t, want, received, "Response should contain details of network lists created")
// }

// func TestRemoveNetworkListItem(t *testing.T) {
// 	mux, server, client := setup()
// 	defer teardown(server)

// 	wantResponseCode := 200

// 	mux.HandleFunc("/network-list/v1/network_lists/345_BOTLIST/element", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(200)
// 		testMethod(t, r, "DELETE")
// 		fmt.Fprintf(w, `{
// 			"message": "Element 1.2.3.4 successfully deleted from list",
// 			"status": 200,
// 			"links": []
// 		 }`)
// 	})

// 	want := &NetworkListResponse{
// 		Status:  wantResponseCode,
// 		Links:   []AkamaiNetworkListLinks{},
// 		Message: "Element 1.2.3.4 successfully deleted from list",
// 	}

// 	received, resp, _ := client.NetworkLists.RemoveNetworkListItem("345_BOTLIST", "1.2.3.5/32")
// 	assert.Equal(t, wantResponseCode, resp.Response.StatusCode, "Response should be 202")
// 	assert.IsType(t, want, received, "Should be type of NetworkListResponse")
// 	assert.Equal(t, want, received, "Response should contain details of network lists created")
// }

// func TestActivateNetworkList(t *testing.T) {
// 	mux, server, client := setup()
// 	defer teardown(server)

// 	wantResponseCode := 200

// 	// Activate a network list
// 	activateListOpts := ActivateNetworkListOptions{
// 		SiebelTicketID:         "test-01",
// 		NotificationRecipients: []string{},
// 		Comments:               "activated by new API client",
// 	}

// 	mux.HandleFunc("/network-list/v1/network_lists/345_BOTLIST/activate", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(200)
// 		testMethod(t, r, "POST")
// 		fmt.Fprintf(w, `{
// 			"status": 200,
// 			"unique-id": "345_BOTLIST",
// 			"links": [],
// 			"sync-point": 1,
// 			"activation-status": "PENDING_ACTIVATION"
// 		 }`)
// 	})

// 	want := &NetworkListResponse{
// 		Status:           wantResponseCode,
// 		Links:            []AkamaiNetworkListLinks{},
// 		ActivationStatus: "PENDING_ACTIVATION",
// 		UniqueID:         "345_BOTLIST",
// 		SyncPoint:        1,
// 	}

// 	received, resp, _ := client.NetworkLists.ActivateNetworkList("345_BOTLIST", Staging, activateListOpts)
// 	assert.Equal(t, wantResponseCode, resp.Response.StatusCode, "Response should be 202")
// 	assert.IsType(t, want, received, "Should be type of NetworkListResponse")
// 	assert.Equal(t, want, received, "Response should contain details of network lists created")
// }
