// PKGPATH: gno.land/r/nft_test
package nft_test

import (
	"gno.land/p/testutils"
	"gno.land/r/nft"
)

func main() {
	addr1 := testutils.TestAddress("addr1")
	addr2 := testutils.TestAddress("addr2")
	grc721 := nft.GetGRC721()
	tid := grc721.Mint(addr1, "NFT#1")
	println(grc721.OwnerOf(tid))
	println(addr1)
	grc721.TransferFrom(addr1, addr2, tid)
}

// Error:
// unauthorized
