//go:build unit

package utils_test

import (
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"

	"encore.app/utils"
)

func TestDecodeRSAPrivateKey_GivenValidBase64_ReturnsPrivateKey(t *testing.T) {
	encodedKey := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBdTRqVW1hdDhzeFVxMHlia29YMzBJM0haN1JyZHBUdjZiS0VCblMvbHFQd3R0R1VXClNWUlQxazRnTzhTdkNCdzFUZnhYQUNvZVZLdkprRWMwM2hqa09yekhIWWo4NVpPZGFUWUEzVE8zdzJvZTIvWjEKb1ZwdUFIMVozdFczUXdlbVRPZzc4T0YvMVp5a0RMV0NpWFpFenVhVGRwQ0JQeHBmcmltU290R3g3TDgvSloySwpkZUxPcEVMeis1ZDRza2VZZERGZXoyZVViaGZyaFNUUkp3NGFFRGV3Zm53QU5BdFdYK1NRYXZxK3ppaVQrTVRtClVWNmhXTGZlU2ZZUWQ3elFNZXNlVWc1STM0cHdyaHJOOE50eTV6dWZqcTJZcTBpZUtldnlUTTBIaEhIVmRZV2wKZHBmU3RaWC90MFJIdnQvU1l0N0VEbnEvc0M3SldwKzN1TWM1N1FJREFRQUJBb0lCQUF2TjN6MUpwelhwMlZkawowQ2oreHgwV0NsYkdML2tQR2UwTmcyUlVBcVkzaUFtZC8vV2ZWUTd5YUUxVmQ0T21BWkJHUW1tcHZRZ0NXaWtoCk9rbUNrcGNaNW9TaGswaEF4RHlUb0NGTVRlbXQwU3h5U253eURORFl2YmZzM0RZS3ZMa3U0d1dxQkFSbVJ0M2YKN1BhMGJEeXVhYzRoa1drWjdXSXhPV0FvblMrZXpGUHNmNjI3VUIvLzZ0L2xsMW4xSFZQR0Y5MFZkR0ZlYythNQpsZDFkdzRUTTllaFN3Z1MzeGxHdEdWYWhXME4rWWNqWW1iTE5tZUpYTTFlYk02UFNCRUVQSzFLbGM0d2YxY1hTCkhHN1pHRzB6V0N5TWJXU3E3OTNyUXp3TGIybTE5eXhUVmRwRVFENWpzWStyN3hnWTFvQVdLZ3hnQit6ajQwTVQKbmo2MHdMY0NnWUVBOTI0b1RLbHQwWXM5V1JTb2JtWHdZcDhUS0hIL1ZWb1RDSEtrM2ZXMGhtcmlMRnlVdFF5WApSQ2hrekt6RHArampGMmRBMWNSWEJ2ZlpmMElkeTh3VDBEbEVaMmQ5RElNdnBWRHplTVVUdHNNcWZQSVRiMk9NCjlWUTZoazZ2RGtQQndUN1hxOWhyV2F6L01KK3hJZFVJdmVyZWFoL3grOFlFUUNuckJKNTlvQXNDZ1lFQXdnZWIKUHVGMjFBbUpKOFp3S0VvNDh4czFEVEx6TjZjTGlwNDhHc1NsclcxT3BWT0UzbkN2T0hmanNxeFVrdGhpcGFCYQoydldaaFdLOFI1V3RCT2VML2xKeGx2cG4rYzhneVd1Sm96NGQyc0l2T3hjR3RGaTB4TmViMjRNOTZsN1FjR256CkM0YVFoU2crSEQwWmZNV3N3cmprd2hLNjJTUzZ4U2xKZVd1VmNPY0NnWUVBc1pEa2F0L01aK3k5QWdqSk56RUEKNmtXdXdmL05kYnRHblhGb3o4RmM1SG9nQmlZS3NrTnU2d0x5RTlDSU9SL3dtUU1JakdCeTlCVTZpUkV4d09lYgo4SHFvbEd5NFdScHNQdjl4T3VKejIxMEVTTmFSREFjdFNZSkEzZEhwM3hyUkdaVzN3U04zSnN2Mzl6VWpNVnhwCkFhSmtLMVR4bS9Yblp3U3VISXFCTFRFQ2dZQnJsM1l4dnlwNlY3TlRlQWszZmpqb0xiMFVUWjFxemRscmkyVCsKM3U2VUpabkh1WEZqTzQ4ZU8vVHFYZjhqMHBPWkRqdHpVVjlKQk1BczNjV3NnNDB3Z0p6MlFIS3BwbjZpMGx2bgpSb21kaGJKVkRPYXQyTWFjcElhTGlkSXFoVnNHQ3VvNENPMVl1VUQvdmEzRmI0UG1Fa3JmaHkxUFBidkhtcVpnCml4UzZId0tCZ0RxbjFjMVhnZUhvQ1RVNFFHTTRjWEJUdHFXaFJEcmtVUTRwQXcwR3hNT1k5VFU4c2oyOThXY0sKczZCSmF5L0dlTHlXN0REMnpWZGJ1UG1qRERjQldFVVNSS2hSbGp3eXc2SVdRU2xZRGpPcGNvdHAxdkRBQVRPNgpzK2hRWUpuYk4xdXV6MlMxZTZaUThSK1BhUlc5MXZkUjdWMEswd0V4SkhobnFTZHJXZXhRCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="

	privateKey, err := utils.DecodeRSAPrivateKey(encodedKey)

	assert.NoError(t, err, "Expected no error when decoding a valid RSA private key")
	assert.IsType(t, &rsa.PrivateKey{}, privateKey, "Expected a valid RSA private key")
}

