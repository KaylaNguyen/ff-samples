# Usage

`docker build -t helloworld . && docker run --rm -p 8080:8080 -e PORT=8080 helloworld`

gcloud functions deploy hello \
--runtime python39 --trigger-http --allow-unauthenticated