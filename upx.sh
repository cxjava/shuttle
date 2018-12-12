#!/bin/bash
find ./dist -xdev -maxdepth 3 -type f -name 'shuttle*' | xargs upx
find ./dist -xdev -maxdepth 3 -type f -name 'upgrade*' | xargs upx