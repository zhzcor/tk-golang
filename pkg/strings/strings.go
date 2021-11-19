package strings

import "regexp"

// 去掉windows/mac中不能保存的文件特殊字符
func ParseFileName(fname string) (string, error) {
	reg, err := regexp.Compile(`[:|\*|\\|/|?|"|<|>|\|]`)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(fname, " "), nil
}
