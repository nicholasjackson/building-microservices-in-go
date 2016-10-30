# Bench
[![GoDoc](https://godoc.org/github.com/nicholasjackson/bench?status.svg)](https://godoc.org/github.com/nicholasjackson/bench) [![CircleCI](https://circleci.com/gh/nicholasjackson/bench.svg?style=svg)](https://circleci.com/gh/nicholasjackson/bench)
[![Coverage Status](https://coveralls.io/repos/github/nicholasjackson/bench/badge.svg?branch=master)](https://coveralls.io/github/nicholasjackson/bench?branch=master)

Bench is a simple load testing application for Microservices, I built this as I wanted to be able run some basic performance tests against the microservice frameworks which I am evaluating for my book "Building Microservices In Go".  Whilst there are many options to be able to test HTTP based services I could find nothing which would serve the purpose of being able to benchmark an RPC server or another framework which did not use JSON or another text based object like XML over REST.  

With bench you write your request as a simple Go function and then by configuring the run options such as the number of threads, timeout and run time, Bench will execute the given function and collect the results.  These results can then be processed by one or more of the given output writers to be able to output simple tabular summaries or more detailed charts.

At present beyond the bounds of this simple project I have limited ambition to extend Bench and because it was originally written as a script, the complexity of which grew beyond my initial desire (classic software) the test coverage is not as great as I would like and the code itself is considered work in progress.  That said it works and it serves the purpose for which it was originally intended.

Because Bench executes a function rather it is possible to use bench in any situation where you would like to test any code in a concurrent manner.


# How Bench works
Bench will execute the function repeatedly for a given time period, every test execution runs in a separate go function so it is possible to execute multiple instances concurrently.  The number of threads which are executed at any one time is user specifiable and is limited to the resources on the machine running the test.  There is also a simple thread ramp up algorythim which will increase the number of threads in a linear manner over a given time period.  

# Example
The basic example which can be found in the ./example directory contains a test to examine the performance of the google.com home page, threads are limited to 10 and ramp up over a period of 10s.  Lets take a look at the code and then we can examine the options:

```go
func main() {

	fmt.Println("Benchmarking application")

	b := bench.New(10, 10*time.Second, 10*time.Second, 2*time.Second)
	b.AddOutput(11*time.Second, os.Stdout, output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	b.RunBenchmarks(AmazonRequest)
}
```

The main function is where we are setting up the test, we create a new instace of bench and pass it parameters for:
* max threads
* test duration
* thread ramp up
* request timeout

We then add an output function to the test, the first output we are adding is the tabular results output, all output functions have the same sinature which take:
* bucket interval - the bucket size which we would like to reduce and aggregate the captured results into;
* output writer - any struct which implements the io.Writer interface;
* output function - the function which will process the given output.

If we look at the first bucket interval we are specifying a size which is larger than the maximum time of the test, this will allow us to output a table which will aggregate the total results from the test.  Because we the stdout implements the io.Writer interface we can use this directly and the processed table will be written like the below example.

```
Start Time: 2016-09-20 19:34:17.008476638 +0000 UTC
Total Requests: 330
Avg Request Time: 0s
Total Success: 0
Total Timeouts: 0
Total Failures: 330
```

The second output we are adding will write 10 tables to a file "./output.txt" this gives more granular information with each outputed table representing one second of the total duration of the test.  For conveneience we are using the `util.NewFile` method, this will create a new file, overwriting any existing file returning an ioWriter.

The third output we are adding will plot the results on a chart like the one which can be seen below.

![](https://raw.githubusercontent.com/nicholasjackson/bench/master/example/output.png)

This chart is useful to see the response time of the given function against the current number of requests executed.  When load testing an API it is common to see the request time increase as the number of executed requests increases.

The final call we make is the `RunBenchmarks` function which will start the tests, we pass a function which contains the work we wish to test which has the signature `func() error`.  When writing a request you should only return an error when the request fails to execute, timeouts are measured by the RunBenchmarks function and should not be returned as an error from your request function.

```go
func AmazonRequest() error {

	resp, err := http.Get("http://www.amazon.co.uk/")
	defer resp.Body.Close()

	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return fmt.Errorf("Oops")
	}

	return nil
}
```

TODO:  
[ ] Update documentation  
[ ] Increase test coverage  
[ ] Write CSV data output  
