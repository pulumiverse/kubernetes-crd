package pkg

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func GenerateSDKs(name string, crd CRDDefinition) error {
	// Clean up the folder for the CRD and recreate it
	if err := os.RemoveAll(name); err != nil {
		return fmt.Errorf("failed to clean up folder %s: %v", name, err)
	}
	if err := os.Mkdir(name, 0755); err != nil {
		return err
	}

	err := generateFor("dotnet", name, crd)
	if err != nil {
		return err
	}

	err = generateFor("go", name, crd)
	if err != nil {
		return err
	}

	err = generateFor("nodejs", name, crd)
	if err != nil {
		return err
	}

	err = generateFor("python", name, crd)
	if err != nil {
		return err
	}

	return nil
}

func generateFor(language string, name string, crd CRDDefinition) error {
	crdTempDir, err := os.MkdirTemp("", "crd")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(crdTempDir)
	crdFiles, err := downloadCRDs(crd, crdTempDir)
	if err != nil {
		return fmt.Errorf("error downloading CRD files: %v", err)
	}

	languageOption := fmt.Sprintf("--%s", language)
	packagePathOption := fmt.Sprintf("--%sPath", language)
	packagePath := fmt.Sprintf("%s/%s", name, language)

	args := append([]string{languageOption, packagePathOption, packagePath}, crdFiles...)
	fmt.Printf("crd2pulumi args: %v\n", args)
	cmd := exec.Command("crd2pulumi", args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return fmt.Errorf("Error running crd2pulumi: %v", err)
	}

	return nil
}

func downloadCRDs(crd CRDDefinition, path string) ([]string, error) {
	versionedCrdUrls := make([]string, len(crd.CRD))
	for i, v := range crd.CRD {
		versionedCrdUrls[i] = strings.Replace(v, "${VERSION}", crd.Version, -1)
	}
	crdFileNames := make([]string, len(versionedCrdUrls))
	for i, v := range versionedCrdUrls {
		crdFileNames[i] = strings.Split(v, "/")[len(strings.Split(v, "/"))-1]
	}
	downloadedCrdFiles := make([]string, len(crdFileNames))
	for i, crdFileName := range crdFileNames {
		downloadedCrdFiles[i] = fmt.Sprintf("%s/%s", path, crdFileName)
		err := downloadFile(downloadedCrdFiles[i], versionedCrdUrls[i])
		if err != nil {
			return []string{}, err
		}
	}
	return downloadedCrdFiles, nil
}

// DownloadFile will download from a given url to a file. It will
// write as it downloads (useful for large files).
func downloadFile(filepath string, url string) error {
	fmt.Printf("Downloading CRD %v to %v\n", url, filepath)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download CRD %v: %v", url, err)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create temporary file %v: %v", filepath, err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
