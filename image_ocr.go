// +build ocr

package docconv

import (
	"fmt"
	"io"
	"sync"

	"github.com/otiai10/gosseract"
)

var langs = struct {
	sync.RWMutex
	lang string
}{lang: "eng"}

// ConvertImage converts images to text.
// Requires gosseract.
func ConvertImage(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	langs.RLock()
	c := gosseract.NewClient()
	c.SetLanguage(langs.lang)
	c.SetImage(f.Name())
	body, err := c.Text()
	langs.RUnlock()

	return body, nil, err
}

// SetImageLanguages sets the languages parameter passed to gosseract.
func SetImageLanguages(l string) {
	langs.Lock()
	langs.lang = l
	langs.Unlock()
}
