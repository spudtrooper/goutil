package curlimport

func Parse(content string) (CurlCmd, error) {
	cc, err := parseCurlCmd(content)
	if err != nil {
		return CurlCmd{}, err
	}
	return *cc, nil
}

func Import(content, outfile string, run, createBodyStruct, unescape bool) error {
	return curlImport(content, outfile, run, createBodyStruct, unescape)
}
