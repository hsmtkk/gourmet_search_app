package parse_test

import (
	"gourmet_serarch_app/parse"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	content, err := os.ReadFile("./sample.xml")
	assert.Nil(t, err)
	wants := []parse.Shop{
		{ID: "J001263511", Name: "発酵串揚げとワイン あげは", LogoImage: "https://imgfp.hotp.jp/SYS/cmn/images/common/diary/custom/m30_img_noimage.gif"},
		{ID: "J001278663", Name: "アジアンさくらダイニング虎ノ門", LogoImage: "https://imgfp.hotp.jp/IMGH/36/27/P038293627/P038293627_69.jpg"},
	}
	gots, err := parse.Parse(content)
	assert.Nil(t, err)
	assert.Equal(t, wants[0], gots[0])
	assert.Equal(t, wants[1], gots[1])
}
