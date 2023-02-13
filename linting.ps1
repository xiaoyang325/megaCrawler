$location = "$(Get-Location)"
golangci-lint.exe run --path-prefix $location $args[0]
