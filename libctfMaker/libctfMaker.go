package libctfMaker

import "os"

// Creates a folder if it doesn't exist
func CreateDir(dir string) {
      if _, err := os.Stat(dir); os.IsNotExist(err) {
              err = os.MkdirAll(dir, 0755)
              if err != nil {
                      panic(err)
              }
      }
}

// Check if a file or folder exists
func FileExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}
