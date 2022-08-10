module ex00

require dbreader v0.0.0 // indirect

replace dbreader v0.0.0 => ./service/dbreader

require handler v0.0.0

replace handler v0.0.0 => ./internal/app/handler

go 1.18
