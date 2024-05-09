package s3

import (
	"context"
	"fmt"
	"os"

	archiver "github.com/mholt/archiver/v4"
)

func (s *s3Client) Zip(
	ctx context.Context,
	folderName string,
	fileNames map[string]string,
) (err error) {
	// zip folder
	tmpDir := os.TempDir()
	zipNameWithPath := fmt.Sprintf("%s/%s", tmpDir, folderName)
	// map files on disk to their paths in the archive
	files, err := archiver.FilesFromDisk(nil, fileNames)
	if err != nil {
		return err
	}

	// create the output file we'll write to
	archiverFile, err := os.Create(fmt.Sprintf("%s.zip", zipNameWithPath))
	if err != nil {
		return err
	}
	defer archiverFile.Close()

	// we can use the CompressedArchive type to gzip a tarball
	// (compression is not required; you could use Tar directly)
	//nolint:exhaustruct // archiver struct, not our struct
	format := archiver.CompressedArchive{
		Archival: archiver.Zip{},
	}

	// create the archive
	err = format.Archive(context.Background(), archiverFile, files)
	if err != nil {
		return err
	}

	return nil
}
