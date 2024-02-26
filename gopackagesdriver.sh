#!/usr/bin/env zsh

exec bazelisk run -- @rules_go//go/tools/gopackagesdriver "${@}"
