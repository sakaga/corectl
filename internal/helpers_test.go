package internal

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestRelativeToProject(t *testing.T) {

	res := RelativeToProject("/etc/corectl.yml", "./myscript.qvs")
	assert.Equal(t, "/etc/myscript.qvs", strings.Replace(res, string(os.PathSeparator), "/", -1))

	res = RelativeToProject("/etc/corectl.yml", "scripts/myscript.qvs")
	assert.Equal(t, "/etc/scripts/myscript.qvs", strings.Replace(res, string(os.PathSeparator), "/", -1))

	res = RelativeToProject("/etc/corectl.yml", "../scripts/myscript.qvs")
	assert.Equal(t, "/scripts/myscript.qvs", strings.Replace(res, string(os.PathSeparator), "/", -1))
}
