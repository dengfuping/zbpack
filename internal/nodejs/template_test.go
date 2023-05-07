package nodejs_test

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"github.com/zeabur/zbpack/internal/nodejs"
)

func TestTemplate_NBuildCmd_NOutputDir(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",

		InstallCmd: "yarn install",
		BuildCmd:   "",
		StartCmd:   "yarn start",

		OutputDir: "",
		SPA:       false,
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_NBuildCmd_OutputDir_NSPA(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",

		InstallCmd: "yarn install",
		BuildCmd:   "",
		StartCmd:   "yarn start",

		OutputDir: "dist",
		SPA:       false,
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_NBuildCmd_OutputDir_SPA(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",

		InstallCmd: "yarn install",
		BuildCmd:   "",
		StartCmd:   "yarn start",

		OutputDir: "dist",
		SPA:       true,
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_BuildCmd_NOutputDir(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",

		InstallCmd: "yarn install",
		BuildCmd:   "yarn build",
		StartCmd:   "yarn start",

		OutputDir: "",
		SPA:       false,
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_BuildCmd_OutputDir(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",

		InstallCmd: "yarn install",
		BuildCmd:   "yarn build",
		StartCmd:   "yarn start",

		OutputDir: "dist",
		SPA:       true,
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}