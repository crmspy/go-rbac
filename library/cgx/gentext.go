package cgx
/*
this is general function collection
it's can use in every time you need
*/

import "math/rand"
import "time"

const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "1234567890"

const charsetint = "1234567890"
var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func CgxGenerateString(length int) string {
  return StringWithCharset(length, charset)
}

func CgxGenerateNumber(length int) string {
    return StringWithCharset(length, charsetint)
  }