func TestDecodeRSAPublicKey_GivenValidBase64_ReturnsPublicKey(t *testing.T) {
	encodedKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF1NGpVbWF0OHN4VXEweWJrb1gzMApJM0haN1JyZHBUdjZiS0VCblMvbHFQd3R0R1VXU1ZSVDFrNGdPOFN2Q0J3MVRmeFhBQ29lVkt2SmtFYzAzaGprCk9yekhIWWo4NVpPZGFUWUEzVE8zdzJvZTIvWjFvVnB1QUgxWjN0VzNRd2VtVE9nNzhPRi8xWnlrRExXQ2lYWkUKenVhVGRwQ0JQeHBmcmltU290R3g3TDgvSloyS2RlTE9wRUx6KzVkNHNrZVlkREZlejJlVWJoZnJoU1RSSnc0YQpFRGV3Zm53QU5BdFdYK1NRYXZxK3ppaVQrTVRtVVY2aFdMZmVTZllRZDd6UU1lc2VVZzVJMzRwd3Jock44TnR5CjV6dWZqcTJZcTBpZUtldnlUTTBIaEhIVmRZV2xkcGZTdFpYL3QwUkh2dC9TWXQ3RURucS9zQzdKV3ArM3VNYzUKN1FJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="

	publicKey, err := utils.DecodeRSAPublicKey(encodedKey)

	assert.NoError(t, err, "Expected no error when decoding a valid RSA public key")
	assert.IsType(t, &rsa.PublicKey{}, publicKey, "Expected a valid RSA public key")
}

func TestGenerateRandomBytes(t *testing.T) {
	length := 16
	bytes, err := utils.GenerateRandomBytes(length)

	assert.NoError(t, err, "Expected no error when generating random bytes")
	assert.Equal(t, length, len(bytes), "Expected byte slice of length %d", length)
}

func TestMustGenerateRandomBytes(t *testing.T) {
	length := 16
	bytes := utils.MustGenerateRandomBytes(length)

	assert.Equal(t, length, len(bytes), "Expected byte slice of length %d", length)
}

func TestGenerateRandomString(t *testing.T) {
	length := 32                // 256 bits
	expectedEncodedLength := 44 // 32 bytes encoded in base64 results in 44 characters

	str, err := utils.GenerateRandomString(length)

	assert.NoError(t, err, "Expected no error when generating random string")
	assert.Equal(t, expectedEncodedLength, len(str), "Expected encoded string length %d", expectedEncodedLength)
}

func TestMustGenerateRandomString(t *testing.T) {
	length := 32                // 256 bits
	expectedEncodedLength := 44 // 32 bytes encoded in base64 results in 44 characters

	str := utils.MustGenerateRandomString(length)

	assert.Equal(t, expectedEncodedLength, len(str), "Expected encoded string length %d", expectedEncodedLength)
}
