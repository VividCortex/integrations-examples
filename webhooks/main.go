package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// ReceiveHook listens for a POST request from VC and does something with it
//
// POST /hook
func ReceiveHook(secret string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Printf("ioutil read error: %v", err)
		}

		// Compute signature with our known Secret Token
		// https://docs.vividcortex.com/how-to-use-vividcortex/integrations/#generic-webhook
		h := sha1.New()
		io.WriteString(h, string(b)+secret)
		computedSignature := hex.EncodeToString(h.Sum(nil))

		signature := r.Header.Get("X-VividCortex-Signature")

		if signature != computedSignature {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Printf("signature didn't match: %v vs %v\n", signature, computedSignature)
		}

		/*

			Do Something

		*/
	})
}

func main() {
	secretToken := "ansecretansecretansecretansecret"

	http.HandleFunc("/hook", ReceiveHook(secretToken))
	log.Fatal(http.ListenAndServe(":1337", nil))
}
