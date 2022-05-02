#!/bin/bash
export VERSION=$(jq -r .version meta.json)
export BUILD=$(TZ=UTC git --no-pager show --quiet --abbrev=40 --format='%h')
export BUILD_HASH=$(TZ=UTC git --no-pager show --quiet --abbrev=8 --format='%h')
export BUILD_DATE=$(TZ=UTC git --no-pager show --quiet --date='format-local:%Y%m%d%H%M%S' --format='%cd')
export BUILD_VAR_PATH=github.com/aklinker1/miasma/internal/shared/constants
