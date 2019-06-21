package zipping

import (
	"archive/zip"
	"fmt"
	"strings"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/zippkg"
)

//AddImage adds an image (either from bytes or from disk) and adds it to the zip file.
func AddImage(z *zip.Writer, img common.ImageRef, imageNum int, dt unioffice.DocType) error {
	filename := unioffice.AbsoluteImageFilename(dt, imageNum, strings.ToLower(img.Format()))
	if img.Data() != nil && len(*img.Data()) > 0 {
		if err := zippkg.AddFileFromBytes(z, filename, *img.Data()); err != nil {
			return err
		}
	} else if img.Path() != "" {
		if err := zippkg.AddFileFromDisk(z, filename, img.Path()); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unsupported image source: %+v", img)
	}
	return nil
}
