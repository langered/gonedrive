package browser

import gobrowser "github.com/pkg/browser"

//Browser is an interface representation of the used functionality of pkg/browser
type Browser interface {
	OpenURL(url string) error
}

//GonedriveBrowser is the own implementation of pkg/browser
type GonedriveBrowser struct{}

//OpenURL wraps the OpenURL mehtod of pkg/browser
func (browser GonedriveBrowser) OpenURL(url string) error {
	return gobrowser.OpenURL(url)
}
