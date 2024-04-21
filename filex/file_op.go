package filex

import (
	"bufio"
	"fmt"

	"github.com/fasthey/go-utils/cmdx"

	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

type FileOp struct {
	Fs afero.Fs
}

func NewFileOp() FileOp {
	return FileOp{
		Fs: afero.NewOsFs(),
	}
}

func (f FileOp) OpenFile(dst string) (fs.File, error) {
	return f.Fs.Open(dst)
}

func (f FileOp) GetContent(dst string) ([]byte, error) {
	afs := &afero.Afero{Fs: f.Fs}
	cByte, err := afs.ReadFile(dst)
	if err != nil {
		return nil, err
	}
	return cByte, nil
}

func (f FileOp) CreateDir(dst string, mode fs.FileMode) error {
	return f.Fs.MkdirAll(dst, mode)
}

func (f FileOp) CreateFile(dst string) error {
	if _, err := f.Fs.Create(dst); err != nil {
		return err
	}
	return nil
}

func (f FileOp) LinkFile(source string, dst string, isSymlink bool) error {
	if isSymlink {
		osFs := afero.OsFs{}
		return osFs.SymlinkIfPossible(source, dst)
	} else {
		return os.Link(source, dst)
	}
}

func (f FileOp) DeleteDir(dst string) error {
	return f.Fs.RemoveAll(dst)
}

func (f FileOp) Stat(dst string) bool {
	info, _ := f.Fs.Stat(dst)
	return info != nil
}

func (f FileOp) DeleteFile(dst string) error {
	return f.Fs.Remove(dst)
}

func (f FileOp) Delete(dst string) error {
	return os.RemoveAll(dst)
}

func (f FileOp) RmRf(dst string) error {
	return cmdx.ExecCmd(fmt.Sprintf("rm -rf %s", dst))
}

func (f FileOp) WriteFile(dst string, in io.Reader, mode fs.FileMode) error {
	file, err := f.Fs.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = io.Copy(file, in); err != nil {
		return err
	}

	if _, err = file.Stat(); err != nil {
		return err
	}
	return nil
}

func (f FileOp) SaveFile(dst string, content string, mode fs.FileMode) error {
	if !f.Stat(path.Dir(dst)) {
		_ = f.CreateDir(path.Dir(dst), mode.Perm())
	}
	file, err := f.Fs.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	_, _ = write.WriteString(content)
	write.Flush()
	return nil
}

func (f FileOp) Chmod(dst string, mode fs.FileMode) error {
	return f.Fs.Chmod(dst, mode)
}

func (f FileOp) Chown(dst string, uid int, gid int) error {
	return f.Fs.Chown(dst, uid, gid)
}

func (f FileOp) ChownR(dst string, uid string, gid string, sub bool) error {
	cmdStr := fmt.Sprintf(`chown %s:%s "%s"`, uid, gid, dst)
	if sub {
		cmdStr = fmt.Sprintf(`chown -R %s:%s "%s"`, uid, gid, dst)
	}
	if cmdx.HasNoPasswordSudo() {
		cmdStr = fmt.Sprintf("sudo %s", cmdStr)
	}
	if msg, err := cmdx.ExecWithTimeOut(cmdStr, 2*time.Second); err != nil {
		if msg != "" {
			return errors.New(msg)
		}
		return err
	}
	return nil
}

func (f FileOp) ChmodR(dst string, mode int64, sub bool) error {
	cmdStr := fmt.Sprintf(`chmod %v "%s"`, fmt.Sprintf("%04o", mode), dst)
	if sub {
		cmdStr = fmt.Sprintf(`chmod -R %v "%s"`, fmt.Sprintf("%04o", mode), dst)
	}
	if cmd.HasNoPasswordSudo() {
		cmdStr = fmt.Sprintf("sudo %s", cmdStr)
	}
	if msg, err := cmd.ExecWithTimeOut(cmdStr, 2*time.Second); err != nil {
		if msg != "" {
			return errors.New(msg)
		}
		return err
	}
	return nil
}

func (f FileOp) Rename(oldName string, newName string) error {
	return f.Fs.Rename(oldName, newName)
}

type WriteCounter struct {
	Total   uint64
	Written uint64
	Key     string
	Name    string
}

type Process struct {
	Total   uint64  `json:"total"`
	Written uint64  `json:"written"`
	Percent float64 `json:"percent"`
	Name    string  `json:"name"`
}

func (f FileOp) Cut(oldPaths []string, dst, name string, cover bool) error {
	for _, p := range oldPaths {
		var dstPath string
		if name != "" {
			dstPath = filepath.Join(dst, name)
			if f.Stat(dstPath) {
				dstPath = dst
			}
		} else {
			base := filepath.Base(p)
			dstPath = filepath.Join(dst, base)
		}
		coverFlag := ""
		if cover {
			coverFlag = "-f"
		}

		cmdStr := fmt.Sprintf(`mv %s "%s" "%s"`, coverFlag, p, dstPath)
		if err := cmd.ExecCmd(cmdStr); err != nil {
			return err
		}
	}
	return nil
}

func (f FileOp) Mv(oldPath, dstPath string) error {
	cmdStr := fmt.Sprintf("mv %s  %s", oldPath, dstPath)
	if err := cmdx.ExecCmd(cmdStr); err != nil {
		return err
	}
	return nil
}

func (f FileOp) Copy(src, dst string) error {
	if src = path.Clean("/" + src); src == "" {
		return os.ErrNotExist
	}
	if dst = path.Clean("/" + dst); dst == "" {
		return os.ErrNotExist
	}
	if src == "/" || dst == "/" {
		return os.ErrInvalid
	}
	if dst == src {
		return os.ErrInvalid
	}
	info, err := f.Fs.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return f.CopyDir(src, dst)
	}
	return f.CopyFile(src, dst)
}

