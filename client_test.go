package andromeda

var testClient Client

func init() {
	testClient = NewClient("http://localhost:9711", "example-password")
}
