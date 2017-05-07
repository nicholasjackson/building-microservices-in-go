package asymmetric

import (
	"fmt"
	"testing"
)

const (
	smallMessage    = "Lorem ipsum dolor sit amet"
	largeMessage    = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
	smallCiphertext = "uAIwFleoavyAurRCOA+icpYYv709zNHfu9cXh1JMzhQ0gneWw+ncWYTiH2Z6FlhTqcOiMq+A1LtlVyP0bQo3PoMegCqi0gFdE4+oB6KLyCFvofMnzJ5vsQedx3ImipQqVmdr5h1MeqiS5EK0vfvdn3e1KszCktrK/aYWwBKql3yY22wPxKn3DvwzgxjUE/VuOYwO0o8sGzTbAxSrXTp5BBaCRvhyrSdQBg6/s6ozrsOfQMhjdlq68Gwdw1DBIxVopmeJYmrd8Aj4yrKTD1Eqt3fBJbHW3Qp6lc+iLRT2wmkP+GeQ7Y0sZjOQ8hHdGz1Hqour/CQ1cPMGWchHzg/k8w=="
	largeCipherkey  = "F+nHN7zeu796dO5XBfnZkNYBx3I+xUPH64hsGRpj6ED+71+J/3nqfxf1PeCcUA3EurtcFrQcD54VKssMiIkXck/iZrxjNbf5df/1zzEbu7mFpOoRb5Bg2MJliy12kYgL7lnS+qoYFCucCfGUViXZLk20M+3k4DY/pY/a0HfyQtYfemKMUKtbodHv2eZZLZ++gmpQrg28DmfkoNql0FHFrjsl25FRTkBlSVWKxnKYcXZXcNqhLe5HAxfUy1IgqUK+9ZfJfqNLADp/qO8gnKqmcD4TGD4awmTdpXUBlrFBTmDgAfyHAoV4PV3AF3Uwol7mRMLBp70VRfW9KenyfCRNXw=="
	largeCiphertext = "x8M/IgEEw62AdtKckUY7WH8l8yUxOV11BXkgjlqM9MX6bwCeHhl9dGsW2Hbvl4QmlsPM69u1VUfW6Jb9w3EdAWz6vbQBHpq9QI3i7XubogpMf3ykLWL15WyXFYpfgMrjeCYa1P45e4ExctkGlgXIuOtGaZNbCEIEygNhywUGQ1TPNKE9kkn+ZOfUDnGc4maEG56jS6bnAyZu0rxqvKKYHFdBFikuqzYMEgzKIw4np4mtAiKEEst1PvBRuUPxGvFPjAW4CBb7NHQJ8LG5v24V3oo3YgM0C2z5XBunp5yGeqQqsgz6Z/4NkzvgWFshjU4g6VOrj0ZEvuKRtaKYBl17c58jzc/GLyUam5fzQxfsLQiyJf/7TRtq3uXf0e6igjwuanh+aqg1dKAqgOnZluUizkIFor9Ip22koF69bIehNpj+Ktw6/Moz44DYdimDUAUs9FGa4b5WRAm5zxfxQwEt66hNflyrMzUAn7fuFyZPGwUNyl7712Hrh4A3Ene3+AFBSy8LysXz73izZHogDzUpkhxSi3wSdgA02kH/5gEDCMYUDjUSrWvlIhIKglkGGmSOVTBtwQTkYzavBt3GJITgKFr5gP2m/GrdLCvPJVpPrKQBJO8ZNQASTg=="
)

func TestEncryptSmallMessageWithPublicKey(t *testing.T) {
	ciphertext, err := EncryptMessageWithPublicKey(smallMessage)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(ciphertext)
}

func TestDecryptSmallMessageWithPrivateKey(t *testing.T) {
	plainData, err := DecryptMessageWithPrivateKey(smallCiphertext)
	if err != nil {
		t.Fatal(err)
	}

	if string(plainData) != smallMessage {
		t.Fatalf("Expecting %v, got %v", smallMessage, string(plainData))
	}
}

func TestEncryptLargeMessageWithPublicKey(t *testing.T) {
	ciphertext, key, err := EncryptLargeMessageWithPublicKey(largeMessage)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Key: ", key)
	fmt.Println("Cipher: ", ciphertext)
}

func TestDecryptLargeMessageWithPrivateKey(t *testing.T) {
	message, err := DecryptLargeMessageWithPrivateKey(largeCiphertext, largeCipherkey)
	if err != nil {
		t.Fatal(err)
	}

	if message != largeMessage {
		t.Fatalf("Expected %s, Got: %s", largeMessage, message)
	}
}
