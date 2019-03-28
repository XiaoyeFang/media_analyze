package storage

import "io"

type Storage interface {
	CopyFile(localPath, remotePath string) error
	CopyFileByUrl(srcLink, remotePath string) (string, error)
	CopyFileByReader(r io.Reader, size int64, fileType, remotePath string) (string, error)
	LsDir(path string) ([]string, error)
	CheckExist(key string) (string, bool)
	RemoveDir(path string) error
}
