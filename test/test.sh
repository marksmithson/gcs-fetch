#!/bin/bash

rm data/out/output.csv

docker run --rm -v $(pwd)/secrets:/config -v $(pwd)/data/out:/out -e GOOGLE_APPLICATION_CREDENTIALS=/config/service-account.json gcs-fetch gs://data-drover-input/input.csv /out/output.csv

cat data/out/output.csv