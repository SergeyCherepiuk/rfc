package load

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
)

type RfcLoader struct {
	latestRfcNumber int
}

func NewRfcLoader() (RfcLoader, error) {
	latestRfcNumber, err := getLatestRfcNumber()
	if err != nil {
		err = fmt.Errorf("failed to get latest RFC number: %w", err)
		return RfcLoader{}, err
	}

	return RfcLoader{latestRfcNumber}, nil
}

func getLatestRfcNumber() (int, error) {
	const lastestRfcsUrl = "https://www.ietf.org/rfc/rfc-index-latest.txt"

	body, err := getRequestBody(lastestRfcsUrl)
	if err != nil {
		err = fmt.Errorf("failed to load list of latest RFCs: %w", err)
		return 0, err
	}

	paragraphs := bytes.Split(body, []byte{'\n', '\n'})
	slices.Reverse(paragraphs)

	const targetParagraphIndex = 2

	if len(paragraphs)-1 < targetParagraphIndex {
		err = fmt.Errorf("incorrectly formatted list")
		err = fmt.Errorf("failed to process list of latest RFCs: %w", err)
		return 0, err
	}

	targetParagraph := paragraphs[targetParagraphIndex]
	spaceIndex := bytes.IndexByte(targetParagraph, ' ')
	latestRfcNumberStr := string(targetParagraph[:spaceIndex])

	latestRfcNumber, err := strconv.ParseInt(latestRfcNumberStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("failed to process list of latest RFCs: %w", err)
		return 0, err
	}

	return int(latestRfcNumber), nil
}

func (l RfcLoader) Load(rfcNumber int) ([]byte, error) {
	if rfcNumber > l.latestRfcNumber {
		err := fmt.Errorf("RFC number (%d) is higher than latest available (%d)", rfcNumber, l.latestRfcNumber)
		return nil, err
	}

	rfcUrl := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", rfcNumber)

	rfc, err := getRequestBody(rfcUrl)
	if err != nil {
		err = fmt.Errorf("failed to load RFC %d: %w", rfcNumber, err)
		return nil, err
	}

	return rfc, nil
}

func getRequestBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("request failed: %w", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil && len(body) != 0 {
			err = fmt.Errorf("request failed: %s", body)
		} else {
			err = fmt.Errorf("request failed: status code %d", resp.StatusCode)
		}
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to process response: %w", err)
		return nil, err
	}

	return body, nil
}
