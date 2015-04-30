// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

package example2

import (
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestDownload validates the http Get function can download content and
// handles different status conditions properly.
func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{"http://www.goinggo.net/feeds/posts/default?alt=rss", http.StatusOK},
		{"http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen the URL is \"%s\" with status code \"%d\"", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err == nil {
					t.Log("\t\tShould be able to make the Get call.",
						succeed)
				} else {
					t.Fatal("\t\tShould be able to make the Get call.",
						failed, err)
				}

				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould receive a \"%d\" status code. %v",
						u.statusCode, succeed)
				} else {
					t.Errorf("\t\tShould receive a \"%d\" status code. %v %v",
						u.statusCode, failed, resp.StatusCode)
				}
			}
		}
	}
}