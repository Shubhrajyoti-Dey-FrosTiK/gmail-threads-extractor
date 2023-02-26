#!/bin/sh
direnv allow . && gcloud functions deploy gmail-labels-extractor --runtime go119 --trigger-http --allow-unauthenticated --region=asia-south1 --gen2 --entry-point threads