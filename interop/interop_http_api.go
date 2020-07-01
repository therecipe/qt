// +build http_interop

package interop

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {

	// into Local

	http.HandleFunc("/syncCallIntoLocal", func(w http.ResponseWriter, r *http.Request) {
		bb := new(bytes.Buffer)
		io.Copy(bb, r.Body)
		r.Body.Close()

		w.Write([]byte(_syncCallIntoLocal(bb.String())))
	})

	http.HandleFunc("/asyncCallIntoRemoteResponse", func(w http.ResponseWriter, r *http.Request) {
		bb := new(bytes.Buffer)
		io.Copy(bb, r.Body)
		r.Body.Close()

		_asyncCallIntoRemoteResponse(bb.String())
		w.Write(nil)
	})

	// into Remote

	AsyncCallIntoRemote = func(i string) {
		fmt.Fprintln(os.Stderr, "async:"+i)
	}

	//

	go http.ListenAndServe("127.0.0.1:8000", nil)
}
