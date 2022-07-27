JWT_VAREJAO :=  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfaWQiOjIsImNsaWVudF9uYW1lIjoidmFyZWphbyIsImV4cCI6MTY2MTA1MjQyOX0.LDNEGM0k345S4oeKNSRN88rA6exVGd_ILBiVjfT95qc
JWT_MACAPA := eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfaWQiOjEsImNsaWVudF9uYW1lIjoibWFjYXBhIiwiZXhwIjoxNjYxMDUzNDUxfQ.gCCWiriJRxXQzHCGbDTvRMtEPerkYDjtYPUsWHfzk8Q 

create-varejao:
	curl -XPOST http://localhost:8082/contact -d@sampledata/contacts-varejao.json -i \
		-H 'Authorization: $(JWT_VAREJAO)'

create-macapa:
	curl -XPOST http://localhost:8082/contact -d@sampledata/contacts-macapa.json -i \
		-H 'Authorization: $(JWT_MACAPA)'

up-db:
	docker-compose up -d mysql postgres

wait-db:
	until docker-compose logs mysql 2>/dev/null | grep 'ready for connections' -q; do \
		sleep 3; \
	done

	until docker-compose logs postgres 2>/dev/null | grep 'ready to accept connections' -q; do \
		sleep 3; \
	done

up-api:
	docker-compose up -d api

wait-api:
	until docker-compose logs api 2>/dev/null | grep -q 'starting api server'; do \
		sleep 3; \
	done

run: up-db wait-db up-api wait-api create-varejao create-macapa