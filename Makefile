default: doc

doc:
	docker run --workdir /opt --rm -v $$PWD:/opt -it josdotso/pp README.pp > README.md

.PHONY: doc
