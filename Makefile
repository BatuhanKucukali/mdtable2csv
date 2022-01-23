release:
	goreleaser release --rm-dist --skip-publish

snapshot:
	goreleaser release --rm-dist --snapshot