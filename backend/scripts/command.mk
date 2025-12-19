.PHONY: gen-api
gen-api:
	powershell -ExecutionPolicy Bypass -File ./scripts/generate.ps1

