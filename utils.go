package auth

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/oklog/ulid"

	errors "github.com/joaosoft/errors"
)

func getEnv() string {
	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}

	return env
}

func exists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func readFile(file string, obj interface{}) ([]byte, error) {
	var err error

	if !exists(file) {
		return nil, errors.New(errors.LevelError, 0, "file don't exist")
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if obj != nil {
		if err := json.Unmarshal(data, obj); err != nil {
			return nil, err
		}
	}

	return data, nil
}

func readFileLines(file string) ([]string, error) {
	lines := make([]string, 0)

	if !exists(file) {
		return nil, errors.New(errors.LevelError, 0, "file don't exist")
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func writeFile(file string, obj interface{}) error {
	if !exists(file) {
		return errors.New(errors.LevelError, 0, "file don't exist")
	}

	jsonBytes, _ := json.MarshalIndent(obj, "", "    ")
	if err := ioutil.WriteFile(file, jsonBytes, 0644); err != nil {
		return err
	}

	return nil
}

func valUI(id string) error {
	_, err := ulid.Parse(id)
	return err
}

func genUI() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
