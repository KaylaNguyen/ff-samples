# Usage

`export FUNCTION_TARGET=com.example.CEFunction` or

`export FUNCTION_TARGET=com.example.HelloWorld`

`mvn function:run`

gcloud functions deploy java-helloworld \
--entry-point com.example.HelloWorld \
--runtime java11 \
--memory 512MB --trigger-http --allow-unauthenticated