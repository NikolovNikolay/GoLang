// [_SHA1 hashes_](http://en.wikipedia.org/wiki/SHA-1) are
// frequently used to compute short identities for binary
// or text blobs. For example, the [git revision control
// system](http://git-scm.com/) uses SHA1s extensively to
// identify versioned files and directories. Here's how to
// compute SHA1 hashes in Go.

package main

import "crypto/sha1"
import "crypto/md5"
import "fmt"

func main() {
	s := "some H@sh this will be"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("SHA1: %x\n", bs)

	md := md5.New()
	md.Write([]byte(s))
	mds := md.Sum(nil)

	fmt.Printf("MD5: %x\n", mds)
}
