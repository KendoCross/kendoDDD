package files

import "github.com/KendoCross/kendoDDD/interfaces"

type FileInfos struct {
	interfaces.FileInfo
	FileBody []byte
}
