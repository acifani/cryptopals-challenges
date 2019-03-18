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
	"math/rand"
	"net/url"

	"github.com/acifani/cryptopals-challenges/set1"
)

func profileFor(email string) string {
	values := url.Values{}
	values.Set("email", email)
	values.Set("uid", string(rand.Intn(1000)))
	values.Set("role", "user")
	return values.Encode()
}

func cookieECBOracle() (encrypt func(email string) string, decrypt func(cookie string) url.Values) {
	key := make([]byte, 16)
	rand.Read(key)

	encrypt = func(email string) string {
		profile := profileFor(email)
		paddedProfile := PadRight([]byte(profile), 16)
		encrypted := EncryptAESinECB(paddedProfile, key)
		return string(encrypted)
	}

	decrypt = func(cookie string) url.Values {
		decrypted := set1.DecryptAESinECB([]byte(cookie), key)
		values, err := url.ParseQuery(string(decrypted))
		if err != nil {
			panic(err)
		}
		return values
	}

	return encrypt, decrypt
}

func BreakCookieECBORacle() {

}
