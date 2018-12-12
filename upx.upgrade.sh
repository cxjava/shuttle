#!/bin/bash
find ./dist -xdev -maxdepth 3 -type f -name 'upgrade*' | xargs upx