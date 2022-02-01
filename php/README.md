# Usage

`composer require google/cloud-functions-framework`

`Remove first line "#!/usr/bin/env php" in vendor/bin/router.php`

`export FUNCTION_TARGET=helloHttp`

or

```
export FUNCTION_TARGET=helloCloudEvent
export FUNCTION_SIGNATURE_TYPE=cloudevent
```

`php -S localhost:8080 vendor/bin/router.php`

gcloud functions deploy helloHttp --runtime php74 --trigger-http --allow-unauthenticated