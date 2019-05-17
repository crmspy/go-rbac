package cgx
/*
this is general function collection
it's can use in every time you need
*/
import "crypto/sha512"
import "crypto/sha256"
import "time"
import "fmt"
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func CgxByteToString(text []byte) string {
    return fmt.Sprintf("%x", text)
}

//get offset and limit for paging
func Calcpage(page int)(int,int){
    page -= 1
    limit := 10
    offset := (page * limit) + 1
    return offset,limit
}

//generate hash
func CgxSha512(text string) string{
    h := sha512.New()
    h.Write([]byte(text))
    ret := h.Sum(nil)
    return CgxByteToString(ret)
}

//generate hash
func CgxSha256(text string) string{
    h := sha256.New()
    h.Write([]byte(text))
    ret := h.Sum(nil)
    return CgxByteToString(ret)
}

func CgxNow() time.Time {
    ret := time.Now()
    return ret
}

