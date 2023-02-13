#!/usr/bin/env bash

for FILE in dist/reconbot_*/*; do
    du -sh ${FILE}
    upx ${FILE}
    du -sh ${FILE}
done
