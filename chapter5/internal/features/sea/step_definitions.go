package sea

import . "github.com/gucumber/gucumber"

func init() {
	Given(`^I have no search criteria$`, func() {
		T.Skip() // pending
	})

	When(`^I call the search endpoint$`, func() {
		T.Skip() // pending
	})

	Then(`^I should receive a bad request message$`, func() {
		T.Skip() // pending
	})

	Given(`^I have a valid search criteria$`, func() {
		T.Skip() // pending
	})

	Then(`^I should receive a list of kittens$`, func() {
		T.Skip() // pending
	})
}
