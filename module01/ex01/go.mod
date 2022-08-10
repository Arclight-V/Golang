module ex01

require dbreader v0.0.0

replace dbreader v0.0.0 => ../ex00/service/dbreader

require handler v0.0.0

replace handler v0.0.0 => ../ex00/internal/app/handler

go 1.18
