module Golang/module01/ex02

require checkextension v0.0.0

replace checkextension v0.0.0 => ./internal/app/checkExtension

require parser v0.0.0

replace parser v0.0.0 => ./internal/app/parser

go 1.18
