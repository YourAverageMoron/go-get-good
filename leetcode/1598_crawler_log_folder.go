package main

func CrawlerLogFolder(logs []string) int {
	l := 0
	for _, log := range logs {
		if log == "./" {
			continue
		}
		if log == "../" {
			if l > 0 {
				l--
			}
			continue
		}
		l++
	}
	return l

}

func CrawlerLogFolderWithStack(logs []string) (int, []string) {
	s := make([]string, len(logs))
	l := 0
	for _, log := range logs {
		if log == "./" {
			continue
		}
		if log == "../" {
			if l > 0 {
				l--
			}
			continue
		}
		s[l] = log
		l++
	}
	res := []string{}
	for i := 0; i < l; i++ {
		res = append(res, s[i])
	}
	return l, res
}
