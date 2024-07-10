# DateTime-Client-Abdelrahman-Mahmoud

## Introduction

A client-server denotes a relationship between cooperating programs in an application, composed of clients initiating requests for services and servers providing that function or service. In this project we use the previously made date time server to get the current date and time.

## Setup

1. Clone the Repository to a directory of your choice.
2. Make sure you have go version 1.22.4 installed on your device
3. Run the server use as a reference https://github.com/codescalersinternships/DateTime-Server-Abdelrahman-Mahmoud/tree/development
4. Create demo.go file inside the working directory
5. import the package using 
   ```GO
	  import "github.com/codescalersinternships/DateTime-Client-Abdelrahman-Mahmoud"
   ```
6. Finish writing your desired code 
7. Open terminal
8. Build the project using
   ```console
   user@user-VirtualBox:~$ go build demo.go
   ```
9. Run the project using
   ```console
   user@user-VirtualBox:~$ ./demo
   ```

## Demo
- Code:
```GO
fmt.Println("Client created")

	myClient := client.NewClient()

	fmt.Println("HTTP server - GET /datetime  -> current date and time")
	returnedDateTime, err := myClient.GetHTTPDateTime(context.Background())

	if err != nil {
		log.Fatalf("error getting current date and time: %s", err)
	} else {
		fmt.Println(returnedDateTime)
	}

	fmt.Println("Gin server - GET /datetime  -> current date and time")
	returnedDateTime, err = myClient.GetGinDateTime(context.Background())

	if err != nil {
		log.Fatalf("error getting current date and time: %s", err)
	} else {
		fmt.Println(returnedDateTime)
	}
```

## Tests

=== RUN   TestOptionFunctions
=== RUN   TestOptionFunctions/happy_path_-_can_add_custom_URLS_using_option_function
=== RUN   TestOptionFunctions/happy_path_-_can_add_custom_port_numbers_using_option_function
=== RUN   TestOptionFunctions/happy_path_-_can_add_custom_client_using_option_function
=== RUN   TestOptionFunctions/happy_path_-_can_add_custom_client_using_option_function#01
--- PASS: TestOptionFunctions (0.00s)
    --- PASS: TestOptionFunctions/happy_path_-_can_add_custom_URLS_using_option_function (0.00s)
    --- PASS: TestOptionFunctions/happy_path_-_can_add_custom_port_numbers_using_option_function (0.00s)
    --- PASS: TestOptionFunctions/happy_path_-_can_add_custom_client_using_option_function (0.00s)
    --- PASS: TestOptionFunctions/happy_path_-_can_add_custom_client_using_option_function#01 (0.00s)
=== RUN   TestClientCanHitHTTPMockServer
=== RUN   TestClientCanHitHTTPMockServer/can_hit_the_mockserver_and_return_date_&_time
=== RUN   TestClientCanHitHTTPMockServer/can_handle_500_status_code
--- PASS: TestClientCanHitHTTPMockServer (0.01s)
    --- PASS: TestClientCanHitHTTPMockServer/can_hit_the_mockserver_and_return_date_&_time (0.01s)
    --- PASS: TestClientCanHitHTTPMockServer/can_handle_500_status_code (0.00s)
=== RUN   TestClientCanHitGinMockServer
=== RUN   TestClientCanHitGinMockServer/can_hit_the_mockserver_and_return_date_&_time
=== RUN   TestClientCanHitGinMockServer/can_handle_500_status_code
=== RUN   TestClientCanHitGinMockServer/can_handle_wrong_json_format
--- PASS: TestClientCanHitGinMockServer (0.01s)
    --- PASS: TestClientCanHitGinMockServer/can_hit_the_mockserver_and_return_date_&_time (0.00s)
    --- PASS: TestClientCanHitGinMockServer/can_handle_500_status_code (0.00s)
    --- PASS: TestClientCanHitGinMockServer/can_handle_wrong_json_format (0.00s)
PASS
ok      github.com/codescalersinternships/DateTime-Client-Abdelrahman-Mahmoud/client    0.032s
