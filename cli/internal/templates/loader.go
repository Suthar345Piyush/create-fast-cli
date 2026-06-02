// loader will help loading templates from embed filesystem

package templates

import (
	"fmt"
	"io/fs"

	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
)

// it will load a subfilesystem - with given app type and framework

// inside files layout will like -

/*

   files/
	    common/
			dev-tool/cobra/
			dev-tool/urfave/
			git-client/cobra/
			ai-assistant/cobra/

*/

// sub file-system with apptype and framework
// we have some common files and some are typed

func SubFS(appType config.AppType, framework config.Framework) (common fs.FS, typed fs.FS, err error) {

	commonRoot := "file/common"
	typedRoot := fmt.Sprintf("files/%s/%s", appType, framework)

	common, err = fs.Sub(FS, commonRoot)

	if err != nil {
		return nil, nil, fmt.Errorf("missing common templates at %q: %w", commonRoot, err)
	}

	typed, err = fs.Sub(FS, typedRoot)

	if err != nil {
		return nil, nil, fmt.Errorf("no template found for apptype=%q and framework=%q (looked at %q): %w", appType, framework, typedRoot, err)
	}

	return common, typed, err
}