func (f FileOp) CopyAndReName(src, dst, name string, cover bool) error {
	if src = path.Clean("/" + src); src == "" {
		return os.ErrNotExist
	}
	if dst = path.Clean("/" + dst); dst == "" {
		return os.ErrNotExist
	}
	if src == "/" || dst == "/" {
		return os.ErrInvalid
	}
	if dst == src {
		return os.ErrInvalid
	}

	srcInfo, err := f.Fs.Stat(src)
	if err != nil {
		return err
	}

	if srcInfo.IsDir() {
		dstPath := dst
		if name != "" && !cover {
			dstPath = filepath.Join(dst, name)
		}
		return cmdx.ExecCmd(fmt.Sprintf(`cp -rf "%s" "%s"`, src, dstPath))
	} else {
		dstPath := filepath.Join(dst, name)
		if cover {
			dstPath = dst
		}
		return cmdx.ExecCmd(fmt.Sprintf(`cp -f "%s" "%s"`, src, dstPath))
	}
}

func (f FileOp) CopyDir(src, dst string) error {
	srcInfo, err := f.Fs.Stat(src)
	if err != nil {
		return err
	}
	dstDir := filepath.Join(dst, srcInfo.Name())
	if err = f.Fs.MkdirAll(dstDir, srcInfo.Mode()); err != nil {
		return err
	}
	return cmdx.ExecCmd(fmt.Sprintf(`cp -rf "%s" "%s"`, src, dst+"/"))
}

func (f FileOp) CopyFile(src, dst string) error {
	dst = filepath.Clean(dst) + string(filepath.Separator)
	return cmdx.ExecCmd(fmt.Sprintf(`cp -f "%s" "%s"`, src, dst+"/"))
}

func (f FileOp) GetDirSize(path string) (float64, error) {
	var m sync.Map
	var wg sync.WaitGroup

	wg.Add(1)
	go ScanDir(f.Fs, path, &m, &wg)
	wg.Wait()

	var dirSize float64
	m.Range(func(k, v interface{}) bool {
		dirSize = dirSize + v.(float64)
		return true
	})

	return dirSize, nil
}

func isIgnoreFile(name string) bool {
	return strings.HasPrefix(name, "__MACOSX") || strings.HasSuffix(name, ".DS_Store") || strings.HasPrefix(name, "._")
}

func decodeGBK(input string) (string, error) {
	decoder := simplifiedchinese.GBK.NewDecoder()
	decoded, _, err := transform.String(decoder, input)
	if err != nil {
		return "", err
	}
	return decoded, nil
}

func (f FileOp) Backup(srcFile string) (string, error) {
	backupPath := srcFile + "_bak"
	info, _ := f.Fs.Stat(backupPath)
	if info != nil {
		if info.IsDir() {
			_ = f.DeleteDir(backupPath)
		} else {
			_ = f.DeleteFile(backupPath)
		}
	}
	if err := f.Rename(srcFile, backupPath); err != nil {
		return backupPath, err
	}

	return backupPath, nil
}

func (f FileOp) CopyAndBackup(src string) (string, error) {
	backupPath := src + "_bak"
	info, _ := f.Fs.Stat(backupPath)
	if info != nil {
		if info.IsDir() {
			_ = f.DeleteDir(backupPath)
		} else {
			_ = f.DeleteFile(backupPath)
		}
	}
	_ = f.CreateDir(backupPath, 0755)
	if err := f.Copy(src, backupPath); err != nil {
		return backupPath, err
	}
	return backupPath, nil
}

func ScanDir(fs afero.Fs, path string, dirMap *sync.Map, wg *sync.WaitGroup) {
	afs := &afero.Afero{Fs: fs}
	files, _ := afs.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			wg.Add(1)
			go ScanDir(fs, filepath.Join(path, f.Name()), dirMap, wg)
		} else {
			if f.Size() > 0 {
				dirMap.Store(filepath.Join(path, f.Name()), float64(f.Size()))
			}
		}
	}
	defer wg.Done()
}
