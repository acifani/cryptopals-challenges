/*
ECB cut-and-paste

Write a k=v parsing routine, as if for a structured cookie. The routine should take:

foo=bar&baz=qux&zap=zazzle

... and produce:

{
  foo: 'bar',
  baz: 'qux',
  zap: 'zazzle'
}

(you know, the object; I don't care if you convert it to JSON).
Now write a function that encodes a user profile in that format, given an email address.
You should have something like:

profile_for("foo@bar.com")

... and it should produce:

{
  email: 'foo@bar.com',
  uid: 10,
  role: 'user'
}
... encoded as:

email=foo@bar.com&uid=10&role=user

Your "profile_for" function should not allow encoding metacharacters (& and =).
Eat them, quote them, whatever you want to do, but don't let people
set their email address to "foo@bar.com&role=admin".

Now, two more easy functions. Generate a random AES key, then:
A. Encrypt the encoded user profile under the key; "provide" that to the "attacker".
B. Decrypt the encoded user profile and parse it.

Using only the user input to profile_for() (as an oracle to generate "valid" ciphertexts)
and the ciphertexts themselves, make a role=admin profile.
*/

package set2

import (
	"log"
	"math/rand"
	"net/url"
	"strconv"

	"github.com/acifani/cryptopals-challenges/set1"
)

func profileFor(email string) string {
	values := url.Values{}
	values.Set("email", email)
	values.Set("uid", strconv.Itoa(rand.Intn(99)))
	values.Set("role", "user")
	return values.Encode()
}

func cookieECBOracle() (encrypt func(email []byte) []byte, decrypt func(cookie []byte) url.Values) {
	key := make([]byte, 16)
	rand.Read(key)

	encrypt = func(email []byte) []byte {
		profile := profileFor(string(email))
		paddedProfile := PadRight([]byte(profile), 16)
		encrypted := EncryptAESinECB(paddedProfile, key)
		return encrypted
	}

	decrypt = func(cookie []byte) url.Values {
		decrypted := set1.DecryptAESinECB([]byte(cookie), key)
		log.Printf("\n%s", decrypted)
		values, err := url.ParseQuery(string(decrypted))
		if err != nil {
			panic(err)
		}
		return values
	}

	return encrypt, decrypt
}

// BreakECBCookie given an encrypt cookie oracle that generates
// cookie profile in the format: `email=foo@bar.me&role=user&uid=99<padding>`,
// this function will forge an encrypted admin cookie.
func BreakECBCookie(encrypt func(email []byte) []byte) []byte {
	blockSize := 16
	// cookie = "email=foo@bar.me | &role=user&uid=9 | 9<padding>"

	// Block 1: email=foo@bar.me
	block1 := encrypt([]byte("foo@bar.me"))[:blockSize]
	// Block 2: 1234567890&role=
	block2 := encrypt([]byte("12345678901234567890"))[blockSize : blockSize*2]
	// Block 3: admin&role=user&uid=99<padding>
	block3 := encrypt([]byte("1234567890admin"))[blockSize:]

	// We have to concat 2 different 'role' keys because of the way
	// Go's url module orders the value keys in the dictionary.
	// It would've been much easier otherwise, but let's pretend we can't.
	// email=foo@bar.me1234567890&role=admin&role=user&uid=99<padding>
	return append(append(block1, block2...), block3...)
}
