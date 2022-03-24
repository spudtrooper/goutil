package fs

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type FindFilesFilterFn func(path string, info os.FileInfo) (bool, error)

type FindFilesResult struct {
	Path string
	Info os.FileInfo
}

func FindFiles(dir string, fOpts ...FindFilesOption) (chan FindFilesResult, chan error, error) {
	opts := MakeFindFilesOptions(fOpts...)
	filter := opts.Filter()
	if filter == nil {
		if opts.Keep() != nil {
			filter = func(path string, info os.FileInfo) (bool, error) {
				res := opts.Keep().MatchString(path)
				return res, nil
			}
		}
	}
	if filter == nil {
		filter = func(path string, info os.FileInfo) (bool, error) {
			return true, nil
		}
	}
	if opts.Recursive() {
		return findFilesRecursive(dir, filter)
	}
	return findFilesFlat(dir, filter)
}

func findFilesFlat(dir string, filter FindFilesFilterFn) (chan FindFilesResult, chan error, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, nil, err
	}
	res, errs := make(chan FindFilesResult), make(chan error)
	go func() {
		for _, f := range files {
			p := path.Join(dir, f.Name())
			ok, err := filter(p, f)
			if err != nil {
				errs <- err
				continue
			}
			if !ok {
				continue
			}
			res <- FindFilesResult{
				Path: p,
				Info: f,
			}
		}
		close(errs)
		close(res)
	}()
	return res, errs, nil
}

func findFilesRecursive(dir string, filter FindFilesFilterFn) (chan FindFilesResult, chan error, error) {
	res, errs := make(chan FindFilesResult), make(chan error)
	go func() {
		if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filter != nil {
				ok, err := filter(path, info)
				if err != nil {
					errs <- err
					return err
				}
				if !ok {
					return nil
				}
			}
			res <- FindFilesResult{
				Path: path,
				Info: info,
			}
			return nil
		}); err != nil {
			errs <- err
		}
		close(errs)
		close(res)
	}()
	return res, errs, nil
}
