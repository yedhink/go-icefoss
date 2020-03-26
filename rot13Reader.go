package main
import (
	"io"
	"os"
	"strings"
)

func rot13Enocoder(b byte) byte {
	if b >= 'a' && b <= 'z' {
		b = (b-'a'+13)%('z'-'a'+1) + 'a'
	} else if b >= 'A' && b <= 'Z' {
		b = (b-'A'+13)%('Z'-'A'+1) + 'A'
	}
	return b
}

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(bytes []byte) (int,error) {
	n,err := rot13.r.Read(bytes)
	for i:=0;i<n;i++{
		bytes[i] = rot13Enocoder(bytes[i])
	}
	return n,err
}

func main(){
	stream := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{stream}
	io.Copy(os.Stdout,&r)
}
