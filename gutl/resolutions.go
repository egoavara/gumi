package gutl

import "sort"

func init() {
	sort.Sort(DefinedResolutions)
}

var DefinedResolutions = STDResolutions([]STDResolution{
	// CIF
	STDResolution{"Sub-QCIF", 128, 96},
	STDResolution{"QCIF", 220, 176},
	STDResolution{"CIF", 352, 288},
	// CGA
	STDResolution{"CGA", 320, 200},
	// VGA
	STDResolution{"qVGA", 320, 240},
	STDResolution{"HVGA", 480, 320},
	STDResolution{"VGA", 640, 480},
	STDResolution{"SVGA", 800, 600},
	STDResolution{"WSVGA", 1024, 576},
	// XGA
	STDResolution{"XGA", 1024, 768},
	STDResolution{"FWXGA", 1366, 768},
	STDResolution{"SXGA", 1280, 1024},
	STDResolution{"UXGA", 1600, 1200},
	// HD
	STDResolution{"nHD", 640, 360},
	STDResolution{"qHD", 960, 540},
	STDResolution{"HD", 1280, 720},
	STDResolution{"720p", 1280, 720},
	STDResolution{"HD+", 1600, 900},
	STDResolution{"FHD", 1920, 1080},
	STDResolution{"1080p", 1920, 1080},
	STDResolution{"1080i", 1920, 1080},
	STDResolution{"QHD", 2560, 1440},
	STDResolution{"2K", 2560, 1440},
	STDResolution{"UHD", 3840, 2160},
	STDResolution{"4K", 3840, 2160},
	STDResolution{"UHD+", 5120, 2880},
	STDResolution{"5K", 5120, 2880},
	STDResolution{"QUHD", 7680, 4320},
	STDResolution{"8K", 7680, 4320},
})

// sorting util
type STDResolution struct {
	Name   string
	Width  int
	Height int
}
type STDResolutions []STDResolution

func (s STDResolutions) Len() int {
	return len(([]STDResolution)(s))
}
func (s STDResolutions) Swap(i, j int) {
	temp := ([]STDResolution)(s)
	temp[i], temp[j] = temp[j], temp[i]
}
func (s STDResolutions) Less(i, j int) bool {
	temp := ([]STDResolution)(s)
	return temp[i].Width*temp[i].Height < temp[j].Width*temp[j].Height
}
func (s STDResolutions) Smaller(w, h int) (res []STDResolution) {
	for _, v := range s{
		if v.Width <= w && v.Height <= h{
			res = append(res, v)
		}
	}
	return res
}
func (s STDResolutions) Get(name string) (w, h int) {
	for _, v := range ([]STDResolution)(s) {
		if v.Name == name {
			return v.Width, v.Height
		}
	}
	return 0, 0
}
func (s STDResolutions) Kinds() []string {
	tmp := ([]STDResolution)(s)
	temp := make([]string, len(tmp))
	for i, v := range tmp {
		temp[i] = v.Name
	}
	return temp
}
func (s STDResolutions) Under(w, h int) []string {
	var temp []string
	for _, v := range ([]STDResolution)(s) {
		if v.Width <= w && v.Height < h {
			temp = append(temp, v.Name)
		}
	}
	return temp
}
