# Code samples to test Function Frameworks

To send a cloud event to localhost:8080
```
curl localhost:8080 \                                                                                                                                                                                                                             [42b7ed6]
-H "Content-Type: application/cloudevents+json" \
-d '{
"specversion" : "1.0",
"type" : "com.github.pull.create",
"source" : "https://github.com/cloudevents/spec/pull",
"subject" : "123",
"id" : "ce",
"time" : "2021-01-27T18:30:00Z",
"data" : "hello"
}' -v
```