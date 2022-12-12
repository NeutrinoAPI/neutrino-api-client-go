package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
	"strings"
)

func main() {
	params := make(url.Values, 3)

	// The character to use to censor out the bad words found
	params.Add("censor-character", "")

	// Which catalog of bad words to use, we currently maintain two bad word catalogs:
	// • strict - the largest database of bad words which includes profanity, obscenity, sexual, rude,
	//   cuss, dirty, swear and objectionable words and phrases. This catalog is suitable for
	//   environments of all ages including educational or children's content
	// • obscene - like the strict catalog but does not include any mild profanities, idiomatic
	//   phrases or words which are considered formal terminology. This catalog is suitable for adult
	//   environments where certain types of bad words are considered OK
	params.Add("catalog", "strict")

	// The content to scan. This can be either a URL to load from, a file upload or an HTML content
	// string
	params.Add("content", "https://en.wikipedia.org/wiki/Profanity")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.BadWordFilter(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// An array of the bad words found
		badWordsList := strings.Fields(fmt.Sprint(data["bad-words-list"]))
		fmt.Printf("bad-words-list: %s\n", strings.Join(badWordsList, ", "))

		// Total number of bad words detected
		fmt.Printf("bad-words-total: %.f\n", data["bad-words-total"])

		// The censored content (only set if censor-character has been set)
		fmt.Printf("censored-content: \"%s\"\n", data["censored-content"])

		// Does the text contain bad words
		fmt.Printf("is-bad: %t\n", data["is-bad"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
