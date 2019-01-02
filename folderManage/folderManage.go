package folderManage

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
