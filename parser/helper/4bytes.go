package helper

const SignatureDBURL = "https://www.4byte.directory/api/v1/event-signatures/"

// SignatureResult represents the response from SignatureDBURL
type SignatureResult struct {
	Results []struct {
		HexSignature  string `json:"hex_signature"`
		TextSignature string `json:"text_signature"`
	} `json:"results"`
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}
