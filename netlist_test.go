package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	"github.com/stretchr/testify/assert"
)

// setup sets up a test HTTP server along with a edgegrid.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (*http.ServeMux, *httptest.Server, *edgegrid.Client) {
	// mux is the HTTP request multiplexer used with the test server.
	mux := http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// Provide struct details needed for apiClient init
	apiClientOpts := &edgegrid.ClientOptions{}
	apiClientOpts.ConfigPath = "/Users/rpieniazek/.edgerc"
	apiClientOpts.ConfigSection = "jmb"
	// apiClientOpts.DebugLevel = "debug"
	apiClientOpts.AccountSwitchKey = ""

	// Initialize client for tests with init options
	apiClient, err := common.EdgeClientInit(apiClientOpts)
	if err != nil {
		log.Fatalln(err)
	}
	apiClient.SetBaseURL(server.URL, true)

	return mux, server, apiClient
}

// teardown closes the test HTTP server.
func teardown(server *httptest.Server) {
	server.Close()
}

func testURL(t *testing.T, r *http.Request, want string) {
	if got := r.RequestURI; got != want {
		t.Errorf("Request url: %+v, want %s", got, want)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %s, want %s", got, want)
	}
}

func TestListNetworkLists(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	listNetListOptsv2 := edgegrid.ListNetworkListsOptionsv2{}
	listNetListOptsv2.Search = "" // Since we are listing all we do not filter results

	mux.HandleFunc("/network-list/v2/network-lists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		body, _ := ioutil.ReadAll(r.Body)
		assert.Equal(t, "", string(body), "Request body should be empty")

		fmt.Fprintf(w, `{
			"network_lists": [
				{
					"name": "test",
					"type": "IP",
					"links": [
						{
							"rel": "get 123_TEST",
							"href": "/network-list/v1/network_lists/123_TEST"
						}
					],
					"unique-id": "123_TEST",
					"list": [
						"1.2.3.4",
						"5.6.7.8"
					],
					"sync-point": 1,
					"numEntries": 2
				}
			]
		}`)
	})

	_, resp, netlistErr := client.NetworkListsv2.ListNetworkLists(listNetListOptsv2)
	if netlistErr != nil {
		fmt.Println(netlistErr)
	}

	assert.Equal(t, 200, resp.Response.StatusCode, "Response should be 200 OK")
}

func TestGetNetworkListById(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	listNetListOptsv2 := edgegrid.ListNetworkListsOptionsv2{}
	listNetListOptsv2.Search = "" // Since we are listing all we do not filter results

	mux.HandleFunc("/network-list/v2/network-lists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		// keys, ok := r.URL.Query()["id"]

		// if !ok || len(keys[0]) < 1 {
		// 	log.Println("Url Param 'key' is missing")
		// 	return
		// }

		body, _ := ioutil.ReadAll(r.Body)
		assert.Equal(t, "", string(body), "Request body should be empty")

		fmt.Fprintf(w, `{
			"network_lists": [
				{
					"name": "test",
					"type": "IP",
					"links": [
						{
							"rel": "get 123_TEST",
							"href": "/network-list/v1/network_lists/123_TEST"
						}
					],
					"unique-id": "123_TEST",
					"list": [
						"1.2.3.4",
						"5.6.7.8"
					],
					"sync-point": 1,
					"numEntries": 2
				}
			]
		}`)
	})

	// netList, _, netlistErr := apiClient.NetworkListsv2.
	_, _, netlistErr := client.NetworkListsv2.GetNetworkList("123_TEST", listNetListOptsv2)
	if netlistErr != nil {
		fmt.Println(netlistErr)
	}

	// assert.Equal(t, 200, resp.Response.StatusCode, "Response should be 200 OK")
}

// func TestGetNetworkList(t *testing.T) {
// 	mux, server, client := setup()
// 	defer teardown(server)

// 	// Set options for working with network lists
// 	opt := ListNetworkListsOptions{
// 		TypeOflist:        "IP",
// 		Extended:          true,
// 		IncludeDeprecated: false,
// 		IncludeElements:   false,
// 	}

// 	mux.HandleFunc("/network-list/v1/network_lists/345_BOTLIST", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		fmt.Fprintf(w, `{
// 			"name": "single",
// 			"type": "IP",
// 			"unique-id": "345_BOTLIST",
// 			"links": [
// 				{
// 				   "rel": "get 345_BOTLIST",
// 				   "href": "/network-list/v1/network_lists/345_BOTLIST"
// 				}
// 			],
// 			"list": [
// 			   "192.168.0.1",
// 			   "192.168.0.2",
// 			   "192.168.0.3",
// 			   "198.168.0.4",
// 			   "198.168.0.5",
// 			   "198.168.0.6"
// 			],
// 			"sync-point": 1,
// 			"numEntries": 6
// 		 }`)
// 	})

// 	received, resp, _ := client.NetworkLists.GetNetworkList("345_BOTLIST", opt)
// 	assert.Equal(t, 200, resp.Response.StatusCode, "Response should be 200 OK")

// 	want := &AkamaiNetworkList{
// 		Name:     "single",
// 		Type:     "IP",
// 		UniqueID: "345_BOTLIST",
// 		List: []string{
// 			"192.168.0.1",
// 			"192.168.0.2",
// 			"192.168.0.3",
// 			"198.168.0.4",
// 			"198.168.0.5",
// 			"198.168.0.6",
// 		},
// 		Links: []AkamaiNetworkListLinks{
// 			AkamaiNetworkListLinks{
// 				Rel:  "get 345_BOTLIST",
// 				Href: "/network-list/v1/network_lists/345_BOTLIST",
// 			},
// 		},
// 		SyncPoint:  1,
// 		NumEntries: 6,
// 	}
// 	assert.Equal(t, want, received, "Response should contain valid array of network lists")

// }

// func TestCreateNetworkList(t *testing.T) {
// 	mux, server, client := setup()
// 	defer teardown(server)

// 	opt := CreateNetworkListOptions{
// 		Name:        "345_BOTLIST",
// 		Type:        "IP",
// 		Description: "created-by-test",
// 	}

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
